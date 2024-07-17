package repository

import (
	"errors"
	"sync"
)

type Repository struct {
	urls sync.Map
}

func NewRepository() *Repository {
	return &Repository{
		urls: sync.Map{},
	}
}

func (r *Repository) SaveURL(url, hash string) error {
	r.urls.Store(hash, url)
	return nil
}

func (r *Repository) GetURL(hash string) (string, error) {
	url, ok := r.urls.Load(hash)
	if !ok {
		return "", errors.New("url not found")
	}
	return url.(string), nil
}
