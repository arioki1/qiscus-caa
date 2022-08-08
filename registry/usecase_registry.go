package registry

import (
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/src/api/v1/model"
	"github.com/arioki1/qiscus-caa/src/api/v1/usecase"
	"sync"
)

type useCaseRegistry struct {
	repo RepositoryRegistry
	cfg  config.Config
}

type UseCaseRegistry interface {
	QiscusWebhook() model.QiscusWebhookUseCase
}

func NewUseCaseRegistry(repo RepositoryRegistry, cfg config.Config) UseCaseRegistry {
	var uc UseCaseRegistry
	var loadOne sync.Once
	loadOne.Do(func() {
		uc = &useCaseRegistry{
			repo: repo,
			cfg:  cfg,
		}
	})

	return uc
}

func (u useCaseRegistry) QiscusWebhook() model.QiscusWebhookUseCase {
	var qu model.QiscusWebhookUseCase
	var loadOne sync.Once

	loadOne.Do(func() {
		qu = usecase.NewQiscusUseCase(u.cfg, u.repo.QiscusCAARepository())
	})

	return qu
}
