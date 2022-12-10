package author

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"Three-Layer-Architecture/datastore"
	"Three-Layer-Architecture/models"
)

// TestStorer_Post function is to test post author details for valid conditions
func TestAuthor_Post(t *testing.T) {
	testcases := []struct {
		desc     string
		req      models.Author
		response models.Author
		err      error
	}{
		{desc: "valid details", req: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			response: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"}, err: nil},
		{desc: "missing first name", req: models.Author{AuthID: 1, LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			err: fmt.Errorf("missing fields")},
		{desc: "invalid id", req: models.Author{AuthID: -11, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001", PenName: "Chetan"},
			err: fmt.Errorf("invalid id")},
		{desc: "missing last name", req: models.Author{AuthID: 1, FirstName: "Chetan", Dob: "06/04/2001", PenName: "Chetan"},
			err: fmt.Errorf("missing fields")},
		{desc: "missing dob", req: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", PenName: "Chetan"},
			err: fmt.Errorf("missing fields")},
		{desc: "missing penname", req: models.Author{AuthID: 1, FirstName: "Chetan", LastName: "Bhagat", Dob: "06/04/2001"},
			err: fmt.Errorf("missing fields")},
	}

	ctr := gomock.NewController(t)
	mockAuthor := datastore.NewMockAuthor(ctr)
	service := New(mockAuthor)

	for i, v := range testcases {
		mockAuthor.EXPECT().Post(v.req).Return(v.response, v.err).AnyTimes()

		resp, err := service.Post(v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.response)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// TestAuthor_Put function is to test update valid author details
func TestAuthor_Put(t *testing.T) {
	testcases := []struct {
		desc string
		id   string
		req  models.Author
		err  error
	}{
		{desc: "valid", id: "1", req: models.Author{AuthID: 1, FirstName: "Rajan", LastName: "Sharma",
			Dob: "26/04/2001", PenName: "Rajan"}},
		{desc: "invalid id", id: "-11", err: errors.New("invalid id")},
		{desc: "missing id", err: errors.New("missing id")},
	}

	ctr := gomock.NewController(t)
	mockAuthor := datastore.NewMockAuthor(ctr)
	service := New(mockAuthor)

	for i, v := range testcases {
		mockAuthor.EXPECT().Update(v.id, v.req).Return(v.req, v.err).AnyTimes()

		resp, err := service.Update(v.id, v.req)

		if !reflect.DeepEqual(resp, v.req) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.req)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}

// TestAuthor_Delete function is to test for remove author
func TestAuthor_Delete(t *testing.T) {
	testcases := []struct {
		desc        string
		id          string
		rowaffected int
		err         error
	}{
		{desc: "valid", id: "1", rowaffected: 1, err: nil},
		{desc: "invalid id", id: "-11", err: fmt.Errorf("invalid id")},
		{desc: "missing id", err: fmt.Errorf("missing id")},
	}

	ctr := gomock.NewController(t)
	mockAuthor := datastore.NewMockAuthor(ctr)
	service := New(mockAuthor)

	for i, v := range testcases {
		mockAuthor.EXPECT().Delete(v.id).Return(v.rowaffected, v.err).AnyTimes()

		resp, err := service.Delete(v.id)

		if !reflect.DeepEqual(resp, v.rowaffected) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, resp, v.rowaffected)
		}

		if !reflect.DeepEqual(err, v.err) {
			t.Errorf("desc : %v ,[TEST%d]Failed. Got %v\tExpected %v\n", v.desc, i+1, err, v.err)
		}
	}
}
