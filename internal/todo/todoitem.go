package todo

import (
	"fmt"
	"strconv"
	"time"
)

type TodoItem struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	dirty       bool
	deleted     bool
}

func NewTodoItem(title, desc string) *TodoItem {
	return &TodoItem{
		ID:          "",
		Title:       title,
		Description: desc,
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		dirty:       false,
		deleted:     false,
	}
}

func NewTodoItemFromStrings(id, title, desc, isDone, createdAt, updatedAt string) (*TodoItem, error) {
	isDoneBool, err := strconv.ParseBool(isDone)
	if err != nil {
		return nil, fmt.Errorf("Expected `IsDone` as boolean, got %s", isDone)
	}

	createdAtDateTime, err := time.Parse("2006-01-02T15:04:05Z07:00", createdAt)
	if err != nil {
		return nil, fmt.Errorf("%s. Expected `CreatedAt` as timestamp, got %s", err.Error(), createdAt)
	}

	updatedAtDateTime, err := time.Parse("2006-01-02T15:04:05Z07:00", updatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s. Expected `UpdatedAt` as timestamp, got %s", err.Error(), updatedAt)
	}

	return &TodoItem{
		ID:          id,
		Title:       title,
		Description: desc,
		IsDone:      isDoneBool,
		CreatedAt:   createdAtDateTime,
		UpdatedAt:   updatedAtDateTime,
	}, nil
}

func (todo *TodoItem) Done() {
	todo.IsDone = true
}

func (todo *TodoItem) Undone() {
	todo.IsDone = false
}
