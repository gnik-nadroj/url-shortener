package data_access

import (
	"context"
	"server/common"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type URLStore struct {
	client *redis.Client
}

type URLClickCount struct {
    ShortURL   string
    ClickCount int
}

func NewURLStore() *URLStore {
	addr := common.GetEnv("REDIS_URL")
	password := common.GetEnv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &URLStore{client: client}
}

func (s *URLStore) Insert(shortURL, originalURL string) error {
    ctx := context.Background()

    err := s.client.Set(ctx, shortURL, originalURL, 0).Err()
    if err != nil {
        return err
    }

    err = s.client.Set(ctx, shortURL+":clicks", "0", 0).Err()
    if err != nil {
        return err
    }

    err = s.client.Incr(ctx, "count:shortened").Err()
    if err != nil {
        return err
    }

    err = s.client.LPush(ctx, "urls:shortened", shortURL).Err()

    return err
}

func (s *URLStore) GetOriginalURL(shortURL string) (string, error) {
	ctx := context.Background()
	result, err := s.client.Get(ctx, shortURL).Result()
	if err != nil {
		return "", err
	}

	err = s.client.Incr(ctx, shortURL+":clicks").Err()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *URLStore) GetClickCount(shortURL string) (int, error) {
	ctx := context.Background()
	result, err := s.client.Get(ctx, shortURL+":clicks").Result()
	if err != nil {
		return 0, err
	}

	clickCount, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return clickCount, nil
}

func (s *URLStore) GetAllShortenedURLs() ([]URLClickCount, error) {
    ctx := context.Background()
    urls, err := s.client.LRange(ctx, "urls:shortened", 0, -1).Result()
    if err != nil {
        return nil, err
    }

    urlClickCounts := make([]URLClickCount, len(urls))
    for i, url := range urls {
        clickCount, err := s.GetClickCount(url)
        if err != nil {
            return nil, err
        }
        urlClickCounts[i] = URLClickCount{ShortURL: common.ComposeUrl(url), ClickCount: clickCount}
    }

    return urlClickCounts, nil
}

func (s *URLStore) GetShortenedURLCount() (int, error) {
    ctx := context.Background()
    result, err := s.client.Get(ctx, "count:shortened").Result()
    if err != nil {
        return 0, err
    }

    count, err := strconv.Atoi(result)
    if err != nil {
        return 0, err
    }

    return count, nil
}
