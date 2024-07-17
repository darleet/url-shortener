package url

import (
	"crypto/sha256"
	"shortener/internal/commands/cmdargs"
)

type Repo interface {
	SaveURL(url, hash string) error
	GetURL(hash string) (string, error)
}

type Usecase struct {
	repo   Repo
	config cmdargs.RunArgs
}

func NewUsecase(repo Repo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Shorten(url string) (string, error) {
	shortURL := sha256.Sum256([]byte(url))
	err := u.repo.SaveURL(url, string(shortURL[]))
	if err != nil {
		return "", err
	}
	return u.config.ShortenerURL + string(shortURL[:]), nil
}

func (u *Usecase) Expand(url string) (string, error) {
	return u.repo.GetURL(url)
}
