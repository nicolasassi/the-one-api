package application

import (
	"context"
	"github.com/nicolasassi/the-one-api/domain/entity/the_one_api/book"
)

var _ TheOneAPI = &TheOneAPIApp{}

type TheOneAPI interface {
	book.Repository
}

type TheOneAPIApp struct {
	b book.Repository
}

func (t TheOneAPIApp) Get(ctx context.Context, id string) (book.Book, error) {
	return t.b.Get(ctx, id)
}
