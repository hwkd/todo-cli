package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Store is an interface that defines the methods to save and load todos from disk.
type Store interface {
	Save(todos []TodoItem) error
	Load() ([]TodoItem, error)
}

// TodoListCsvStore is a store that saves and loads todos to and from a CSV file.
type TodoListCsvStore struct {
	filepath string
}

// NewTodoListCsvStore creates a new instance of TodoListCsvStore.
func NewTodoListCsvStore(filepath string) *TodoListCsvStore {
	return &TodoListCsvStore{
		filepath: filepath,
	}
}

// Save writes the list of todos to a CSV file.
func (t *TodoListCsvStore) Save(todos []TodoItem) error {
	file, err := os.Create(t.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	records := make([][]string, len(todos))
	for i := range todos {
		todo := &todos[i]
		records[i] = append(records[i], todo.ID)
		records[i] = append(records[i], todo.Title)
		records[i] = append(records[i], todo.Description)
		records[i] = append(records[i], strconv.FormatBool(todo.IsDone))
		records[i] = append(records[i], todo.CreatedAt.Format("2006-01-02T15:04:05Z07:00"))
		records[i] = append(records[i], todo.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"))
	}

	if err = writer.WriteAll(records); err != nil {
		return fmt.Errorf("%w: Error writing data to CSV", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("%w: Error flushing data to CSV", err)
	}

	return nil
}

// Load reads the CSV file and returns the list of todos.
func (t *TodoListCsvStore) Load() ([]TodoItem, error) {
	file, err := os.OpenFile(t.filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	todos := make([]TodoItem, len(records))
	for i, rec := range records {
		todo, err := NewTodoItemFromStrings(
			rec[0],
			rec[1],
			rec[2],
			rec[3],
			rec[4],
			rec[5],
		)
		if err != nil {
			return nil, err
		}
		todos[i] = *todo
	}

	return todos, nil
}
