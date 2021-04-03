package persistence

import (
	"context"
	"github.com/nicolasassi/the-one-api/domain/entity/book"
	"github.com/nicolasassi/the-one-api/domain/values"
)

var _ Book = &BookDB{}

type Book interface {
	book.Repository
}

type BookDB struct {
	b book.Repository
}

func (b BookDB) Get(ctx context.Context, id string) (book.Book, error) {
	return b.b.Get(ctx, id)
}

func (b BookDB) List(ctx context.Context, params values.QueryParams) ([]book.Book, error) {
	return b.b.List(ctx, params)
}

func (b BookDB) Save(ctx context.Context, data *book.Book) (string, error) {
	return b.b.Save(ctx, data)
}

func (b BookDB) Update(ctx context.Context, id string, data *book.Book) error {
	return b.b.Update(ctx, id, data)
}

func (b BookDB) Delete(ctx context.Context, id string) error {
	return b.b.Delete(ctx, id)
}
