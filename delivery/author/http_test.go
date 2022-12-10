package author

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

// TestPostAuthor function is to test post method
func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc               string
		req                any
		resp               models.Author
		expectedStatusCode int
		err                error
	}{
		{desc: "valid", req: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, resp: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, expectedStatusCode: http.StatusOK},
		{desc: "unmashal error", req: []models.Author{}, expectedStatusCode: http.StatusBadRequest},
		{desc: "errors from svc", req: models.Author{AuthID: -21, FirstName: "Sagar", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, expectedStatusCode: http.StatusBadRequest, err: errors.New("invalid id")},
	}

	ctr := gomock.NewController(t)
	mockAuthor := service.NewMockAuthor(ctr)
	delivery := New(mockAuthor)

	for i, v := range testcases {
		body, err := json.Marshal(v.req)
		if err != nil {
			log.Printf("Not able to marshal : %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/author", bytes.NewReader(body))
		w := httptest.NewRecorder()

		mockAuthor.EXPECT().Post(v.req).Return(v.resp, v.err).AnyTimes()

		delivery.Post(w, req)

		var author models.Author

		res := w.Result()

		author = Helper(author, res)

		if !reflect.DeepEqual(author, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, author, v.resp)
		}

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

// TestUpdateAuthor function is to test put method
func TestUpdateAuthor(t *testing.T) {
	testcases := []struct {
		desc               string
		reqid              string
		reqbody            any
		resp               models.Author
		expectedStatusCode int
		err                error
	}{
		{desc: "valid case", reqid: "1", reqbody: models.Author{AuthID: 1, FirstName: "Rajan", LastName: "Sharma",
			Dob: "26/04/2001", PenName: "Sharma"}, expectedStatusCode: http.StatusOK},
		{desc: "unmashel error", reqid: "1", reqbody: "something", expectedStatusCode: http.StatusBadRequest},
		{desc: "errors from svc", reqid: "11", reqbody: models.Author{AuthID: -21, FirstName: "Sagar", LastName: "Bhagat", Dob: "06/04/2001",
			PenName: "Chetan"}, expectedStatusCode: http.StatusBadRequest, err: errors.New("invalid id")},
	}

	ctr := gomock.NewController(t)
	mockAuthor := service.NewMockAuthor(ctr)
	delivery := New(mockAuthor)

	for i, v := range testcases {
		params := url.Values{}
		params.Add("bookId", v.reqid)

		body, err := json.Marshal(v.reqbody)
		if err != nil {
			log.Printf("Not able to marshal : %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/author?"+params.Encode(), bytes.NewReader(body))
		w := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{"id": v.reqid})

		mockAuthor.EXPECT().Update(v.reqid, v.reqbody).Return(v.resp, v.err).AnyTimes()

		// Mocking Update
		delivery.Update(w, req)

		var author models.Author

		res := w.Result()

		author = Helper(author, res)

		if !reflect.DeepEqual(author, v.resp) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, author, v.resp)
		}

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

// TestDeleteAuthor function is to test delete method
func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc               string
		reqid              string
		expectedStatusCode int
		rowAffected        int
		err                error
	}{
		{desc: "valid case", reqid: "1", expectedStatusCode: http.StatusNoContent, rowAffected: 1},
		{desc: "error from svc", reqid: "", expectedStatusCode: http.StatusBadRequest, err: errors.New("missing id")},
	}

	ctr := gomock.NewController(t)
	mockAuthor := service.NewMockAuthor(ctr)
	delivery := New(mockAuthor)

	for i, v := range testcases {
		params := url.Values{}
		params.Add("bookId", v.reqid)

		req := httptest.NewRequest(http.MethodPost, "/author/{id}"+v.reqid, nil)
		w := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{"id": v.reqid})

		mockAuthor.EXPECT().Delete(v.reqid).Return(v.rowAffected, v.err).AnyTimes()

		delivery.Delete(w, req)

		res := w.Result()

		if res.StatusCode != v.expectedStatusCode {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, res, v.expectedStatusCode)
		}

		res.Body.Close()
	}
}

func Helper(author models.Author, res *http.Response) models.Author {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v", err)
	}

	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Printf("%v", err)
	}

	return author
}
