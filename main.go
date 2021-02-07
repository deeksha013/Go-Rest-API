package main

import (
        "github.com/gin-gonic/gin"
        // "net/http"
        //"fmt"
        "project2/models"
        "project2/controllers"
      )

var router *gin.Engine

func main(){

  router = gin.Default()
  // book := models.Book{Id:5,Name:"d",Author:"swd",Cost:100}
  // fmt.Println(book)
  models.ConnectDatabase()

  // router.GET("/",controllers.)
  router.GET("/books",controllers.GetAllBooks)
  router.GET("/book/:id",controllers.GetBook)
  router.POST("/createBook",controllers.CreateBook)
  router.POST("/updateBookCost",controllers.UpdateBookCost)



router.Run()


  }
