package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	datastoreauthor "Three-Layer-Architecture/datastore/author"
	datastorebook "Three-Layer-Architecture/datastore/book"
	deliveryauthor "Three-Layer-Architecture/delivery/author"
	deliverybook "Three-Layer-Architecture/delivery/book"
	"Three-Layer-Architecture/driver"
	serviceauthor "Three-Layer-Architecture/service/author"
	servicebook "Three-Layer-Architecture/service/book"
)

func main() {
	db, err := driver.ConnectDB()
	if err != nil {
		log.Println("could not connect to sql, err:", err)

		return
	}

	authorDatastore := datastoreauthor.New(db)
	authorService := serviceauthor.New(authorDatastore)
	authorHandler := deliveryauthor.New(authorService)

	bookDatastore := datastorebook.New(db)
	bookService := servicebook.New(bookDatastore)
	bookHandler := deliverybook.New(bookService)

	r := mux.NewRouter()

	// Author endpoints
	r.HandleFunc("/author", authorHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/author/{id}", authorHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/author/{id}", authorHandler.Delete).Methods(http.MethodDelete)

	// Book endpoints
	r.HandleFunc("/books", bookHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/book/{id}", bookHandler.Getbyid).Methods(http.MethodGet)
	r.HandleFunc("/book", bookHandler.Post).Methods(http.MethodPost)
	r.HandleFunc("/book/{id}", bookHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/book/{id}", bookHandler.Delete).Methods(http.MethodDelete)

	fmt.Println("Server Started And Listening..!!")
	log.Fatal(http.ListenAndServe(":8000", r))
}
