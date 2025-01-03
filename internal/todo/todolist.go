package todo

type TodoList struct {
	Todos    []TodoItem
	modified bool
	store    Store
}

// NewTodoList creates a new TodoList.
func NewTodoList(store Store) (*TodoList, error) {
	list := &TodoList{
		modified: false,
		store:    store,
	}
	todos, err := list.store.Load()
	if err != nil {
		return nil, err
	}
	list.setTodos(todos)
	return list, nil
}

// List returns all TodoItems.
func (todoList *TodoList) List() []TodoItem {
	return todoList.Todos
}

// Get returns a TodoItem by ID.
func (todoList *TodoList) Get(id string) *TodoItem {
	idLen := len(id)
	for i := range todoList.Todos {
		todo := &todoList.Todos[i]
		if todo.ID[:idLen] == id {
			return todo
		}
	}
	return nil
}

// Add appends a TodoItem to the list.
func (todoList *TodoList) Add(todo TodoItem) {
	todoList.Todos = append(todoList.Todos, todo)
	todoList.modified = true
}

// Update updates a TodoItem in the list that matches the ID.
func (todoList *TodoList) Update(todo TodoItem) {
	for i, _todo := range todoList.Todos {
		if _todo.ID == todo.ID {
			todoList.Todos[i] = todo
			todoList.modified = true
			break
		}
	}
}

// Delete removes a TodoItem from the list that matches the ID.
func (todoList *TodoList) Delete(id string) {
	idLen := len(id)
	for i := range todoList.Todos {
		todo := &todoList.Todos[i]
		if todo.ID[:idLen] == id {
			todoList.Todos = append(todoList.Todos[:i], todoList.Todos[i+1:]...)
			todoList.modified = true
			break
		}
	}
}

// Flush writes to the store if the todolist was modified.
func (todoList *TodoList) Flush() error {
	if !todoList.modified {
		return nil
	}

	// Save to disk.
	todoList.store.Save(todoList.Todos)
	todoList.modified = false

	return nil
}

// setTodos sets the list of TodoItems.
func (todoList *TodoList) setTodos(todos []TodoItem) {
	todoList.Todos = todos
}
