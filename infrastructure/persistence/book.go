package persistence

import (
	"context"
	"github.com/nicolasassi/the-one-api/domain/entity/book"
	"github.com/nicolasassi/the-one-api/domain/values"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookDocument struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Publish *values.Publish    `bson:"publish,omitempty"`
}

func (bd bookDocument) newBookEntity() *book.Book {
	return &book.Book{
		ID:      bd.ID.Hex(),
		Name:    bd.Name,
		Publish: bd.Publish,
	}
}

func makeBookDocument(b *book.Book) (bookDocument, error) {
	var id primitive.ObjectID
	if b.ID != "" {
		v, err := primitive.ObjectIDFromHex(b.ID)
		if err != nil {
			return bookDocument{}, err
		}
		id = v
	}
	return bookDocument{
		ID:      id,
		Name:    b.Name,
		Publish: b.Publish,
	}, nil
}

var _ book.Repository = &bookRepository{}

type bookRepository struct {
	db             *mongo.Database
	collectionName string
	collection     *mongo.Collection
}

func NewBookRepository(db *mongo.Database, collectionName string) *bookRepository {
	return &bookRepository{
		db:             db,
		collectionName: collectionName,
		collection:     db.Collection(collectionName),
	}
}

func (bk *bookRepository) Get(ctx context.Context, id string) (book.Book, error) {
	panic("")
}

func (bk *bookRepository) Save(ctx context.Context, data *book.Book) (string, error) {
	panic("")
}

func (bk *bookRepository) List(ctx context.Context, params values.QueryParams) ([]book.Book, error) {
	panic("implement me")
}

func (bk *bookRepository) Update(ctx context.Context, id string, data *book.Book) error {
	panic("implement me")
}

func (bk *bookRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
