package registry

import (
	"github.com/arioki1/qiscus-caa/config"
	"github.com/arioki1/qiscus-caa/src/api/v1/model"
	"github.com/arioki1/qiscus-caa/src/api/v1/repository"
	"sync"
)

type repositoryRegistry struct {
	cfg config.Config
}

func (r repositoryRegistry) QiscusCAARepository() model.QiscusCAARepository {
	var qu model.QiscusCAARepository
	var loadOne sync.Once

	loadOne.Do(func() {
		qu = repository.NewQiscusCAARepository(r.cfg)
	})
	return qu
}

type RepositoryRegistry interface {
	QiscusCAARepository() model.QiscusCAARepository
}

func NewRepositoryRegistry(cfg config.Config) RepositoryRegistry {
	var repoRegistry RepositoryRegistry
	var loadOne sync.Once

	loadOne.Do(func() {
		repoRegistry = &repositoryRegistry{
			cfg: cfg,
		}
	})

	return repoRegistry
}
