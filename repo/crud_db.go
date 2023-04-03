package repo

import (
	"DTS/Chapter-2/chapter2-challenge-sesi-4/models"
	"errors"
	"log"

	"gorm.io/gorm"
)

func GetAllBookDB(book []models.Book) (allBook []models.Book, err error) {
	db := GetDB()

	err = db.Find(&book).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book data not found")
		err = errors.New("error getting data")
		return
	}

	allBook = append(allBook, book...)

	return
}

func GetBookByIdDB(id int, book models.Book) (bookData models.Book, err error) {
	db := GetDB()

	err = db.First(&book, "id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	bookData = book
	return
}

func CreateBookDB(book models.Book) (bookData models.Book, err error) {
	var lastID int
	var lastBook models.Book
	db := GetDB()

	_ = db.Select("id").Last(&lastBook).Scan(&lastID)

	book = models.Book{
		ID:       uint(lastID) + 1,
		NameBook: book.NameBook,
		Author:   book.Author,
	}

	err = db.Create(&book).Error
	if err != nil {
		return
	}

	bookData = book
	return
}

func UpdateBookDB(id int, book models.Book) (bookData models.Book, err error) {
	db := GetDB()
	var findBook models.Book

	err = db.Where("id = ?", id).First(&findBook).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book not found")
		err = errors.New("book not found")
		return
	}

	err = db.Model(&book).Where("id = ?", id).Updates(models.Book{
		ID:        uint(id),
		NameBook:  book.NameBook,
		Author:    book.Author,
		CreatedAt: findBook.CreatedAt,
	}).Error

	if err != nil {
		log.Println("Error updating book data", err)
		err = errors.New("error updating book data")
		return
	}

	bookData = book

	return
}

func DeleteBookDB(id int) (err error) {
	db := GetDB()
	var book models.Book

	err = db.Where("id = ?", id).First(&book).Delete(&book).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return
}
