package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"strconv"
	"time"
	"todo-example/graph/generated"
	"todo-example/graph/model"

	"github.com/pkg/errors"
)

type Resolver struct {
	todos []*model.Todo
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, title string, completed bool) (*model.Todo, error) {
	todo := &model.Todo{
		ID:        strconv.Itoa(len(r.todos) + 1),
		Title:     title,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, title string, completed bool) (*model.Todo, error) {
	panic("not implemented")
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic("not implemented")
}

// AllTodos is the resolver for the allTodos field.
func (r *queryResolver) AllTodos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// GetTodoByID is the resolver for the getTodoById field.
func (r *queryResolver) GetTodoByID(ctx context.Context, id string) (*model.Todo, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return r.todos[i+1], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
