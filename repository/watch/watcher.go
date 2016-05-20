package watch

import (
	"sync"

	"github.com/apex/log"
	"github.com/cleung2010/go-git2consul/repository"
)

type Watcher struct {
	sync.Mutex

	Repositories []*repository.Repository

	RepoChangeCh chan *repository.Repository
	ErrCh        chan error

	logger log.Interface
}

func New(repos []*repository.Repository) *Watcher {
	repoChangeCh := make(chan *repository.Repository, len(repos))
	errCh := make(chan error)

	logger := log.Log
	logger.WithField("caller", "repo-watcher")

	return &Watcher{
		Repositories: repos,
		RepoChangeCh: repoChangeCh,
		ErrCh:        errCh,
		logger:       logger,
	}
}