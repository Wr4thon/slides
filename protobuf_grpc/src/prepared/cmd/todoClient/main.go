package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Wr4thon/grpc_talk/src/prepared/pkg/services/todo"
	"google.golang.org/grpc"
)

const (
	cmdAdd      string = "add"
	cmdComplete string = "complete"
	cmdList     string = "list"
)

var (
	subCommands []string = []string{
		cmdAdd,
		cmdComplete,
		cmdList,
	}
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "missing subcommand: %s\n", strings.Join(subCommands, ", "))
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while trying to establish a connection: ", err)
	}

	todoClient := todo.NewTodoClient(conn)
	subCommand := flag.Arg(0)

	switch subCommand {
	case cmdAdd:
		err = add(context.Background(), todoClient, flag.Args()[1:])
	case cmdComplete:
		err = complete(context.Background(), todoClient, flag.Args()[1:])
	case cmdList:
		err = list(context.Background(), todoClient, flag.Args()[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command, choose one of: %s\n", strings.Join(subCommands, ", "))
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error while executing %s : %v\n", subCommand, err)
		os.Exit(1)
	}
}

func add(ctx context.Context, todoClient todo.TodoClient, args []string) error {
	id, err := todoClient.Add(ctx, &todo.Item{
		Text: strings.Join(args, " "),
	})
	if err != nil {
		return err
	}

	if id.Value < 1 {
		return errors.New("could not create Item")
	}

	fmt.Println("added new Item with ID: ", id.Value)

	return nil
}

func complete(ctx context.Context, todoClient todo.TodoClient, args []string) error {
	if len(args) != 1 {
		return errors.New("invalid argument count")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	_, err = todoClient.Complete(ctx, &todo.ID{
		Value: int64(id),
	})

	if err != nil {
		return err
	}

	fmt.Println("completed Item with ID: ", id)

	return nil
}

func list(ctx context.Context, todoClient todo.TodoClient, args []string) error {
	items, err := todoClient.List(ctx, &todo.Void{})
	if err != nil {
		return err
	}

	sort.Slice(items.Elements, func(i, j int) bool {
		return items.Elements[i].ID.Value < items.Elements[j].ID.Value
	})

	for _, item := range items.Elements {
		done := ""

		if item.Done {
			done = "✅"
		}

		fmt.Printf("[%d] - \"%s\" %s\n", item.ID.Value, item.Text, done)
	}

	return nil
}
