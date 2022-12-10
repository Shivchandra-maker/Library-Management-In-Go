package book

import (
	"Three-Layer-Architecture/models"
	"database/sql"
	"log"
	"strconv"
)

type Datastore struct {
	db *sql.DB
}

func New(db *sql.DB) Datastore {
	return Datastore{db: db}
}

// Post method is to Post data in Book
func (d Datastore) Post(book *models.Book) (models.Book, error) {
	// inserting data into Db
	_, err := d.db.Exec("insert into Book(bookId,title,authorId,Publication,PublishedDate) values (?,?,?,?,?)",
		book.BookID, book.Title, book.AuthorID, book.Publication, book.PublishedDate)
	if err != nil {
		return models.Book{}, err
	}

	return *book, nil
}

// GetAll method is to get all Books with Author
func (d Datastore) GetAll() ([]models.Book, error) {
	// reading all books from Db
	allRows, err := d.db.Query("SELECT * FROM Book")
	if err != nil {
		return nil, err
	}

	if allRows.Err() != nil {
		log.Print(allRows.Err())
	}

	// Closing db.query
	defer allRows.Close()

	// To store books
	var book []models.Book

	// Iterating to each book
	for allRows.Next() {
		var b models.Book

		err = allRows.Scan(&b.BookID, &b.Title, &b.AuthorID, &b.Publication, &b.PublishedDate)
		if err != nil {
			return []models.Book{}, err
		}

		// for storing author details
		row := d.db.QueryRow("SELECT * FROM Author where authorId=?", b.AuthorID)

		var author models.Author

		if err2 := row.Scan(&author.AuthID, &author.FirstName, &author.LastName, &author.Dob, &author.PenName); err2 != nil {
			return []models.Book{}, err2
		}

		b.Auth = author

		if err != nil {
			return []models.Book{}, err
		}

		book = append(book, b)
	}

	return book, nil
}

// Getbyid method is to get book by its ID
func (d Datastore) Getbyid(iD string) (models.Book, error) {
	// converting string to integer to check for invalid id
	id, err := strconv.Atoi(iD)
	if err != nil {
		return models.Book{}, err
	}

	// reading all data of book with given id
	row := d.db.QueryRow("select * from Book where bookId=?", id)

	// to store d book
	var book models.Book

	// fetching data of book at given id and storing in book
	if err := row.Scan(&book.BookID, &book.Title, &book.AuthorID, &book.Publication, &book.PublishedDate); err != nil {
		return models.Book{}, err
	}

	// for storing author details
	result := d.db.QueryRow("SELECT * FROM Author where authorId=?", book.AuthorID)

	// To store author
	var author models.Author

	if err := result.Scan(&author.AuthID, &author.FirstName, &author.LastName, &author.Dob, &author.PenName); err != nil {
		return models.Book{}, err
	}

	book.Auth = author

	return book, nil
}

// Update method is to change data of Particular book
func (d Datastore) Update(iD string, book *models.Book) (models.Book, error) {
	// converting string to integer to check for invalid id
	id, err := strconv.Atoi(iD)
	if err != nil {
		return models.Book{}, nil
	}

	var scanbook models.Book

	row := d.db.QueryRow("select * from Book where bookId=?", id)
	if err2 := row.Scan(&scanbook.BookID, &scanbook.Title, &scanbook.AuthorID, &scanbook.Publication, &scanbook.PublishedDate); err2 != nil {
		return models.Book{}, err2
	}

	// Updating book data
	_, err = d.db.Exec("UPDATE Book SET title=?, Publication=? , PublishedDate=? WHERE bookId=?",
		book.Title, book.Publication, book.PublishedDate, id)
	if err != nil {
		return models.Book{}, err
	}

	return *book, nil
}

// Delete method is remove Book by its ID
func (d Datastore) Delete(iD string) (int, error) {
	// converting string to integer to check for invalid id
	id, err := strconv.Atoi(iD)
	if err != nil {
		return 0, err
	}

	// Checking book exist or not
	var book models.Book

	row := d.db.QueryRow("select * from Book where bookId=?", id)

	if err2 := row.Scan(&book.BookID, &book.Title, &book.AuthorID, &book.Publication, &book.PublishedDate); err2 != nil {
		return 0, err2
	}

	if err != nil {
		return 0, err
	}

	// Now deleting data from table
	res, err := d.db.Exec("DELETE FROM Book where bookId=?", id)
	if err != nil {
		return 0, err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowAffected), nil
}
