package the_one_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nicolasassi/the-one-api/domain/entity/the_one_api/book"
	"io/ioutil"
	"net/http"
)

const (
	bookPATH = "book/"
)

var (
	_            book.Repository = &bookRepository{}
	BookNotFound                 = errors.New("the book with given ID was not found")
)

type bookResponse struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

func (br bookResponse) newBookEntity() *book.Book {
	return &book.Book{
		ID:   br.ID,
		Name: br.Name,
	}
}

type bookRepository struct {
	client *http.Client
}

func NewBookRepository() *bookRepository {
	return &bookRepository{client: &http.Client{Timeout: Timeout}}
}

func (br bookRepository) Get(ctx context.Context, id string) (book.Book, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s%s", endpoint, bookPATH, id), nil)
	if err != nil {
		return book.Book{}, err
	}
	resp, err := br.client.Do(req)
	if err != nil {
		return book.Book{}, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return book.Book{}, err
	}
	if err := resp.Body.Close(); err != nil {
		return book.Book{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return book.Book{},
			fmt.Errorf("status %v while requesting 'the one api': %v", resp.StatusCode, string(b))
	}
	respParsed, err := newResponse(b)
	if err != nil {
		return book.Book{}, err
	}
	if len(respParsed.Docs) > 0 {
		return book.Book{}, BookNotFound
	}
	doc, err := json.Marshal(respParsed.Docs[0])
	if err != nil {
		return book.Book{}, err
	}
	var bookResp bookResponse
	if err := json.Unmarshal(doc, &bookResp); err != nil {
		return book.Book{}, err
	}
	return *bookResp.newBookEntity(), nil
}
