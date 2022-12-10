package book

import (
	"Three-Layer-Architecture/datastore"
	"Three-Layer-Architecture/models"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Service struct {
	datastore datastore.Book
}

func New(book datastore.Book) Service {
	return Service{book}
}

// Post method is to post Book details
func (a Service) Post(book *models.Book) (models.Book, error) {
	if book.BookID <= 0 {
		return models.Book{}, errors.New("invalid id")
	}

	// missing book fields
	if isBookFieldsMissing(book) {
		return models.Book{}, errors.New("missing book fields")
	}

	// missing author fields
	if isAuthorFieldsMissing(book.Auth) {
		return models.Book{}, errors.New("missing author fields")
	}

	if !isValidPublishedDate(book.PublishedDate) {
		return models.Book{}, errors.New("invalid publishedDate")
	}

	if !isValidPublication(book.Publication) {
		return models.Book{}, errors.New("invalid publication")
	}

	_, err := a.datastore.Post(book)
	if err != nil {
		return models.Book{}, err
	}

	return *book, nil
}

// Getbyid method is to get Book details by id
func (a Service) Getbyid(id string) (models.Book, error) {
	// checking missing id
	if id == "" {
		return models.Book{}, errors.New("missing id")
	}

	// converting string to integer to check for invalid id
	iD, err := strconv.Atoi(id)
	if err != nil {
		return models.Book{}, err
	}

	// Checking invalid id
	if iD <= 0 {
		return models.Book{}, errors.New("invalid id")
	}

	book, err := a.datastore.Getbyid(id)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

// Update method is to update Book details
func (a Service) Update(id string, book *models.Book) (models.Book, error) {
	// checking missing id
	if id == "" {
		return models.Book{}, errors.New("missing id")
	}

	// missing author fields
	if isAuthorFieldsMissing(book.Auth) {
		return models.Book{}, errors.New("missing author fields")
	}

	// missing book fields
	if isBookFieldsMissing(book) {
		return models.Book{}, errors.New("missing book fields")
	}

	if !isValidPublishedDate(book.PublishedDate) {
		return models.Book{}, errors.New("invalid publishedDate")
	}

	if !isValidPublication(book.Publication) {
		return models.Book{}, errors.New("invalid publication")
	}

	// converting string to integer to check for invalid id
	iD, _ := strconv.Atoi(id)
	if iD <= 0 {
		return models.Book{}, fmt.Errorf("invalid id")
	}

	bk, err := a.datastore.Update(id, book)
	if err != nil {
		return models.Book{}, err
	}

	return bk, nil
}

// Delete method is to delete Book details
func (a Service) Delete(id string) (int, error) {
	// Checking missing id
	if id == "" {
		return 0, errors.New("missing id")
	}

	// converting string to integer to check for invalid id
	iD, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	// Checking invalid id
	if iD <= 0 {
		return 0, errors.New("invalid id")
	}

	rowAffected, err := a.datastore.Delete(id)
	if err != nil {
		return 0, err
	}

	return rowAffected, nil
}

// GetAll method is to details of book
func (a Service) GetAll() ([]models.Book, error) {
	// To store book details
	var book []models.Book

	var err error

	book, err = a.datastore.GetAll()
	if err != nil {
		return []models.Book{}, err
	}

	return book, nil
}

func isValidPublishedDate(date string) bool {
	split := strings.Split(date, "/")

	yearInstr := split[2]

	yearInint, err := strconv.Atoi(yearInstr)

	if err != nil {
		log.Printf("Cannot convert dob in integer : %v", yearInint)
	}

	if yearInint < 2022 && yearInint > 1880 {
		return true
	}

	return false
}

func isValidPublication(pub string) bool {
	if pub == "Scholastic" || pub == "Arihant" || pub == "Penguin" {
		return true
	}

	return false
}

func isAuthorFieldsMissing(auth models.Author) bool {
	if auth.FirstName == "" || auth.LastName == "" || auth.PenName == "" || auth.Dob == "" {
		return true
	}

	return false
}

func isBookFieldsMissing(book *models.Book) bool {
	if book.PublishedDate == "" || book.Publication == "" || book.Title == "" {
		return true
	}

	return false
}
