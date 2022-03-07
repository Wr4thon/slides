package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/Wr4thon/grpc_talk/src/live/pkg/services/todo"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		print("you have to choose an subCommand")
		return
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	todoClient := todo.NewTodoClient(conn)
	subCommand := flag.Arg(0)

	switch subCommand {
	case "add":
		id, err := todoClient.Add(context.Background(), &todo.Item{
			Text: strings.Join(flag.Args()[1:], " "),
		})

		if err != nil {
			panic(err)
		}

		println(id.ID)
	case "list":
		items, err := todoClient.List(context.Background(), &todo.Void{})
		if err != nil {
			panic(err)
		}

		for _, item := range items.Element {
			done := ""
			if item.Done {
				done = "✅"
			}

			fmt.Printf("[%d] - \"%s\" %s\n", item.ID.ID, item.Text, done)
		}
	case "complete":
		id, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		_, err = todoClient.Complete(context.Background(), &todo.ID{
			ID: int64(id),
		})
		if err != nil {
			panic(err)
		}
	default:
		panic("unknown option")
	}

}
