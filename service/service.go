package service

import (
	"gin-learn/service/book"
	"gin-learn/service/people"
)

type Services struct {
	BookService *book.BookService
	PersonService *people.PersonService

}

func NewServices(bookService *book.BookService,personService *people.PersonService) (*Services, error) {
	return &Services{
		BookService: bookService,
		PersonService: personService,
	}, nil
}

