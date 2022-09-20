package database

import (
	"day4/models"
	"fmt"
)

var books = []models.Book{
	{
		Id:     1,
		Title:  "Filosofi Teras",
		Writer: "Henry Manampiring",
	},
	{
		Id:     2,
		Title:  "Rich Dad Poor Dad",
		Writer: "Robert T. Kiyosaki",
	},
}

func GetBooks() *[]models.Book {
	return &books
}
func GetBooksById(id int) *models.Book {
	for _, book := range books {
		if book.Id == id {
			return &book
		}
	}
	return &models.Book{}
}
func CreateBook(book models.Book) *models.Book {
	books = append(books, book)
	return &book
}

func UpdateBookById(id int, bookUpdate models.Book) *models.Book {
	fmt.Println(books)
	for _, book := range books {
		if book.Id == id {
			book.Title = bookUpdate.Title
			book.Writer = bookUpdate.Writer
			return &book
		}
	}
	return &models.Book{}
}

func DeleteBooksById(id int) interface{} {
	for _, book := range books {
		if book.Id == id {
			books = append(books[:id-1], books[id:]...)
		}
		return fmt.Sprint("success delete")
	}
	return fmt.Sprint("record not found")
}
