package main

import (
	"context"
	"fmt"
	proto "gRPC/example.com/services"
	"google.golang.org/grpc"
	"io"
)

type GrpcClient interface {
	GetBooks() ()
	GetBook(id int64) ()
	AddBook(bookName string, AuthorName string) ()
	UpdateBook(id int64,bookName string, AuthorName string) ()
	DeleteBook(id int64) ()

}

type client struct {
	client proto.BooksServicesClient
}


func newClient() GrpcClient {
	conn, _ := grpc.Dial("localhost:8000",grpc.WithInsecure())
	grpcClient := proto.NewBooksServicesClient(conn)
	return &client{
		client: grpcClient,
	}
}

func (c *client) GetBooks() {
	request:=proto.RequestForGetBooks{}
	responseStream, error := c.client.GetBooks(context.Background(), &request)
 	if error != nil {
 		fmt.Println(error.Error())
 		return
	}
	for {
		book, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("ID = %v, Book Name = %v, Author = %v\n",book.ID,book.BookName,book.AuthorName)
	}

}
func (c *client) AddBook(bookName string, authorName string) {
	request := proto.RequestForAddBook{BookName: bookName, AuthorName: authorName}
	response, error := c.client.AddBook(context.Background(), &request)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Printf("Book added with ID = %v, Book Name = %v, Author = %v\n", response.ID, response.BookName, response.AuthorName)
}
func (c *client) GetBook( id int64) {
	request := proto.RequestForGetBook{ID: id}
	response, error := c.client.GetBook(context.Background(), &request)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Printf("ID = %v, Book Name = %v, Author = %v\n", response.ID, response.BookName, response.AuthorName)
}
func (c *client) UpdateBook( id int64, bookName string, authorName string) {
	request := proto.RequestForUpdateBook{ID: id, BookName: bookName, AuthorName: authorName}
	response, error := c.client.UpdateBook(context.Background(), &request)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	if response.Flag == 1 {
		fmt.Printf("Book updated successfully\n")
	} else {
		fmt.Printf("book with given id not exists\n")
	}
}
func (c *client) DeleteBook( id int64) {
	request := proto.RequestForDeleteBook{ID: id}
	response, error := c.client.DeleteBook(context.Background(), &request)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	if response.Flag == 1 {
		fmt.Printf("Book deleted with id = %v\n",id)
	} else {
		fmt.Printf("book with given id not exists\n")
	}

}


func main() {

	client:=newClient()

	client.AddBook( "My first book", "abc")
	client.AddBook( "My second book", "abc")
	client.GetBooks()
	client.GetBook( 2)
	client.GetBook( 1)
	client.UpdateBook( 2, "My second book(updated)", "abc(xyz)")
	client.GetBook( 2)
	client.DeleteBook( 1)
	client.GetBooks( )
	client.DeleteBook( 1)

}
