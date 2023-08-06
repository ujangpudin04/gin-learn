package service

import (
	"gin-learn/service/book"
	"gin-learn/service/people"
	"gin-learn/service/upload"
)

type Services struct {
	BookService *book.BookService
	PersonService *people.PersonService
	UploadService *upload.UploadService

}

func NewServices(bookService *book.BookService,personService *people.PersonService, uploadService *upload.UploadService) (*Services, error) {
	return &Services{
		BookService: bookService,
		PersonService: personService,
		UploadService: uploadService,
	}, nil
}

