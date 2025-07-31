package services

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joaogdoqr/taskfy/store"
	"github.com/stretchr/testify/assert"
)

type MockTaskService struct {
}

func (m *MockTaskService) CreateTask(title, description string, priority int32) (store.Task, error) {
	return store.Task{
		ID:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}, nil
}

func (m *MockTaskService) GetTaskById(id int32) (store.Task, error) {
	return store.Task{
		ID:          id,
		Title:       "test",
		Description: "test",
		Priority:    8000,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}, nil
}

func (m *MockTaskService) ListTasks() ([]store.Task, error) {
	return []store.Task{
		{
			ID:          1,
			Title:       "test",
			Description: "test",
			Priority:    8000,
			CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
		{
			ID:          2,
			Title:       "test",
			Description: "test",
			Priority:    8000,
			CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
		{
			ID:          3,
			Title:       "test",
			Description: "test",
			Priority:    8000,
			CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
	}, nil
}

func (m *MockTaskService) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {
	return store.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}, nil
}

func (m *MockTaskService) DeleteTask(id int32) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	mockService := &MockTaskService{}

	task, err := mockService.CreateTask("test", "test desc", 1)
	if err != nil {
		t.Fatal("Failed to create task")
	}

	assert.NoError(t, err)
	assert.Equal(t, "test", task.Title)
	assert.Equal(t, "test desc", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestGetTask(t *testing.T) {
	mockService := &MockTaskService{}

	task, err := mockService.GetTaskById(1)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.ID)
}

func TestListTask(t *testing.T) {
	mockService := &MockTaskService{}

	tasks, err := mockService.ListTasks()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(tasks))
}

func TestUpdateTask(t *testing.T) {
	mockService := &MockTaskService{}

	task, err := mockService.UpdateTask(1, "test updated", "test desc updated", 2)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.ID)
	assert.Equal(t, "test updated", task.Title)
	assert.Equal(t, "test desc updated", task.Description)
	assert.Equal(t, int32(2), task.Priority)
}

func TestDeleteTask(t *testing.T) {
	mockService := &MockTaskService{}

	err := mockService.DeleteTask(1)

	assert.NoError(t, err)
}
