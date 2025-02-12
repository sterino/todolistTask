package repository

import (
	"todoList/internal/config"
	"todoList/internal/domain/task"
	"todoList/internal/repository/postgres"
	"todoList/pkg/store"
)

type Configuration func(r *Repository) error

type Repository struct {
	postgres store.SQLX

	Task task.Repository
}

func New(configs ...Configuration) (s *Repository, err error) {
	s = &Repository{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func (r *Repository) Close() {
	if r.postgres.Client != nil {
		r.postgres.Client.Close()
	}

}

func WithPostgresStore(cfg config.Config) Configuration {
	return func(s *Repository) (err error) {

		s.postgres, err = store.NewSQL(cfg)
		if err != nil {
			return
		}

		if err = store.Migrate(s.postgres); err != nil {
			return
		}

		s.Task = postgres.NewTaskRepository(s.postgres.Client)

		return
	}
}
