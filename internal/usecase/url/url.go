package url

import (
	"crypto/sha256"
	"encoding/hex"
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

func NewUsecase(repo Repo, config cmdargs.RunArgs) *Usecase {
	return &Usecase{
		repo:   repo,
		config: config,
	}
}

func (u *Usecase) Shorten(url string) (string, error) {
	hash := sha256.Sum256([]byte(url))
	hashString := hex.EncodeToString(hash[:])[:8]
	err := u.repo.SaveURL(url, hashString)
	if err != nil {
		return "", err
	}
	return u.config.ShortenerURL + hashString, nil
}

func (u *Usecase) Expand(url string) (string, error) {
	return u.repo.GetURL(url)
}
