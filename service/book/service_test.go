package book

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"Three-Layer-Architecture/datastore"
	"Three-Layer-Architecture/models"
)

// TestBook_Post function is to test post author details
func TestBook_Post(t *testing.T) {
	testcases := []struct {
		desc     string
		req      models.Book
		response models.Book
		err      error
	}{
		{desc: "valid details", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			response: models.Book{BookID: 1, AuthorID: 1,
				Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
				Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}},
		{desc: "valid details", req: models.Book{BookID: 2, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			response: models.Book{BookID: 2, AuthorID: 1,
				Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
				Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}},
		{desc: "valid details", req: models.Book{BookID: 3, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			response: models.Book{BookID: 3, AuthorID: 1,
				Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
				Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}},
		{desc: "invalid id", req: models.Book{BookID: -11, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, err: errors.New("invalid id")},
		{desc: "invalid publication", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Lenin", PublishedDate: "16/03/2016"}, err: errors.New("invalid publication")},
		{desc: "invalid publishedDate", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2061"}, err: errors.New("invalid publishedDate")},
		{desc: "missing title", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:        models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Publication: "Scholastic", PublishedDate: "16/03/2016"}, err: errors.New("missing book fields")},
		{desc: "missing publication", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", PublishedDate: "16/03/2016"}, err: errors.New("missing book fields")},
		{desc: "missing publishedDate", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic"}, err: errors.New("missing book fields")},
		{desc: "missing Author fields", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, err: errors.New("missing author fields")},
	}

	for i, v := range testcases {
		ctr := gomock.NewController(t)
		mockBook := datastore.NewMockBook(ctr)
		service := New(mockBook)

		mockBook.EXPECT().Post(&v.req).Return(v.response, v.err).AnyTimes()

		resp, err := service.Post(&v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.response)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// TestBook_GetAll function is to test for getting all books
func TestBook_GetAll(t *testing.T) {
	testcases := []struct {
		desc string
		resp []models.Book
		err  error
	}{

		{desc: "valid details ", resp: []models.Book{{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"}}},
	}

	for i, v := range testcases {
		ctr := gomock.NewController(t)
		mockBook := datastore.NewMockBook(ctr)
		service := New(mockBook)

		mockBook.EXPECT().GetAll().Return(v.resp, v.err).AnyTimes()

		resp, err := service.GetAll()

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// TestBook_Getbyid function is to test for get a book
func TestBook_Getbyid(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		resp models.Book
		err  error
	}{

		{desc: "valid detail", id: "1", resp: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"}},
		{desc: "invalid id", id: "-11", err: errors.New("invalid id")},
		{desc: "missing id", err: errors.New("missing id")},
	}

	ctr := gomock.NewController(t)
	mockBook := datastore.NewMockBook(ctr)
	service := New(mockBook)

	for i, v := range testcases {
		mockBook.EXPECT().Getbyid(v.id).Return(v.resp, v.err).AnyTimes()

		resp, err := service.Getbyid(v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("Desc : %v,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// TestBook_Put function is to test for updating Book
func TestBook_Put(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		req  models.Book
		resp models.Book
		err  error
	}{
		{desc: "valid ", id: "1", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}, resp: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}},
		{desc: "missing id", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "lenin", PublishedDate: "17/03/2016"}, err: errors.New("missing id"), resp: models.Book{}},
		{desc: "invalid id", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "Arihant", PublishedDate: "17/03/2016"}, id: "-11", err: errors.New("invalid id")},
		{desc: "invalid publication", id: "1", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "lenin", PublishedDate: "17/03/2016"}, err: errors.New("invalid publication")},
		{desc: "invalid publishedDate", id: "1", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2061"}, err: errors.New("invalid publishedDate")},
		{desc: "missing book fields", id: "1", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:        models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Publication: "lenin", PublishedDate: "17/03/2016"}, err: errors.New("missing book fields")},
		{desc: "missing author fields", id: "1", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:        models.Author{AuthID: 1, LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Publication: "Arihant", PublishedDate: "17/03/2016"}, err: errors.New("missing author fields")},
	}

	for i, v := range testcases {
		ctr := gomock.NewController(t)
		mockBook := datastore.NewMockBook(ctr)
		service := New(mockBook)

		mockBook.EXPECT().Update(v.id, &v.req).Return(v.resp, v.err).AnyTimes()

		resp, err := service.Update(v.id, &v.req)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.resp)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// TestBook_Delete functio is to test for deleting a valid book
func TestBook_Delete(t *testing.T) {
	testcases := []struct {
		desc        string
		id          string
		rowAffected int
		err         error
	}{
		{desc: "valid", id: "1", rowAffected: 1},
		{desc: "missing id", err: errors.New("missing id")},
		{desc: "invalid id", id: "-11", err: errors.New("invalid id")},
	}

	ctr := gomock.NewController(t)
	mockBook := datastore.NewMockBook(ctr)
	service := New(mockBook)

	for i, v := range testcases {
		mockBook.EXPECT().Delete(v.id).Return(v.rowAffected, v.err).AnyTimes()

		resp, err := service.Delete(v.id)

		if !reflect.DeepEqual(resp, v.rowAffected) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.rowAffected)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}
