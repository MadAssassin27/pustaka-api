package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"pustaka-api/book"
	"pustaka-api/handler"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234567890@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	router.SetTrustedProxies([]string{"192.168.1.2"})

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run()
}

//main
//handler
//service
//repository
//db
//mysql
