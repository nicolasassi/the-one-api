package persistence

import (
	"context"
	"errors"
	"github.com/nicolasassi/the-one-api/domain/entity/book"
	"github.com/nicolasassi/the-one-api/domain/values"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	BookNotFound      = errors.New("book with the given ID was not found")
	NoBookMatchParams = errors.New("no book matches the given params")
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
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book.Book{}, err
	}
	var bookDoc bookDocument
	if err := bk.collection.FindOne(ctx, bson.M{"_id": primitiveID}).Decode(&bookDoc); err != nil {
		if err == mongo.ErrNoDocuments {
			return book.Book{}, BookNotFound
		}
		return book.Book{}, err
	}
	return *bookDoc.newBookEntity(), nil
}

func (bk *bookRepository) List(ctx context.Context, params values.QueryParams) ([]book.Book, error) {
	for k, v := range params {
		if k == "id" {
			primitiveID, err := primitive.ObjectIDFromHex(v.(string))
			if err != nil {
				return nil, err
			}
			delete(params, "id")
			params["_id"] = primitiveID
		}
	}
	cur, err := bk.collection.Find(ctx, params)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, NoBookMatchParams
		}
		return nil, err
	}
	var books []book.Book
	for cur.Next(ctx) {
		var bookDoc bookDocument
		if err := cur.Decode(&bookDoc); err != nil {
			return nil, err
		}
		books = append(books, *bookDoc.newBookEntity())
	}
	return books, nil
}

func (bk *bookRepository) Save(ctx context.Context, data *book.Book) (string, error) {
	doc, err := makeBookDocument(data)
	if err != nil {
		return "", err
	}
	if _, err := bk.List(ctx, map[string]interface{}{
		"name": doc.Name,
	}); err != nil && err != NoBookMatchParams {
		return "", err
	}
	result, err := bk.collection.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (bk *bookRepository) Update(ctx context.Context, id string, data *book.Book) error {
	doc, err := makeBookDocument(data)
	if err != nil {
		return err
	}
	if _, err := bk.Get(ctx, id); err != nil {
		return err
	}
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err := bk.collection.FindOneAndUpdate(ctx,
		bson.M{"_id": primitiveID},
		bson.M{"$set": bson.M{"publish.date": doc.Publish.Date}}).Err(); err != nil {
		return err
	}
	return nil
}

func (bk *bookRepository) Delete(ctx context.Context, id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err := bk.collection.FindOneAndDelete(ctx, bson.M{"_id": primitiveID}).Err(); err != nil {
		return err
	}
	return nil
}
