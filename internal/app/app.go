package app

import (
	"shortener/internal/commands/cmdargs"
	delivery "shortener/internal/delivery/http/url"
	usecase "shortener/internal/usecase/url"
)

type Entrypoint struct {
	repo   usecase.Repo
	config cmdargs.RunArgs
}

func NewEntrypoint(repo usecase.Repo, config cmdargs.RunArgs) *Entrypoint {
	return &Entrypoint{
		repo:   repo,
		config: config,
	}
}

func (e *Entrypoint) Run() error {
	uc := usecase.NewUsecase(e.repo, e.config)
	return delivery.InitRouter(uc).Start(":8080")
}
