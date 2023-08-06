package people

import (
	"fmt"
	"gin-learn/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PersonService struct{
	storage *storage.Storage
}

var Persons = []Person{}
var PersonsBook = []PersonBook{}


func NewPerson(storage *storage.Storage) (*PersonService, error) {
	return &PersonService{storage:storage}, nil
}

func (p *PersonService) PostPerson(c *gin.Context){
	var inputPerson Person

	err:=c.ShouldBindJSON(&inputPerson)
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

	p.storage.StorePerson(inputPerson.Name,inputPerson.Email,inputPerson.Alamat)

	c.JSON(http.StatusCreated,gin.H{
		"name":inputPerson.Name,
		"email":inputPerson.Email,
		"alamat":inputPerson.Alamat,
	})

	Persons = append(Persons,inputPerson)

}

func (p *PersonService) GetPerson(c *gin.Context) {
	persons := p.storage.GetPerson()
	c.JSON(http.StatusOK,persons)
}

func (p *PersonService) PostPersonBook(c *gin.Context){
	var inputPersonBook PersonBook

	err:=c.ShouldBindJSON(&inputPersonBook)
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

	p.storage.PostPersonBook(inputPersonBook.IdBook,inputPersonBook.IdPerson)

	c.JSON(http.StatusCreated,gin.H{
		"id_person":inputPersonBook.IdPerson,
		"id_book":inputPersonBook.IdBook,
	})

	PersonsBook = append(PersonsBook,inputPersonBook)

}
