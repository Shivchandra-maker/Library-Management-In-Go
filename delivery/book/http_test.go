package book

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"Three-Layer-Architecture/models"
	"Three-Layer-Architecture/service"
)

// TestPostBook function is to test Post book method for posting books
func TestPostBook(t *testing.T) {
	testcase := []struct {
		desc               string
		req                models.Book
		resp               models.Book
		expectedStatusCode int
		err                error
	}{
		{desc: "valid details ", req: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			resp: models.Book{BookID: 1, AuthorID: 1, Auth: models.Author{AuthID: 1, FirstName: "Chetan",
				LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"}},
			expectedStatusCode: http.StatusOK},
		{desc: "error from svc", req: models.Book{BookID: -11, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "2 States", Publication: "Scholastic", PublishedDate: "16/03/2016"},
			expectedStatusCode: http.StatusBadRequest, err: errors.New("dsadc")},
	}

	ctr := gomock.NewController(t)
	mockBook := service.NewMockBook(ctr)
	delivery := New(mockBook)

	for i, v := range testcase {
		body, err := json.Marshal(v.req)
		if err != nil {
			log.Printf("Not able to marshal : %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/book", bytes.NewReader(body))
		w := httptest.NewRecorder()

		if v.req.BookID == 1 {
			mockBook.EXPECT().Post(&v.req).Return(v.resp, v.err)
		} else {
			mockBook.EXPECT().Post(&v.req).Return(v.resp, v.err).AnyTimes()
		}

		delivery.Post(w, req)

		res := w.Result()

		var book models.Book

		book = HelperReader(&book, res)

		if !reflect.DeepEqual(book, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, book, v.resp)
		}

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res.StatusCode, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

// TestGetAllBooks function is to test GetAll method for fetching details of books
func TestGetAllBooks(t *testing.T) {
	testcases := []struct {
		desc               string
		output             any
		expectedStatusCode int
		err                error
	}{
		{desc: "valid details", output: []models.Book{
			{BookID: 1, AuthorID: 1, Auth: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
				Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"}}, expectedStatusCode: http.StatusOK},
	}

	ctr := gomock.NewController(t)
	mockBook := service.NewMockBook(ctr)
	delivery := New(mockBook)

	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)

		w := httptest.NewRecorder()

		mockBook.EXPECT().GetAll().Return(v.output, v.err).AnyTimes()

		delivery.GetAll(w, req)

		res := w.Result()

		var output []models.Book

		output = GetAllHelper(&output, res)

		if !reflect.DeepEqual(output, v.output) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, output, v.output)
		}

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res.StatusCode, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

// TestGetBook function is to test Getbyid method for fetching a book
func TestGetBook(t *testing.T) {
	testcases := []struct {
		desc               string
		reqid              string
		resp               models.Book
		expectedStatusCode int
		err                error
	}{
		{desc: "valid details", reqid: "1", resp: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, expectedStatusCode: http.StatusOK},
		{desc: "error from svc", reqid: "", resp: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			Title: "States", Publication: "Scholastic", PublishedDate: "16/03/2016"}, expectedStatusCode: http.StatusBadRequest,
			err: errors.New("missing id")},
	}

	ctr := gomock.NewController(t)
	mockBook := service.NewMockBook(ctr)
	delivery := New(mockBook)

	for i, v := range testcases {
		params := url.Values{}
		params.Add("bookId", v.reqid)

		req := httptest.NewRequest(http.MethodGet, "/books?"+params.Encode(), nil)

		w := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{"id": v.reqid})

		mockBook.EXPECT().Getbyid(v.reqid).Return(v.resp, v.err).AnyTimes()

		delivery.Getbyid(w, req)

		res := w.Result()

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res.StatusCode, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

// TestUpdateBook function is to test Put method for updating details of book
func TestUpdateBook(t *testing.T) {
	testcases := []struct {
		desc               string
		reqid              string
		reqbody            models.Book
		resp               models.Book
		expectedStatusCode int
		err                error
	}{
		{desc: "valid", reqid: "1", reqbody: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"},
			resp: models.Book{BookID: 1, AuthorID: 1, Auth: models.Author{AuthID: 1, FirstName: "Gaurav",
				LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
				Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}, expectedStatusCode: http.StatusOK},
		{desc: "error from svc", reqid: "", reqbody: models.Book{BookID: 1, AuthorID: 1,
			Auth:  models.Author{AuthID: 1, FirstName: "Gaurav", LastName: "Singh", Dob: "07/04/2001", PenName: "Gaurav"},
			Title: "300 Days", Publication: "Penguin", PublishedDate: "17/03/2016"}, err: errors.New("missing id"),
			expectedStatusCode: http.StatusBadRequest},
	}

	ctr := gomock.NewController(t)
	mockBook := service.NewMockBook(ctr)
	delivery := New(mockBook)

	for i, v := range testcases {
		params := url.Values{}
		params.Add("bookId", v.reqid)

		body, err := json.Marshal(v.reqbody)
		if err != nil {
			log.Printf("Not able to marshal : %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/books?"+params.Encode(), bytes.NewReader(body))

		w := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{"id": v.reqid})

		mockBook.EXPECT().Update(v.reqid, &v.reqbody).Return(v.resp, v.err).AnyTimes()

		delivery.Update(w, req)

		res := w.Result()

		var book models.Book

		book = HelperReader(&book, res)

		if !reflect.DeepEqual(book, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, book, v.resp)
		}

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res.StatusCode, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

// TestDeleteBook function is to test delete method to remove any book
func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc               string
		reqid              string
		rowAffected        int
		expectedStatusCode int
		err                error
	}{
		{desc: "valid", reqid: "1", rowAffected: 1, expectedStatusCode: http.StatusOK},
		{desc: "missing id", reqid: "", rowAffected: 0, err: errors.New("missing id"), expectedStatusCode: http.StatusBadRequest},
	}

	ctr := gomock.NewController(t)
	mockBook := service.NewMockBook(ctr)
	delivery := New(mockBook)

	for i, v := range testcases {
		params := url.Values{}
		params.Add("bookId", v.reqid)

		req := httptest.NewRequest(http.MethodGet, "/books?"+params.Encode(), nil)

		w := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{"id": v.reqid})

		mockBook.EXPECT().Delete(v.reqid).Return(v.rowAffected, v.err).AnyTimes()

		delivery.Delete(w, req)

		res := w.Result()

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

func HelperReader(book *models.Book, res *http.Response) models.Book {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v", err)
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("%v", err)
	}

	return *book
}

func GetAllHelper(output *[]models.Book, res *http.Response) []models.Book {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v", err)
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Printf("%v", err)
	}

	return *output
}
