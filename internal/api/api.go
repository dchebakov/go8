package api

import (
	"eight/internal/domain/authors"
	"eight/internal/domain/books"
)

// API holds all the dependencies required to expose APIs. And each API is a function with *API as its receiver
type API struct {
	books   *books.HandlerBooks
	authors *authors.HandlerAuthors
}

// add microservice to the PARAM
func NewService(bs *books.HandlerBooks, as *authors.HandlerAuthors) (*API, error) {
	return &API{
		books:   bs,
		authors: as,
	}, nil
}
