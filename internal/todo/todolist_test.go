package todo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewTodoList(t *testing.T) {
	todoList, err := NewTodoList(NewTodoListCsvStore("todos.csv"))
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	if len(todoList.List()) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(todoList.List()))
	}

	makeTodo := func(i int) TodoItem {
		return TodoItem{
			ID:          strconv.Itoa(i),
			Title:       fmt.Sprintf("Task %d", i),
			Description: "Created for test",
		}
	}

	// Add some todo items
	count := 5
	for i := 0; i < count; i++ {
		todo := makeTodo(i)
		todoList.Add(todo)
	}

	// Check length of todo list
	todos := todoList.List()
	if len(todos) != count {
		t.Errorf("Expected %d, got %d", count, len(todos))
	}

	for i, todo := range todos {
		want := makeTodo(i)
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
	}
}

func TestTodoListGet(t *testing.T) {
	makeTodo := func(i int) TodoItem {
		return TodoItem{
			ID:          strconv.Itoa(i),
			Title:       fmt.Sprintf("Task %d", i),
			Description: "Created for test",
		}
	}

	todoList, err := NewTodoList(NewTodoListCsvStore("todos.csv"))
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}
	count := 3
	for i := 0; i < count; i++ {
		todo := makeTodo(i)
		todoList.Add(todo)
	}

	id := 2
	todo := todoList.Get(strconv.Itoa(id))
	want := makeTodo(id)
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
}

// func TestTodoListUpdate(t *testing.T) {
// 	todoList := NewTodoList()
// 	todo := NewTodoItem("Task 1", "Created for test")
// 	todoList.Add(todo)
// 	todo.Title = "Task 2"
// 	// Ensure todo is copied by value, not reference.
// 	if todoList.Todos[0].Title == "Task 2" {
// 		t.Errorf("Expected Task 1, got %s", todoList.Todos[0].Title)
// 	}
// 	todoList.Update(todo)
// 	// todoList should be updated.
// 	if todoList.Todos[0].Title != "Task 2" {
// 		t.Errorf("Expected Task 2, got %s", todoList.Todos[0].Title)
// 	}
// 	todo.Done()
// 	todoList.Update(todo)
// 	if todoList.Todos[0].IsDone != true {
// 		t.Errorf("Expected true, got %t", todo.IsDone)
// 	}
// }
//
// func TestTodoDone(t *testing.T) {
// 	todo := NewTodoItem("Task 1", "Created for test")
// 	todo.Done()
// 	if todo.IsDone != true {
// 		t.Errorf("Expected true, got %t", todo.IsDone)
// 	}
// 	todo.Undone()
// 	if todo.IsDone != false {
// 		t.Errorf("Expected false, got %t", todo.IsDone)
// 	}
// }
//
// func TestTodoDelete(t *testing.T) {
// 	todoList := NewTodoList()
// 	origin_len := 5
// 	for i := 0; i < origin_len; i++ {
// 		todo := NewTodoItem(fmt.Sprintf("Task %d", i), fmt.Sprintf("Desc %d", i))
// 		todo.ID = i
// 		todoList.Add(todo)
// 	}
//
// 	id := 2
// 	todoList.Delete(id)
// 	new_len := origin_len - 1
// 	if len(todoList.Todos) != new_len {
// 		t.Errorf("Expected %d, got %d", new_len, len(todoList.Todos))
// 	}
// 	for _, todo := range todoList.Todos {
// 		if todo.ID == id {
// 			t.Errorf("ID %d should be deleted", id)
// 		}
// 	}
// }
