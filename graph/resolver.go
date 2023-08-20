package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"todo-example/graph/generated"
	"todo-example/graph/model"
)

type Resolver struct{}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, title string, completed bool) (*model.Todo, error) {
	panic("not implemented")
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
	panic("not implemented")
}

// GetTodoByID is the resolver for the getTodoById field.
func (r *queryResolver) GetTodoByID(ctx context.Context, id string) (*model.Todo, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
