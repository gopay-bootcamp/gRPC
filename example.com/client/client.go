package main

import (
	"context"
	"fmt"
	proto "gRPC/example.com/services"
	"google.golang.org/grpc"
	"io"
)
func main() {
	conn,_ :=grpc.Dial("localhost:8000",grpc.WithInsecure())
	client:=proto.NewBooksServicesClient(conn)
	addBook(client,"My first book","abc")
	addBook(client,"My second book","abc")
	getBooks(client)
	getBook(client,2)
	getBook(client, 1)
	getBook(client,3)
	updateBook(client,2,"My second book(updated)","abc(xyz)")
	getBook(client,2)
	deleteBook(client,1)
	getBook(client,1)
	deleteBook(client,1)

}

func getBooks(client proto.BooksServicesClient) {
	request:=proto.RequestForGetBooks{}
	responseStream,_:=client.GetBooks(context.Background(),&request)

	for {
		book,err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("ID = %v, Book Name = %v, Author = %v\n",book.ID,book.BookName,book.AuthorName)
	}

}
func addBook(client proto.BooksServicesClient,bookName string,authorName string) {
	request:=proto.RequestForAddBook{BookName: bookName,AuthorName: authorName}
	response,_:=client.AddBook(context.Background(),&request)
	fmt.Printf("Book added with ID = %v, Book Name = %v, Author = %v\n", response.ID, response.BookName, response.AuthorName)
}
func getBook(client proto.BooksServicesClient,id int64) {
	request:=proto.RequestForGetBook{ID: id}
	response,_:=client.GetBook(context.Background(),&request)
	fmt.Printf("ID = %v, Book Name = %v, Author = %v\n", response.ID, response.BookName, response.AuthorName)
}
func updateBook(client proto.BooksServicesClient,id int64,bookName string,authorName string) {
	request:=proto.RequestForUpdateBook{ID:id,BookName: bookName,AuthorName: authorName}
	response,_:=client.UpdateBook(context.Background(),&request)
	fmt.Printf("Book updated with ID = %v, Book Name = %v, Author = %v\n", response.ID, response.BookName, response.AuthorName)
}
func deleteBook(client proto.BooksServicesClient,id int64) {
	request := proto.RequestForDeleteBook{ID: id}
	response, _ := client.DeleteBook(context.Background(), &request)
	if response.Flag==1 {
		fmt.Printf("Book deleted with id = %v\n",id)
	} else {
		fmt.Printf("book with given id not exists\n")
	}

}