package repository

import (
	"errors"
	"sync"
)

type Repository struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SaveURL(url, hash string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.urls[hash] = url
	return nil
}

func (r *Repository) GetURL(hash string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, ok := r.urls[hash]
	if !ok {
		return "", errors.New("url not found")
	}
	return url, nil
}
