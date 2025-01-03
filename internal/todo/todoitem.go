package todo

import (
	"crypto/sha1"
	"encoding/hex"
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
	todo := &TodoItem{
		ID:          "",
		Title:       title,
		Description: desc,
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		dirty:       false,
		deleted:     false,
	}
	todo.assignID()
	return todo
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

// assignID assigns a unique ID to the TodoItem.
func (todo *TodoItem) assignID() {
	now := time.Now()
	hasher := sha1.New()
	hasher.Write([]byte(todo.Title + todo.Description + now.Format("2006-01-02 15:04:05 MST")))
	hashSum := hasher.Sum(nil)
	todo.ID = hex.EncodeToString(hashSum)[:16]
}

func (todo *TodoItem) Done() {
	todo.IsDone = true
}

func (todo *TodoItem) Undone() {
	todo.IsDone = false
}
