package app

import (
	delivery "shortener/internal/delivery/http/url"
	usecase "shortener/internal/usecase/url"
)

type Entrypoint struct {
	repo usecase.Repo
}

func NewEntrypoint(repo usecase.Repo) *Entrypoint {
	return &Entrypoint{
		repo: repo,
	}
}

func (e *Entrypoint) Run() error {
	uc := usecase.NewUsecase(e.repo)
	return delivery.InitRouter(uc).Start(":8080")
}
