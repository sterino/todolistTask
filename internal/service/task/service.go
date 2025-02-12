package task

import (
	"todoList/internal/domain/task"
)

type Configuration func(s *Service) error

type Service struct {
	taskRepository task.Repository
	//taskCache      task.Cache
}

func New(configs ...Configuration) (s *Service, err error) {

	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithTaskRepository(taskRepository task.Repository) Configuration {
	return func(s *Service) error {
		s.taskRepository = taskRepository
		return nil
	}
}

//func WithTaskCache(taskCache task.Cache) Configuration {
//	return func(s *Service) error {
//		s.taskCache = taskCache
//		return nil
//	}
//}
