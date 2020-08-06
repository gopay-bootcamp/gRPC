package bookUtil

import (
	"gRPC/example.com/Model"
)

var books []Model.Book
var id int64 = 1

func GetBooksUtil() []Model.Book {
	return books
}

func GetBookUtil(id int64) Model.Book {
	for _, item := range books {
		if item.ID == id {
			return item
		}
	}
	return Model.Book{}
}

func AddBookUtil(newBook Model.Book) Model.Book {
	newBook.ID=id
	id=id+1
	books = append(books, newBook)
	return newBook
}

func UpdateBookUtil(id int64,updatedBook Model.Book) bool{

	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)

			updatedBook.ID = id
			books = append(books, updatedBook)
			return true
		}
	}
	return false
}

func DeleteBookUtil(id int64) bool {
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			return true
		}
	}
	return false
}