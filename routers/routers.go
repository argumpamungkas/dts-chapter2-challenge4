package routers

import (
	"DTS/Chapter-2/chapter2-challenge-sesi-4/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	route.GET("/books", controllers.GetAllBook)

	route.GET("/books/:bookID", controllers.GetBookById)

	route.POST("/books", controllers.CreateBook)

	route.PUT("/books/:bookID", controllers.UpdateBook)

	route.DELETE("/books/:bookID", controllers.DeleteBook)

	return route
}
