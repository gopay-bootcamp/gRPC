package main

import (
	"context"
	proto "gRPC/example.com/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"io"
	"log"
	"net"

	"testing"
)

const bufSize = 1024 * 1024
var lis *bufconn.Listener


func TestFunc(t *testing.T) {
	createFakeServer()
	client:=createClient(t)
	testGetBooks(t, client)
	testCreateBook(t, client)
	testGetBook(t, client)
	testUpdateBook(t, client)
	testDeleteBook(t, client)

}

func createFakeServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	proto.RegisterBooksServicesServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}
func createClient(t *testing.T) proto.BooksServicesClient {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	client := proto.NewBooksServicesClient(conn)
	return client
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}


func testGetBooks(t *testing.T, client proto.BooksServicesClient) {
	books = append(books, book{ID: 1, BookName: "My first book", AuthorName: "abc"})
	books = append(books, book{ID: 2, BookName: "My second book", AuthorName: "abc"})

	request:=proto.RequestForGetBooks{}
	responseStream,err:=client.GetBooks(context.Background(),&request)
	if err!=nil {
		t.Errorf("getBooks failed")
	}
	book1,err1:=responseStream.Recv()
	book2,err2:=responseStream.Recv()
	_,err3:=responseStream.Recv()
	if err1!=nil || err2!=nil {
		t.Errorf("GetBooks failed")
	}
	if book1.ID!=1 || book1.BookName!="My first book" || book1.AuthorName!="abc" {
		t.Errorf("GetBooks failed")
	}
	if book2.ID!=2 || book2.BookName!="My second book" || book2.AuthorName!="abc" {
		t.Errorf("GetBooks failed")
	}
	if err3!=io.EOF {
		t.Errorf("GetBooks failed")
	}
	books = nil
}

func testCreateBook(t *testing.T, client proto.BooksServicesClient) {
	request:=proto.RequestForAddBook{BookName: "My first book",AuthorName: "abc"}
	response,err:=client.AddBook(context.Background(),&request)
	if err!=nil {
		t.Errorf("CreateBook failed")
	}
	if response.ID!=1 {
		t.Errorf("CreateBook failed")
	}
	if books[0].ID!=1 || books[0].BookName!="My first book" || books[0].AuthorName!="abc"  {
		t.Errorf("CreateBook failed")
	}
	books = nil
}

func testGetBook(t *testing.T, client proto.BooksServicesClient) {
	books = append(books, book{ID: 1, BookName: "My first book", AuthorName: "abc"})
	books = append(books, book{ID: 2, BookName: "My second book", AuthorName: "xyz"})

	request:=proto.RequestForGetBook{ID: 1}
	response,err:=client.GetBook(context.Background(),&request)
	if err!=nil {
		t.Errorf("GetBook failed")
	}
	if response.ID!=1 || response.BookName!="My first book" || response.AuthorName!="abc" {
		t.Errorf("GetBook failed")
	}

	request=proto.RequestForGetBook{ID: 2}
	response,err=client.GetBook(context.Background(),&request)
	if err!=nil {
		t.Errorf("GetBook failed")
	}
	if response.ID!=2 || response.BookName!="My second book" || response.AuthorName!="xyz" {
		t.Errorf("GetBook failed")
	}
	books=nil
}

func testUpdateBook(t *testing.T, client proto.BooksServicesClient) {
	books = append(books, book{ID: 1, BookName: "My first book", AuthorName: "abc"})
	books = append(books, book{ID: 2, BookName: "My second book", AuthorName: "abc"})

	request:=proto.RequestForUpdateBook{ID: 2,BookName: "My updated second book",AuthorName: "abc(xyz)"}
	response,err:=client.UpdateBook(context.Background(),&request)
	if err!=nil || response==nil{
		t.Errorf("UpdateBook failed")
	}
	if books[0].ID!=1 || books[0].BookName!="My first book" || books[0].AuthorName!="abc"  {
		t.Errorf("UpdateBook failed")
	}
	if books[1].ID!=2 || books[1].BookName!="My updated second book" || books[1].AuthorName!="abc(xyz)"  {
		t.Errorf("UpdateBook failed")
	}
	books=nil
}

func testDeleteBook(t *testing.T, client proto.BooksServicesClient) {
	books = append(books, book{ID: 1, BookName: "My first book", AuthorName: "abc"})
	books = append(books, book{ID: 2, BookName: "My second book", AuthorName: "abc"})
	books = append(books, book{ID: 3, BookName: "My third book", AuthorName: "abc"})

	request:=proto.RequestForDeleteBook{ID: 2}
	response,err:=client.DeleteBook(context.Background(),&request)
	if err!=nil {
		t.Errorf("DeleteBook failed")
	}
	if response.Flag!=1 {
		t.Errorf("DeleteBook failed")
	}
	if books[0].ID!=1 || books[0].BookName!="My first book" || books[0].AuthorName!="abc"  {
		t.Errorf("DeleteBook failed")
	}
	if books[1].ID!=3 || books[1].BookName!="My third book" || books[1].AuthorName!="abc"  {
		t.Errorf("DeleteBook failed")
	}
	books=nil

}




