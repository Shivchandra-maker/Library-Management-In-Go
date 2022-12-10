package book

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"Three-Layer-Architecture/models"
)

// Test_Post Book
func Test_Post(t *testing.T) {
	testcases := []struct {
		desc         string
		req          models.Book
		response     models.Book
		lastInsertID int64
		rowAffected  int64
		err          error
	}{
		{desc: "valid details", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, response: models.Book{BookID: 1,
			AuthorID: 1, Auth: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, lastInsertID: 1, rowAffected: 1},
		{desc: "duplicate id", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, err: errors.New(" Duplicate entry '1' for key 'PRIMARY'")},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// Mocking insert query for book
		mock.ExpectExec("insert into Book(bookId,title,authorId,Publication,PublishedDate) values (?,?,?,?,?)").
			WithArgs(v.req.BookID, v.req.Title, v.req.AuthorID, v.req.Publication, v.req.PublishedDate).
			WillReturnResult(sqlmock.NewResult(v.lastInsertID, v.rowAffected)).WillReturnError(v.err)

		// injecting mock db
		d := New(db)

		resp, err := d.Post(&v.req)

		// comparing body
		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.response)
		}

		// comparing error
		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// Test_GetAll all book
func Test_GetAll(t *testing.T) {
	testcases := []struct {
		desc    string
		resp    []models.Book
		columns []string
		err     error
	}{
		{desc: "valid details ", resp: []models.Book{
			{BookID: 1, AuthorID: 1,
				Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
				Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			{BookID: 2, AuthorID: 1,
				Auth:  models.Author{AuthID: 1, FirstName: "Vikram", LastName: "Seth", Dob: "26/04/2001", PenName: "Vikram"},
				Title: "3 States", Publication: "Penguin", PublishedDate: "11/03/2016"}},
			columns: []string{"bookId", "auth", "title", "authorId", "Publication", "PublishedDate"}},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// Mocking select all books query
		mock.ExpectQuery("SELECT * FROM Book").WillReturnRows(sqlmock.NewRows(v.columns)).WillReturnError(v.err)

		// Mocking select author for book query
		mock.ExpectQuery("SELECT * FROM Author where authorId=?").WithArgs(v.resp[0].Auth.AuthID).
			WillReturnRows(sqlmock.NewRows(v.columns).FromCSVString("1, Chetan, Bhagat, 06/04/2001, Chetan")).WillReturnError(v.err)

		// injecting mock db
		datastore := New(db)

		resp, err := datastore.GetAll()

		// Comparing body
		if reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		// Comparing errors
		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// Test_GetbyidInvalid Testing book Get by id
func Test_GetbyidInvalid(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		resp models.Book
		err  error
	}{
		{desc: "id not exist", id: "11", resp: models.Book{BookID: 1, AuthorID: 1,
			Auth: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
				PenName: "Chetan"}, Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			err: errors.New("sql: no rows in result set")},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// converting id string to integer
		id, err := strconv.Atoi(v.id)
		if err != nil {
			log.Printf("%v", err)
		}

		// Mocking Query for reading book
		mock.ExpectQuery("select * from Book where bookId=?").WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{""})).WillReturnError(v.err)

		// Mocking Query for reading author of that book
		mock.ExpectQuery("SELECT * FROM Author where authorId=?").WithArgs(v.resp.AuthorID).
			WillReturnRows(sqlmock.NewRows([]string{""})).WillReturnError(v.err)

		// Injecting mock DB
		d := New(db)

		resp, err := d.Getbyid(v.id)

		// Comparing body
		if reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		// Comparing errors
		if err != v.err {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// Test_GetbyidValid Testing book Get by id
func Test_GetbyidValid(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		resp models.Book
		err  error
	}{
		{desc: "valid", id: "1", resp: models.Book{BookID: 1, AuthorID: 1,
			Auth: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
				PenName: "Chetan"}, Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"}},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// converting id string to integer
		id, err := strconv.Atoi(v.id)
		if err != nil {
			log.Printf("%v", err)
		}

		// Mocking Query for reading book
		mock.ExpectQuery("select * from Book where bookId=?").WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"bookId", "title", "authorId", "Publication", "PublishedDate"}).
				FromCSVString("1,States,1,Scholastic,16/03/2016")).WillReturnError(v.err)

		// Mocking Query for reading author of that book
		mock.ExpectQuery("SELECT * FROM Author where authorId=?").WithArgs(v.resp.AuthorID).
			WillReturnRows(sqlmock.NewRows([]string{"authorId", "firstName", "lastName", "dob", "penName"}).
				FromCSVString("1,Chetan,Bhagat,06/04/2001,Chetan")).WillReturnError(v.err)

		// Injecting mock DB
		d := New(db)

		resp, err := d.Getbyid(v.id)

		// Comparing body
		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		// Comparing errors
		if err != v.err {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// Test_Put book
func Test_Put(t *testing.T) {
	testcases := []struct {
		desc         string
		id           string
		req          models.Book
		resp         models.Book
		row          *sqlmock.Rows
		lastInsertID int64
		rowAffected  int64
		err          error
	}{
		{desc: "valid", id: "1", req: models.Book{BookID: 1, AuthorID: 1,
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}, lastInsertID: 1, rowAffected: 1,
			resp: models.Book{BookID: 1, AuthorID: 1, Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}, row: sqlmock.
				NewRows([]string{"bookId", "title", "authorId", "Publication", "PublishedDate"}).
				AddRow(1, "300 Days", 1, "Penguin", "17/03/2016")},
		{desc: "id not exist", id: "11", req: models.Book{BookID: 1, AuthorID: 1,
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}, err: errors.New("sql: no rows in result set"), row: sqlmock.
			NewRows([]string{"bookId", "title", "authorId", "Publication", "PublishedDate"})},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		id, err := strconv.Atoi(v.id)
		if err != nil {
			log.Printf("%v", err)
		}

		mock.ExpectQuery("select * from Book where bookId=?").WithArgs(id).
			WillReturnRows(v.row).WillReturnError(v.err)

		// Mocking Exec query for updating data
		mock.ExpectExec("UPDATE Book SET title=?, Publication=? , PublishedDate=? WHERE bookId=?").
			WithArgs(v.resp.Title, v.resp.Publication, v.resp.PublishedDate, id).
			WillReturnResult(sqlmock.NewResult(v.lastInsertID, v.rowAffected)).WillReturnError(v.err)

		// Injecting mock Db
		d := New(db)

		resp, err := d.Update(v.id, &v.req)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// Test_Delete book
func Test_Delete(t *testing.T) {
	testcases := []struct {
		desc           string
		id             string
		row            *sqlmock.Rows
		rowAffected    int64
		lastInsertedID int64
		err            error
	}{
		{desc: "valid", id: "1", rowAffected: 1, lastInsertedID: 1, row: sqlmock.
			NewRows([]string{"bookId", "title", "authorId", "Publication", "PublishedDate"}).AddRow(1, "Journey", 1, "Penguin", "12/04/2001")},
		{desc: "id not exist", id: "11", err: errors.New("sql: no rows in result set"), row: sqlmock.
			NewRows([]string{"bookId", "title", "authorId", "Publication", "PublishedDate"})},
	}

	// Customize SQL query matching
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Closing DB after all things done
	defer db.Close()

	for i, v := range testcases {
		// converting string to integer to check for invalid id
		id, err := strconv.Atoi(v.id)
		if err != nil {
			log.Printf("%v", err)
		}

		// Mocking for checking book Id
		mock.ExpectQuery("select * from Book where bookId=?").WithArgs(id).WillReturnRows(v.row).WillReturnError(v.err)

		// Mocking delete query from book
		mock.ExpectExec("DELETE FROM Book where bookId=?").WithArgs(id).
			WillReturnResult(sqlmock.NewResult(v.lastInsertedID, v.rowAffected)).WillReturnError(v.err)

		// Injecting mock DB
		d := New(db)

		resp, err := d.Delete(v.id)

		// Comparing body
		if reflect.DeepEqual(resp, v.rowAffected) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.rowAffected)
		}

		// Comparing errors
		if err != v.err {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}
