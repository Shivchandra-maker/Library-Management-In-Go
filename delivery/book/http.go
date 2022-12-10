package book

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"Three-Layer-Architecture/models"
	"Three-Layer-Architecture/service"
)

type Delivery struct {
	service service.Book
}

func New(book service.Book) Delivery {
	return Delivery{service: book}
}

// Post method is post details of Book
func (a Delivery) Post(w http.ResponseWriter, r *http.Request) {
	book, err := ReadReqbody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	book2, err := a.service.Post(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	result, err := json.Marshal(book2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	w.WriteHeader(http.StatusCreated)

	fmt.Println("Successfully Post data")
}

// GetAll method is get all details of Books
func (a Delivery) GetAll(w http.ResponseWriter, r *http.Request) {
	// Getting all books
	allbooks, err := a.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	// Updating status
	w.WriteHeader(http.StatusOK)

	// Encoding
	body, err := json.Marshal(allbooks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	// writing in body
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Println("Successfully get all books")
}

// Getbyid method is get the book by its id
func (a Delivery) Getbyid(w http.ResponseWriter, r *http.Request) {
	// storing id in map
	vars := mux.Vars(r)

	book, err := a.service.Getbyid(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	// Encoding
	body, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	// Updating Status
	w.WriteHeader(http.StatusOK)

	fmt.Println("Successfully Get Book")
}

// Update method is to update details of Book
func (a Delivery) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	book, err := ReadReqbody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	bk, err := a.service.Update(vars["id"], &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	result, err := json.Marshal(bk)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Successfully Update data")
}

// Delete method is to delete details of Book by its id
func (a Delivery) Delete(w http.ResponseWriter, r *http.Request) {
	// storing id in map
	vars := mux.Vars(r)

	_, err := a.service.Delete(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	w.WriteHeader(http.StatusNoContent)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	fmt.Println("Successfully Deleted..!!")
}

func writeError(err error, w http.ResponseWriter) {
	_, errs := w.Write([]byte(err.Error()))
	if errs != nil {
		log.Printf("%v", errs)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
}

func ReadReqbody(r *http.Request) (models.Book, error) {
	// Reading body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return models.Book{}, err
	}

	var book models.Book

	// Decoding
	err = json.Unmarshal(body, &book)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}
