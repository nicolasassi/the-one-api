package book

import "context"

type Book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type Repository interface {
	Get(ctx context.Context, id string) (Book, error)
}
