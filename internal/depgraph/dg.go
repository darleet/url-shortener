package depgraph

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"shortener/internal/commands/cmdargs"
	"sync"
)

type dgEntity[T any] struct {
	sync.Once
	value   T
	initErr error
}

func (e *dgEntity[T]) get(init func() (T, error)) (T, error) {
	e.Do(func() {
		e.value, e.initErr = init()
	})
	if e.initErr != nil {
		return *new(T), e.initErr
	}
	return e.value, nil
}

type DepGraph struct {
	logger *dgEntity[*zap.SugaredLogger]
	repo   *dgEntity[*pgxpool.Pool]
	config cmdargs.RunArgs
}

func NewDepGraph(config cmdargs.RunArgs) *DepGraph {
	return &DepGraph{
		logger: &dgEntity[*zap.SugaredLogger]{},
		repo:   &dgEntity[*pgxpool.Pool]{},
		config: config,
	}
}

func (d *DepGraph) GetLogger() (*zap.SugaredLogger, error) {
	return d.logger.get(func() (*zap.SugaredLogger, error) {
		return zap.S(), nil
	})
}

func (d *DepGraph) GetRepo() (*pgxpool.Pool, error) {
	return d.repo.get(func() (*pgxpool.Pool, error) {
		return pgxpool.New(
			context.Background(),
			d.config.DatabaseURL,
		)
	})
}
