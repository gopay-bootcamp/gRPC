package main

import (
	"context"
	"fmt"
	"gRPC/example.com/Model"
	proto "gRPC/example.com/services"
	"google.golang.org/grpc"
	"net"
	"gRPC/example.com/bookUtil"
)

type server struct {
}



func main()  {


	listener, err:=net.Listen("tcp",":8000")

	srvr := grpc.NewServer()

	proto.RegisterBooksServicesServer(srvr,&server{})

	if err == nil {
		fmt.Println("Server running successfully....")
		srvr.Serve(listener)
	}
}


func (s *server) AddBook(ctx context.Context, request *proto.RequestForAddBook) (*proto.Response, error) {
	var book Model.Book
	book.BookName=request.BookName
	book.AuthorName=request.AuthorName

	newBook:=bookUtil.AddBookUtil(book)
	return &proto.Response{ID: newBook.ID, BookName: newBook.BookName, AuthorName: newBook.AuthorName}, nil
}

func (s *server) GetBook(ctx context.Context, request *proto.RequestForGetBook) (*proto.Response, error) {
	requestedBookId := request.ID
	book:=bookUtil.GetBookUtil(requestedBookId)
	return &proto.Response{ID: book.ID, BookName: book.BookName, AuthorName: book.AuthorName}, nil
}

func (s *server) GetBooks(request *proto.RequestForGetBooks, stream proto.BooksServices_GetBooksServer) error {
	books:=bookUtil.GetBooksUtil()
	for _, book := range books {
		stream.Send(&proto.Response{ID: book.ID, BookName: book.BookName, AuthorName: book.AuthorName})
	}
	return nil
}

func (s *server) UpdateBook(ctx context.Context, request *proto.RequestForUpdateBook) (*proto.UpdateResponse, error) {


			var book Model.Book
			book.BookName = request.BookName
			book.AuthorName = request.AuthorName

			updateBookResponse:=bookUtil.UpdateBookUtil(request.ID,book)
			if updateBookResponse==true {
				return &proto.UpdateResponse{Flag: 1},nil
			} else {
				return &proto.UpdateResponse{Flag: 0},nil
			}

}

func (s *server) DeleteBook(ctx context.Context, request *proto.RequestForDeleteBook) (*proto.DeleteResponse, error) {

	deleteBookResponse:=bookUtil.DeleteBookUtil(request.ID)
	if deleteBookResponse==true {
		return &proto.DeleteResponse{Flag: 1},nil
	} else {
		return &proto.DeleteResponse{Flag: 0},nil
	}
}









