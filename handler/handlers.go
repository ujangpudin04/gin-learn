package handler

// import (
// 	"gin-learn/entity"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func Test(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "success",
// 	})
// }

// func BookHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})
// }

// func QueryHandler(c *gin.Context) {
// 	title := c.Query("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"title": title,
// 	})
// }

// func PostBook(c *gin.Context)  {
// 	var b entity.Book
// 	err :=c.ShouldBindJSON(&b)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}

// 	c.JSON(http.StatusCreated,gin.H{

// 	})

// }