package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NameBook  string    `json:"name_book" gorm:"not null;type:varchar(200)"`
	Author    string    `json:"author" gorm:"not null;type:varchar(200)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return "tb_book"
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {

	if len(b.NameBook) <= 2 {
		err = errors.New("name book is too short")
	}

	if len(b.Author) <= 2 {
		err = errors.New("author is too short")
	}

	return
}
