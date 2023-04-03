package controllers

import (
	"DTS/Chapter-2/chapter2-challenge-sesi-4/models"
	"DTS/Chapter-2/chapter2-challenge-sesi-4/repo"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	var books []models.Book

	allBook, err := repo.GetAllBookDB(books)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid Request %+v", err.Error()),
		})
		return
	}

	if allBook == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": allBook,
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var book models.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert id")
		return
	}

	bookData, err := repo.GetBookByIdDB(idBook, book)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Data with id %d not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)

}

func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookData, err := repo.CreateBookDB(book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint("Invalid create data book because ", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, bookData)
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var book models.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Error convert id book")
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookData, err := repo.UpdateBookDB(idBook, book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid update data because %+v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert id")
		return
	}

	err = repo.DeleteBookDB(idBook)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %d successfully deleted", idBook),
	})
}
