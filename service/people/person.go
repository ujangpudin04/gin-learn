package people

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PersonService struct{}

func NewPerson() (*PersonService, error) {
	return &PersonService{}, nil
}

func (p *PersonService) GetPerson(c *gin.Context) {
	log.Println("get person endpoint")
}