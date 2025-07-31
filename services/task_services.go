package services

import (
	"context"

	"github.com/joaogdoqr/taskfy/store"
)

type TaskService struct {
	Store *store.Queries
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string, priority int32) (store.Task, error) {
	task, err := s.Store.CreateTask(ctx, store.CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	task, err := s.Store.GetTaskById(ctx, id)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) ListTasks(ctx context.Context) ([]store.Task, error) {
	tasks, err := s.Store.ListTasks(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int32, title, description string, priority int32) (store.Task, error) {
	task, err := s.Store.UpdateTask(ctx, store.UpdateTaskParams{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int32) error {
	return s.Store.DeleteTask(ctx, id)
}
