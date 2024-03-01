package data_access

import (
	"testing"
)

func TestNewURLStore(t *testing.T) {
    store := NewURLStore()
    if store == nil {
        t.Errorf("NewURLStore() = nil; want non-nil")
    }
}

func TestInsert(t *testing.T) {
    store := NewURLStore()
    err := store.Insert("shortURL", "originalURL", user.ID)
    if err != nil {
        t.Errorf("Insert() error = %v; want nil", err)
    }
}

func TestGetOriginalURL(t *testing.T) {
    store := NewURLStore()

    url, err := store.GetOriginalURL("shortURL")
    if err != nil {
        t.Errorf("GetOriginalURL() error = %v; want nil", err)
    }
    if url != "originalURL" {
        t.Errorf("GetOriginalURL() = %v; want 'originalURL'", url)
    }
}

func TestGetClickCount(t *testing.T) {
    store := NewURLStore()

    _, err := store.GetClickCount("shortURL")
    if err != nil {
        t.Errorf("GetClickCount() error = %v; want nil", err)
    }
}

func TestGetShortenedURLCount(t *testing.T) {
    store := NewURLStore()

    count, err := store.GetShortenedURLCount()
    if err != nil {
        t.Errorf("GetShortenedURLCount() error = %v; want nil", err)
    }
    if count <= 0 {
        t.Errorf("GetShortenedURLCount() = %v; want > 0", count)
    }
}

func TestGetAllShortenedURLs(t *testing.T) {
    store := NewURLStore()

    urls, err := store.GetAllShortenedURLs()
    if err != nil {
        t.Errorf("GetAllShortenedURLs() error = %v; want nil", err)
    }
    if len(urls) <= 0 {
        t.Errorf("GetAllShortenedURLs() = %v; want > 0", len(urls))
    }
}

func TestIncrementClicksCount(t *testing.T) {
    store := NewURLStore()

    err := store.IncrementClicksCount("shortURL")

    if err  != nil {
        t.Errorf("IncrementClicksCount() error = %v; want nil", err)
    }

    count, err := store.GetClickCount("shortURL")

    if err  != nil {
        t.Errorf("IncrementClicksCount() error = %v; want nil", err)
    }

    if(count != 2) {
        t.Errorf("IncrementClicksCount() = %v; want 2", count)
    }  
}

func TestGetOriginalURL_NotFound(t *testing.T) {
    store := NewURLStore()

    _, err := store.GetOriginalURL("nonexistentURL")
    if err == nil {
        t.Errorf("GetOriginalURL() error = nil; want non-nil")
    }
}

func TestGetClickCount_NotFound(t *testing.T) {
    store := NewURLStore()

    _, err := store.GetClickCount("nonexistentURL")
    if err == nil {
        t.Errorf("GetClickCount() error = nil; want non-nil")
    }
}
