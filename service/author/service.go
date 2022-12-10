package author

import (
	"Three-Layer-Architecture/datastore"
	"Three-Layer-Architecture/models"
	"errors"
	"fmt"
	"strconv"
)

type Service struct {
	datastore datastore.Author
}

func New(author datastore.Author) Service {
	return Service{author}
}

// Post Author details
func (a Service) Post(auth models.Author) (models.Author, error) {
	// Checking for invalid id
	if auth.AuthID <= 0 {
		return models.Author{}, errors.New("invalid id")
	}

	if isMissingFields(auth) {
		return models.Author{}, fmt.Errorf("missing fields")
	}

	author, err := a.datastore.Post(auth)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

// Update Author details
func (a Service) Update(id string, auth models.Author) (models.Author, error) {
	if id == "" {
		return models.Author{}, errors.New("missing id")
	}

	// converting string to integer
	iD, err := strconv.Atoi(id)
	if err != nil {
		return models.Author{}, err
	}

	// Checking invalid id
	if iD <= 0 {
		return models.Author{}, errors.New("invalid id")
	}

	if isMissingFields(auth) {
		return models.Author{}, errors.New("missing fields")
	}

	author, err := a.datastore.Update(id, auth)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

// Delete Author by its ID
func (a Service) Delete(id string) (int, error) {
	// Checking for missing id
	if id == "" {
		return 0, errors.New("missing id")
	}

	// converting string to integer
	iD, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	// Checking for invalid id
	if iD <= 0 {
		return 0, fmt.Errorf("invalid id")
	}

	rowaffected, err := a.datastore.Delete(id)
	if err != nil {
		return 0, err
	}

	return rowaffected, nil
}

func isMissingFields(auth models.Author) bool {
	if auth.FirstName == "" || auth.LastName == "" || auth.PenName == "" || auth.Dob == "" {
		return true
	}

	return false
}
