package main

import (
	"context"
	"net"

	"github.com/Wr4thon/grpc_talk/src/live/pkg/services/todo"
	"google.golang.org/grpc"
)

type todoServer struct {
	items map[int64]*todo.Item
}

func main() {
	srv := grpc.NewServer()

	todoServer := &todoServer{
		items: make(map[int64]*todo.Item),
	}
	todo.RegisterTodoServer(srv, todoServer)

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	if err := srv.Serve(l); err != nil {
		panic(err)
	}
}

func (server *todoServer) Add(ctx context.Context, foo *todo.Item) (*todo.ID, error) {
	id := int64(len(server.items) + 1)

	foo.ID = &todo.ID{
		ID: id,
	}

	server.items[id] = foo

	return foo.ID, nil
}

func (server *todoServer) Complete(ctx context.Context, foo *todo.ID) (*todo.Void, error) {

	server.items[foo.ID].Done = true

	return &todo.Void{}, nil
}

func (server *todoServer) List(ctx context.Context, foo *todo.Void) (*todo.Items, error) {
	var result todo.Items
	result.Element = make([]*todo.Item, len(server.items))

	var i int
	for _, item := range server.items {
		result.Element[i] = item
		i++
	}

	return &result, nil
}
