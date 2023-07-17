package main

import (
	"log"
	"os"

	"gin-learn/service"
	"gin-learn/service/book"
	"gin-learn/service/people"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading env file")
	}
  
	gin.SetMode(os.Getenv("GIN_MODE"))
    r:=gin.Default()
	// r.GET("/test", handler.Test)
	// r.GET("/books/:id/:title", handler.BookHandler)
	// r.GET("/query", handler.QueryHandler)
	// r.POST("/book",handler.PostBook)
	

	bookService, err := book.NewBook()
	if err != nil {
		log.Println("Err initializing book service", err)
	}
	
	
	personService,err := people.NewPerson()
	if err != nil {
		log.Println("Err initializing person service", err)
	}
	
	
	
	services, err := service.NewServices(bookService,personService)
	if err != nil {
		log.Println("Err initializing services", err)
	}
	
	bookRoute:=r.Group("book")
	bookRoute.GET("/get", services.BookService.GetBook)
	bookRoute.GET("/getone/:id", services.BookService.GetOneBook)
	bookRoute.POST("/post", services.BookService.PostBook)
	r.GET("/person/get", services.PersonService.GetPerson)
	
	// running on Port 5000
	r.Run(os.Getenv("PORT"))
}

