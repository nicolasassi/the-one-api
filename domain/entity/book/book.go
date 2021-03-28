package book

import (
	"context"
	"github.com/nicolasassi/the-one-api/domain/values"
)

type Book struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Publish values.Publish `json:"publish"`
}

type Repository interface {
	Get(ctx context.Context, id string) (Book, error)
	List(ctx context.Context, params values.QueryParams) ([]Book, error)
	Save(ctx context.Context, data *Book) (string, error)
	Update(ctx context.Context, id string, data *Book) error
	Delete(ctx context.Context, id string) error
}
