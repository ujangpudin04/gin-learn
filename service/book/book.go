package book

import (
	"fmt"
	"gin-learn/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookService struct{
	storage *storage.Storage
}

var Books = []Book{}

func NewBook(storage *storage.Storage) (*BookService, error) {
	return &BookService{storage: storage}, nil
}

func (b *BookService) GetBook(c *gin.Context) {
	books := b.storage.GetData()
	c.JSON(http.StatusOK,books)
}

func (b *BookService) PostBook(c *gin.Context){
	var inputBook Book

	err:=c.ShouldBindJSON(&inputBook)
	if err!=nil {
		errorMessages:=[]string{}
		for _,e:=range err.(validator.ValidationErrors){
			errorMessage:=fmt.Sprintf("Error on filed %s,condition:%s",e.Field(),e.ActualTag())
			errorMessages=append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest,gin.H{
			"errors":errorMessages,
		})
		
		return
	}

	b.storage.Store(inputBook.Title,inputBook.Description,inputBook.Price,inputBook.Image,inputBook.Qty,inputBook.IdPerson)

	c.JSON(http.StatusCreated,gin.H{
		"title":inputBook.Title,
		"description":inputBook.Description,
		"price":inputBook.Price,
		"image":inputBook.Image,
		"qty":inputBook.Qty,
		"id_person":inputBook.IdPerson,
	})

	Books = append(Books,inputBook)

}

type NotFound struct {
	Message string `json:"message"`
}

func (b *BookService)GetOneBook(c *gin.Context) {
	id:=c.Param("id")
	log.Println(id)

	idx, err := strconv.Atoi(id)

	if err!=nil{
		log.Println(err)
	}
	
	book,err := b.storage.GetOneData(idx)
	if err!=nil{
		log.Println(err)
		return
	}
	
	if book.Title == "" {
		c.JSON(http.StatusOK, NotFound{
			Message: "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, book)

}

func (b *BookService) UpdateOneBook(c *gin.Context)  {
	var updateBook Book
	err:=c.ShouldBindJSON(&updateBook)
	if err!=nil {
		errorMessages:=[]string{}
		for _,e:=range err.(validator.ValidationErrors){
			errorMessage:=fmt.Sprintf("Error on filed %s,condition:%s",e.Field(),e.ActualTag())
			errorMessages=append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest,gin.H{
			"errors":errorMessages,
		})
		
		return
	}

	book, err := b.storage.GetOneData(updateBook.ID)

	log.Println("book",book)
	if err != nil {
		log.Println("error when getting data", err.Error())
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"errors": "data not found",
		})
		return
	}

	b.storage.UpdateBook(updateBook.ID,updateBook.Title,updateBook.Description,updateBook.Price,updateBook.Image,updateBook.Qty,updateBook.IdPerson)

	log.Println(updateBook)

	c.JSON(http.StatusCreated,gin.H{
		"title":updateBook.Title,
		"description":updateBook.Description,
		"price":updateBook.Price,
		"image":updateBook.Image,
		"qty":updateBook.Qty,
		"id_person":updateBook.IdPerson,
	})
}

func (b *BookService)DeleteBook(c *gin.Context)  {
	var delete Book

	// err:=c.ShouldBindJSON(&delete)
	book, err := b.storage.GetOneData(delete.ID)
	if err != nil {
		log.Println("error when getting data", err.Error())
	}

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"errors": "data not found",
		})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"errors": "data not found",
		})
		return
	}


	book,err=b.storage.DeleteBook(book.ID)
	if err!=nil{
		log.Println(err.Error())
	}

	log.Println(book)

	c.JSON(http.StatusOK,gin.H{
		"deleted":"success",
	})
}