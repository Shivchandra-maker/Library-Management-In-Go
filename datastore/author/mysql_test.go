package author

import (
	"database/sql/driver"
	"errors"
	"log"
	"reflect"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"Three-Layer-Architecture/models"
)

// Testing Post Author
func TestAuthor_Post(t *testing.T) {
	testcases := []struct {
		desc         string
		req          models.Author
		resp         models.Author
		lastInsertID int64
		rowAffected  int64
		err          error
	}{
		{desc: "valid details", req: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, resp: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, lastInsertID: 1, rowAffected: 1},
		{desc: "duplicate id", req: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, err: errors.New(" Duplicate entry '1' for key 'PRIMARY'")},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// mocking insert exec query
		mock.ExpectExec("insert into Author(authorId,firstName,lastName,dob,penName) values (?,?,?,?,?)").
			WithArgs(v.req.AuthID, v.req.FirstName, v.req.LastName, v.req.Dob, v.req.PenName).
			WillReturnResult(sqlmock.NewResult(v.lastInsertID, v.rowAffected)).WillReturnError(v.err)

		d := New(db)

		resp, err := d.Post(v.req)

		// Comparing body
		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		// Comparing errors
		if err != v.err {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// Testing Put Author
func TestAuthor_Put(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		resp models.Author
		rows *sqlmock.Rows
		res  driver.Result
		err  error
	}{
		{desc: "valid", id: "1", rows: sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"}).
			AddRow(1, "Rajan", "Sharma", "26/04/2001", "Rajan"), resp: models.Author{AuthID: 1, FirstName: "Rajan",
			LastName: "Sharma", Dob: "26/04/2001", PenName: "Rajan"}, res: sqlmock.NewResult(1, 1)},
		{desc: "id not exist", id: "11", rows: sqlmock.NewRows([]string{"authorId", "firstName",
			"lastName", "dob", "penName"})},
		{desc: "error id conversion", id: "a", rows: sqlmock.NewRows([]string{"authorId", "firstName",
			"lastName", "dob", "penName"})},
		{desc: "error in exec", id: "5", rows: sqlmock.NewRows([]string{"authorId", "firstName",
			"lastName", "dob", "penName"}).AddRow(1, "Sonu", "Sharma", "26/04/2000", "Sonu"),
			err: errors.New("err"),
		},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// conveting id string to integer
		id, err := strconv.Atoi(v.id)
		if err != nil {
			log.Printf("%v", err)
		}

		mock.ExpectQuery("select * from Author where authorId=?").WithArgs(id).
			WillReturnRows(v.rows)

		// Mocking Exec for updating data
		mock.ExpectExec("UPDATE Author SET firstName=?, lastName=? , dob=? , penName=? WHERE authorId=?").
			WithArgs(v.resp.FirstName, v.resp.LastName, v.resp.Dob, v.resp.PenName, id).WillReturnResult(v.res).WillReturnError(v.err)

		d := New(db)

		resp, err := d.Update(v.id, v.resp)

		// Comparing body
		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		// Getting error
		if err != nil {
			log.Printf("desc : %v ,[TEST%d]Failed. Got %v", v.desc, i+1, err)
		}
	}
}

// Testing Delete Author
func TestAuthor_Delete(t *testing.T) {
	testcases := []struct {
		desc           string
		ID             string
		rows           *sqlmock.Rows
		rowAffected    int64
		lastInsertedID int64
	}{
		{desc: "valid", ID: "1", rowAffected: 1, rows: sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"}).
			AddRow(1, "Rajan", "Sharma", "12/03/2012", "Sharma")},
		{desc: "id not exist", ID: "11",
			rows: sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"})},
		{desc: "id to string err", ID: "a",
			rows: sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"})},
		{desc: "missig fields", ID: "1", rowAffected: 1, rows: sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"}).
			AddRow(1, "", "Sharma", "12/03/2012", "Sharma")},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// conveting id string to integer
		id, err := strconv.Atoi(v.ID)
		if err != nil {
			log.Printf("%v", err)
		}

		// Mocking Query for checking authorId
		mock.ExpectQuery("select * from Author where authorId=?").WithArgs(id).
			WillReturnRows(v.rows)

		// Mocking Exec for deleting book
		mock.ExpectExec("delete from Book where bookId=?").WithArgs(id).
			WillReturnResult(sqlmock.NewResult(v.lastInsertedID, v.rowAffected))

		// Mocking Exec for deleting author
		mock.ExpectExec("delete from Author where authorId=?").WithArgs(id).
			WillReturnResult(sqlmock.NewResult(v.lastInsertedID, v.rowAffected))

		d := New(db)

		resp, err := d.Delete(v.ID)

		if reflect.DeepEqual(resp, v.rowAffected) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.rowAffected)
		}

		if err != nil {
			log.Printf("desc : %v ,[TEST%d]Failed. Got %v", v.desc, i+1, err)
		}
	}
}
