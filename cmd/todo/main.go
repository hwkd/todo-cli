package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hwkd/todo-cli/internal/args"
	"github.com/hwkd/todo-cli/internal/todo"
)

func main() {
	if err := run(); err != nil {
		if err, ok := err.(args.ArgError); ok {
			fmt.Println(err)
			displayUsage(err.Action)
		} else {
			fmt.Println(err)
		}
		return
	}
}

func run() error {
	result, err := args.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	todoList, err := todo.NewTodoList(todo.NewTodoListCsvStore("todo.csv"))
	if err != nil {
		return err
	}

	err = nil
	switch result.Action {
	case args.ActionHelp:
		handleHelpAction()
	case args.ActionList:
		handleListAction(*todoList)
	case args.ActionAdd:
		err = handleAddAction(todoList, result.ParseAddActionValues())
	case args.ActionUpdate:
		err = handleUpdateAction(todoList, result.ParseUpdateActionValues())
	case args.ActionDelete:
		err = handleDeleteAction(todoList, result.ParseDeleteActionValues())
	case args.ActionMarkComplete:
		err = handleMarkCompleteAction(todoList, result.ParseMarkCompleteActionValues())
	case args.ActionMarkIncomplete:
		err = handleMarkInompleteAction(todoList, result.ParseMarkIncompleteActionValues())
	}
	return err
}

var actions = []struct {
	action      string
	flag        string
	params      string
	description string
}{
	{args.ActionHelp, "-h", "", ""},
	{args.ActionList, "-l", "", ""},
	{args.ActionAdd, "-a", "<title> [description]", "Add a todo item"},
	{args.ActionUpdate, "-u", "<id> [-t title] [-d description]", "Update a todo item"},
	{args.ActionDelete, "-d", "<id>...", "Delete todo items by id"},
	{args.ActionMarkComplete, "-c", "<id>...", "Mark complete by id"},
	{args.ActionMarkIncomplete, "-r", "<id>...", "Mark incomplete by id"},
}

func displayUsage(action string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprint(writer, "Usage:\n")
	for _, act := range actions {
		if act.action == action {
			fmt.Fprintf(writer, "  todo %s\t%s\t%s\n", act.flag, act.params, act.description)
		}
	}
	writer.Flush()
}

func handleHelpAction() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprint(writer, "Usage:\n")
	for _, action := range actions {
		fmt.Fprintf(writer, "  todo %s\t%s\t%s\n", action.flag, action.params, action.description)
	}
	writer.Flush()
}

func handleListAction(todoList todo.TodoList) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, "ID\tTitle\tDescription\tDone\tCreated At")
	fmt.Fprintln(writer, "--\t-----\t-----------\t----\t----------")
	todos := todoList.List()
	for _, todoItem := range todos {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\t%t\t%s\n",
			todoItem.ID,
			todoItem.Title,
			todoItem.Description,
			todoItem.IsDone,
			todoItem.CreatedAt.Format("2006-01-02 03:04:05 PM"),
		)
	}
	writer.Flush()
}

func handleAddAction(todoList *todo.TodoList, values args.ParsedAddActionValues) error {
	todoItem := todo.NewTodoItem(values.Title, values.Description)
	todoList.Add(*todoItem)
	return todoList.Flush()
}

func handleUpdateAction(todoList *todo.TodoList, values args.ParsedUpdateActionValues) error {
	todo := todoList.Get(values.ID)
	if todo == nil {
		return fmt.Errorf("Todo with %s not found\n", values.ID)
	}
	todo.Title = values.Title
	todo.Description = values.Description
	todoList.Update(*todo)
	return todoList.Flush()
}

func handleDeleteAction(todoList *todo.TodoList, result args.ParsedIdValues) error {
	for _, id := range result.IDs {
		todoList.Delete(id)
	}
	return todoList.Flush()
}

func handleMarkCompleteAction(todoList *todo.TodoList, result args.ParsedIdValues) error {
	for _, id := range result.IDs {
		todo := todoList.Get(id)
		if todo == nil {
			return fmt.Errorf("Todo with %s not found\n", id)
		}
		todo.Done()
		todoList.Update(*todo)
	}
	return todoList.Flush()
}

func handleMarkInompleteAction(todoList *todo.TodoList, result args.ParsedIdValues) error {
	for _, id := range result.IDs {
		todo := todoList.Get(id)
		if todo == nil {
			return fmt.Errorf("Todo with %s not found\n", id)
		}
		todo.Undone()
		todoList.Update(*todo)
	}
	return todoList.Flush()
}
