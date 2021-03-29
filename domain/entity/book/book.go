package book

import (
	"context"
	"fmt"
	"github.com/nicolasassi/the-one-api/domain/values"
)

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Publish *values.Publish `json:"publish"`
}

func (b Book) Validate(action string) map[string]string {
	errorMessages := make(map[string]string)
	switch action {
	case "create":
		if b.Name == "" {
			errorMessages["name"] = fmt.Sprintf("name field should not be nil on %s", action)
		}
		if !b.Publish.IsValid() {
			errorMessages["publish_date"] = "invalid publish date value"
		}
	case "update":
		if !b.Publish.IsValid() {
			errorMessages["publish_date"] = "invalid publish date value"
		}
	}
	return errorMessages
}

type Repository interface {
	Get(ctx context.Context, id string) (Book, error)
	List(ctx context.Context, params values.QueryParams) ([]Book, error)
	Save(ctx context.Context, data *Book) (string, error)
	Update(ctx context.Context, id string, data *Book) error
	Delete(ctx context.Context, id string) error
}
