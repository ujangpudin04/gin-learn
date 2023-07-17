package book

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookService struct{}

var books = []Book{}

func NewBook() (*BookService, error) {
	return &BookService{}, nil
}

func (b *BookService) GetBook(c *gin.Context) {
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

	c.JSON(http.StatusCreated,gin.H{
		"id":inputBook.ID,
		"title":inputBook.Title,
		"description":inputBook.Description,
		"price":inputBook.Price,
		"image":inputBook.Image,
		"qty":inputBook.Qty,
	})

	books = append(books,inputBook)

}

func (b *BookService)GetOneBook(c *gin.Context) {
	id:=c.Param("id")
	log.Println(id)

	for _, v := range books {
		if id == string(v.ID) {
			
			c.JSON(http.StatusOK,gin.H{
				
			})
			// fmt.Printf("books: %v\n", books)
		}
	}
}