package data_access

import (
	"context"
	"server/common"
	"server/model"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	clicksKey      = ":clicks"
	urlCountsKey   = "count:shortened"
	urlsIdsListKey = "urls:shortened"
	originalURLKey = "originalURL"
)

type URLStore struct {
	client *redis.Client
}

func urlClicks(shortUrl string) string {
	return shortUrl + clicksKey
}

func NewURLStore() *URLStore {
	addr := common.GetEnv(common.RedisUrl)
	password := common.GetEnv(common.RedisPassword)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &URLStore{client: client}
}

func (s *URLStore) Insert(urlId, originalURL, user string) error {
	ctx := context.Background()

	err := s.client.HSet(ctx, urlId, originalURLKey, originalURL, userIdKey, user).Err()
	if err != nil {
		return err
	}

	err = s.client.Set(ctx, urlClicks(urlId), "0", 0).Err()
	if err != nil {
		return err
	}

	err = s.client.Incr(ctx, urlCountsKey).Err()
	if err != nil {
		return err
	}

	err = s.client.LPush(ctx, urlsIdsListKey, urlId).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *URLStore) GetOriginalURL(urlId string) (string, error) {
	ctx := context.Background()
	result, err := s.client.HGet(ctx, urlId, originalURLKey).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *URLStore) IncrementClicksCount(urlId string) error {
	ctx := context.Background()
	err := s.client.Incr(ctx, urlClicks(urlId)).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *URLStore) GetClickCount(urlId string) (int, error) {
	ctx := context.Background()
	result, err := s.client.Get(ctx, urlClicks(urlId)).Result()
	if err != nil {
		return 0, err
	}

	clickCount, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return clickCount, nil
}

func (s *URLStore) GetAllShortenedURLs() ([]model.Url, error) {
	ctx := context.Background()

	urlIds, err := s.client.LRange(ctx, urlsIdsListKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	urlClickCounts := make([]model.Url, len(urlIds))

	for i, urlId := range urlIds {
		clickCount, err := s.GetClickCount(urlId)
		if err != nil {
			return nil, err
		}

		result, err := s.client.HGetAll(ctx, urlId).Result()
		if err != nil {
			return nil, err
		}

		urlClickCounts[i] = model.Url{ShortURL: common.IdToUrl(urlId), OriginalURL: result[originalURLKey], ClickCount: clickCount, UserID: result[userIdKey]}
	}

	return urlClickCounts, nil
}

func (s *URLStore) GetShortenedURLCount() (int, error) {
	ctx := context.Background()
	result, err := s.client.Get(ctx, urlCountsKey).Result()

	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return count, nil
}
