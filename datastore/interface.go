package datastore

import (
	"Three-Layer-Architecture/models"
)

type Book interface {
	Post(book *models.Book) (models.Book, error)
	GetAll() ([]models.Book, error)
	Getbyid(id string) (models.Book, error)
	Update(id string, book *models.Book) (models.Book, error)
	Delete(id string) (int, error)
}

type Author interface {
	Post(models.Author) (models.Author, error)
	Update(id string, author models.Author) (models.Author, error)
	Delete(id string) (int, error)
}
