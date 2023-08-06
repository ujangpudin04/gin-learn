package main

import (
	"log"
	"os"

	"gin-learn/service"
	"gin-learn/service/book"
	"gin-learn/service/people"
	"gin-learn/service/upload"
	"gin-learn/storage"

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
	storage := storage.NewStorage()

	bookService, err := book.NewBook(storage)
	if err != nil {
		log.Println("Err initializing book service", err)
	}
	
	personService,err := people.NewPerson(storage)
	if err != nil {
		log.Println("Err initializing person service", err)
	}

	uploadService,err :=upload.NewUpload(storage)
	if err != nil {
		log.Println("Err initializing upload service", err)
	}
	
	services, err := service.NewServices(bookService,personService,uploadService)
	if err != nil {
		log.Println("Err initializing services", err)
	}
	
	bookRoute:=r.Group("book")
	bookRoute.GET("/get", services.BookService.GetBook)
	bookRoute.GET("/getone/:id", services.BookService.GetOneBook)
	bookRoute.POST("/post", services.BookService.PostBook)
	bookRoute.PUT("/update", services.BookService.UpdateOneBook)
	bookRoute.DELETE("/delete", services.BookService.DeleteBook)
	
	personRoute:=r.Group("person")
	personRoute.POST("/post", services.PersonService.PostPerson)
	personRoute.POST("/postbooks", services.PersonService.PostPersonBook)
	personRoute.GET("/get", services.PersonService.GetPerson)
	
	uploadRoute :=r.Group("upload")
	uploadRoute.POST("/post", services.UploadService.PostUpload)
	
	// running on Port 5000
	r.Run(os.Getenv("PORT"))
}

