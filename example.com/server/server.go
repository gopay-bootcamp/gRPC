package main

import (
	"context"
	"fmt"
	proto "gRPC/example.com/services"
	"google.golang.org/grpc"
	"net"
)

type server struct {
}

type book struct {
	ID int64
	BookName string
	AuthorName string
}

var books []book

var id int64=1
func main()  {
	listener,err:=net.Listen("tcp",":8000")

	srvr := grpc.NewServer()

	proto.RegisterBooksServicesServer(srvr,&server{})

	if err == nil {
		fmt.Println("Server running successfully....")
		srvr.Serve(listener)
	}
}


func (s *server) AddBook(ctx context.Context,request *proto.RequestForAddBook) (*proto.Response,error) {
	var book book
	book.ID=id
	book.BookName=request.BookName
	book.AuthorName=request.AuthorName
	books=append(books,book)
	id=id+1
	return &proto.Response{ID: book.ID, BookName: book.BookName, AuthorName: book.AuthorName}, nil
}

func (s *server) GetBook(ctx context.Context,request *proto.RequestForGetBook) (*proto.Response,error) {
	requestedBookId:=request.ID
	for _, item := range books {
		if item.ID == requestedBookId {

			return &proto.Response{ID: item.ID,BookName: item.BookName,AuthorName: item.AuthorName},nil
		}
	}
	return &proto.Response{ID: 0, BookName: "", AuthorName: ""}, nil
}
func (s *server) GetBooks(request *proto.RequestForGetBooks,stream proto.BooksServices_GetBooksServer) error {
	for _,book := range books {
		stream.Send(&proto.Response{ID: book.ID,BookName: book.BookName,AuthorName: book.AuthorName})
	}
	return nil
}

func (s *server) UpdateBook(ctx context.Context,request *proto.RequestForUpdateBook) (*proto.Response,error) {
	requestedUpdateId:=request.ID
	for index, item := range books {
		if item.ID == requestedUpdateId {
			books = append(books[:index], books[index+1:]...)
			var book book
			book.ID = requestedUpdateId
			book.BookName=request.BookName
			book.AuthorName=request.AuthorName
			books = append(books, book)
			return &proto.Response{ID: book.ID,BookName: book.BookName,AuthorName: book.AuthorName},nil
		}
	}
	return &proto.Response{ID: 0, BookName: "", AuthorName: ""}, nil
}

func (s *server) DeleteBook(ctx context.Context,request *proto.RequestForDeleteBook) (*proto.DeleteResponse,error) {
	for index, item := range books {
		if item.ID == request.ID {
			books = append(books[:index], books[index+1:]...)
			return &proto.DeleteResponse{Flag: 1},nil
		}
	}
	return &proto.DeleteResponse{Flag: 0}, nil
}









