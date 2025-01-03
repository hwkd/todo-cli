package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hwkd/todo-cli/internal/args"
	"github.com/hwkd/todo-cli/internal/todo"
)

func main() {
	result, err := args.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	todoList, err := todo.NewTodoList(todo.NewTodoListCsvStore("todo.csv"))
	if err != nil {
		panic(err)
	}

	switch result.Action {
	case args.ActionHelp:
		handleHelpAction()
	case args.ActionList:
		handleListAction(*todoList)
	case args.ActionAdd:
		handleAddAction(todoList, result.ParseAddActionValues())
	case args.ActionUpdate:
		handleUpdateAction(todoList, result.ParseUpdateActionValues())
	case args.ActionDelete:
		handleDeleteAction(todoList, result.ParseDeleteActionValues())
	case args.ActionMarkComplete:
		handleMarkCompleteAction(todoList, result.ParseMarkCompleteActionValues())
	case args.ActionMarkIncomplete:
		handleMarkInompleteAction(todoList, result.ParseMarkIncompleteActionValues())
	}
}

func handleHelpAction() {
	fmt.Println("Usage:")
	fmt.Println("  todo -h")
	fmt.Println("  todo -l")
	fmt.Println("  todo -a <title> [description]")
	fmt.Println("  todo -u <id> <title> [description]")
	fmt.Println("  todo -d <id>...")
	fmt.Println("  todo -c <id>...")
	fmt.Println("  todo -r <id>...")
}

func handleListAction(todoList todo.TodoList) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tTitle\tDescription\tDone")
	fmt.Fprintln(writer, "--\t-----\t-----------\t----")
	todos := todoList.List()
	for _, todoItem := range todos {
		fmt.Fprintf(writer, "%s\t%s\t%s\t%t\n", todoItem.ID, todoItem.Title, todoItem.Description, todoItem.IsDone)
	}
	writer.Flush()
}

func handleAddAction(todoList *todo.TodoList, values args.ParsedAddActionValues) {
	todoItem := todo.NewTodoItem(values.Title, values.Description)
	todoList.Add(*todoItem)
	todoList.Flush()
}

func handleUpdateAction(todoList *todo.TodoList, values args.ParsedUpdateActionValues) {
	todo := todoList.Get(values.ID)
	if todo == nil {
		fmt.Printf("Todo not found: %s\n", values.ID)
		return
	}
	todo.Title = values.Title
	todo.Description = values.Description
	todoList.Update(*todo)
	todoList.Flush()
}

func handleDeleteAction(todoList *todo.TodoList, result args.ParsedIdValues) {
	for _, id := range result.IDs {
		todoList.Delete(id)
	}
	todoList.Flush()
}

func handleMarkCompleteAction(todoList *todo.TodoList, result args.ParsedIdValues) {
	for _, id := range result.IDs {
		todo := todoList.Get(id)
		if todo == nil {
			fmt.Printf("Todo not found: %s\n", id)
			return
		}
		todo.Done()
		todoList.Update(*todo)
	}
	todoList.Flush()
}

func handleMarkInompleteAction(todoList *todo.TodoList, result args.ParsedIdValues) {
	for _, id := range result.IDs {
		todo := todoList.Get(id)
		if todo == nil {
			fmt.Printf("Todo not found: %s\n", id)
			return
		}
		todo.Undone()
		todoList.Update(*todo)
	}
	todoList.Flush()
}
