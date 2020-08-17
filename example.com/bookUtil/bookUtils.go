package bookUtil

import (
	"gRPC/example.com/Model"
)

var Books []Model.Book
var id int64 = 1

func GetBooksUtil() []Model.Book {
	return Books
}

func GetBookUtil(id int64) Model.Book {
	for _, item := range Books {
		if item.ID == id {
			return item
		}
	}
	return Model.Book{}
}

func AddBookUtil(newBook Model.Book) Model.Book {
	newBook.ID=id
	id=id+1
	Books = append(Books, newBook)
	return newBook
}

func UpdateBookUtil(id int64,updatedBook Model.Book) bool{

	for index, item := range Books {
		if item.ID == id {
			Books = append(Books[:index], Books[index+1:]...)

			updatedBook.ID = id
			Books = append(Books, updatedBook)
			return true
		}
	}
	return false
}

func DeleteBookUtil(id int64) bool {
	for index, item := range Books {
		if item.ID == id {
			Books = append(Books[:index], Books[index+1:]...)
			return true
		}
	}
	return false
}