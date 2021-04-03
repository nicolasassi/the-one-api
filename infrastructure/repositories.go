package infrastructure

import (
	"github.com/nicolasassi/the-one-api/domain/entity/book"
	theOneBook "github.com/nicolasassi/the-one-api/domain/entity/the_one_api/book"
	"github.com/nicolasassi/the-one-api/infrastructure/persistence"
	"github.com/nicolasassi/the-one-api/infrastructure/the_one_api"
	"go.mongodb.org/mongo-driver/mongo"
)

const mongoCollection = "books"

type Persistence struct {
	Book    book.Repository
	mongoDB *mongo.Database
}

func (p *Persistence) setRepositories() {
	p.Book = persistence.NewBookRepository(p.mongoDB, mongoCollection)
}

type TheOneAPI struct {
	Book theOneBook.Repository
}

func (t *TheOneAPI) setRepositories() {
	t.Book = the_one_api.NewBookRepository()
}

type Repositories struct {
	Persistence Persistence
	TheOneAPI   TheOneAPI
}

func NewRepositories(opts ...Options) *Repositories {
	reps := new(Repositories)
	for _, opt := range opts {
		opt(reps)
	}
	reps.TheOneAPI.setRepositories()
	reps.Persistence.setRepositories()
	return reps
}
