package author

import (
	"Three-Layer-Architecture/models"
	"database/sql"
	"errors"
	"strconv"
)

type Datastore struct {
	db *sql.DB
}

func New(db *sql.DB) Datastore {
	return Datastore{db}
}

// Post method is to post the data in Author table
func (d Datastore) Post(auth models.Author) (models.Author, error) {
	// inserting data into db
	_, err := d.db.Exec("insert into Author(authorId,firstName,lastName,dob,penName) values (?,?,?,?,?)",
		auth.AuthID, auth.FirstName, auth.LastName, auth.Dob, auth.PenName)
	if err != nil {
		return models.Author{}, err
	}

	return auth, nil
}

// Update method is to update the data in Author table
func (d Datastore) Update(iD string, auth models.Author) (models.Author, error) {
	// conveting id string to integer
	id, err := strconv.Atoi(iD)
	if err != nil {
		return models.Author{}, errors.New("strconv.Atoi: parsing a")
	}

	var author models.Author

	row := d.db.QueryRow("select * from Author where authorId=?", id)

	if err2 := row.Scan(&author.AuthID, &author.FirstName, &author.LastName, &author.Dob, &author.PenName); err2 != nil {
		return models.Author{}, err2
	}

	// now updating the table
	_, err = d.db.Exec("UPDATE Author SET firstName=?, lastName=? , dob=? , penName=? WHERE authorId=?",
		auth.FirstName, auth.LastName, auth.Dob, auth.PenName, id)
	if err != nil {
		return models.Author{}, err
	}

	return auth, nil
}

// Delete method is to delete the data in Author
func (d Datastore) Delete(iD string) (int, error) {
	id, err := strconv.Atoi(iD)
	if err != nil {
		return 0, err
	}

	// Checking author is present or not
	var author models.Author

	result := d.db.QueryRow("select * from Author where authorId=?", id)

	if err2 := result.Scan(&author.AuthID, &author.FirstName, &author.LastName, &author.Dob, &author.PenName); err2 != nil {
		return 0, err2
	}

	// Firstly deleting data from book table because book can't exist without author ( Foreign key )
	res, err := d.db.Exec("delete from Book where bookId=?", id)
	if err != nil {
		return 0, err
	}

	// now row is affected
	rowaffected, err2 := res.RowsAffected()
	if err2 != nil {
		return 0, err2
	}

	// Now deleting data from Author
	_, err = d.db.Exec("delete from Author where authorId=?", id)
	if err != nil {
		return 0, err
	}

	return int(rowaffected), nil
}
