package main

import (
	"context"
	"log"
	"net"

	"github.com/Wr4thon/grpc_talk/src/prepared/pkg/services/todo"
	"github.com/pkg/errors"
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
		log.Fatalf("error while creating listener: %v", err)
	}

	if err := srv.Serve(l); err != nil {
		log.Fatalf("error while listening: %v", err)
	}
}

func (s *todoServer) Add(ctx context.Context, item *todo.Item) (*todo.ID, error) {
	id := int64(len(s.items) + 1)

	item.ID = &todo.ID{
		Value: id,
	}

	item.Done = false

	s.items[id] = item

	return item.ID, nil
}

func (s *todoServer) Complete(ctx context.Context, id *todo.ID) (*todo.Void, error) {
	var item *todo.Item
	var ok bool

	if item, ok = s.items[id.Value]; !ok {
		return &todo.Void{}, errors.Errorf("Item with ID %v not found", id.Value)
	}

	item.Done = true

	return &todo.Void{}, nil
}

func (s *todoServer) List(ctx context.Context, void *todo.Void) (*todo.Items, error) {
	items := &todo.Items{
		Elements: make([]*todo.Item, len(s.items)),
	}

	var i int
	for key := range s.items {
		items.Elements[i] = s.items[key]
		i++
	}

	return items, nil
}
