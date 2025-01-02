package todo

import "testing"

func TestTodoItem(t *testing.T) {
	t.Run("NewTodoItem", func(t *testing.T) {
		todo := NewTodoItem("Task 1", "Created for test")
		want := TodoItem{
			ID:          "",
			Title:       "Task 1",
			Description: "Created for test",
			IsDone:      false,
		}

		// Unsaved todo ID should be -1.
		if todo.ID != want.ID {
			t.Errorf("Expected %s, got %s", want.ID, todo.ID)
		}
		if todo.Title != want.Title {
			t.Errorf("Expected %s, got %s", want.Title, todo.Title)
		}
		if todo.Description != want.Description {
			t.Errorf("Expected %s, got %s", want.Description, todo.Description)
		}
		// IsDone should default to false.
		if todo.IsDone != want.IsDone {
			t.Errorf("Expected %t, got %t", want.IsDone, todo.IsDone)
		}
		if todo.CreatedAt.IsZero() {
			t.Errorf("Expected non-zero time, got zero")
		}
		if todo.UpdatedAt.IsZero() {
			t.Errorf("Expected non-zero time, got zero")
		}
	})

	t.Run("NewTodoItemFromStrings", func(t *testing.T) {
		todo, err := NewTodoItemFromStrings("1", "Task 1", "Created for test", "false", "2021-09-01T00:00:00Z", "2021-09-01T00:00:00Z")
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
		want := TodoItem{
			ID:          "1",
			Title:       "Task 1",
			Description: "Created for test",
			IsDone:      false,
		}
		if todo.ID != want.ID {
			t.Errorf("Expected %s, got %s", want.ID, todo.ID)
		}
		if todo.Title != want.Title {
			t.Errorf("Expected %s, got %s", want.Title, todo.Title)
		}
		if todo.Description != want.Description {
			t.Errorf("Expected %s, got %s", want.Description, todo.Description)
		}
		if todo.IsDone != want.IsDone {
			t.Errorf("Expected %t, got %t", want.IsDone, todo.IsDone)
		}
		if todo.CreatedAt.IsZero() {
			t.Errorf("Expected non-zero time, got zero")
		}
		if todo.UpdatedAt.IsZero() {
			t.Errorf("Expected non-zero time, got zero")
		}
	})

	t.Run("Mark done", func(t *testing.T) {
		todo := NewTodoItem("Task 1", "Created for test")
		todo.Done()
		if todo.IsDone != true {
			t.Errorf("Expected true, got %t", todo.IsDone)
		}
	})

	t.Run("Mark undone", func(t *testing.T) {
		todo := NewTodoItem("Task 1", "Created for test")
		todo.Done()
		if todo.IsDone != true {
			t.Errorf("Expected true, got %t", todo.IsDone)
		}
		todo.Undone()
		if todo.IsDone != false {
			t.Errorf("Expected false, got %t", todo.IsDone)
		}
	})
}
