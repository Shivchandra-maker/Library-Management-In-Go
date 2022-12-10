package author

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"Three-Layer-Architecture/models"
	"Three-Layer-Architecture/service"
)

type Delivery struct {
	service service.Author
}

func New(author service.Author) Delivery {
	return Delivery{author}
}

// Post Request method is to post request
func (a Delivery) Post(w http.ResponseWriter, r *http.Request) {
	// reading body
	auth, err := ReadReqbody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	author, err := a.service.Post(auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	result, err := json.Marshal(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)
	}

	w.WriteHeader(http.StatusCreated)

	fmt.Println("Successfully Post data")
}

// Update Request method is to update request
func (a Delivery) Update(w http.ResponseWriter, r *http.Request) {
	// storing id in map
	vars := mux.Vars(r)

	// reading body of request
	author, err := ReadReqbody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	auth, err := a.service.Update(vars["id"], author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	result, err := json.Marshal(auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Println("Successfully Update data")
}

// Delete method is to delete data from request
func (a Delivery) Delete(w http.ResponseWriter, r *http.Request) {
	// storing id in map
	vars := mux.Vars(r)

	_, err := a.service.Delete(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(err, w)

		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Println("Successfully Deleted..!!")
}

func writeError(err error, w http.ResponseWriter) {
	_, errs := w.Write([]byte(err.Error()))
	if errs != nil {
		log.Printf("%v", errs)
		w.WriteHeader(http.StatusBadRequest)

		return
	}
}

func ReadReqbody(r *http.Request) (models.Author, error) {
	// Reading body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return models.Author{}, err
	}

	var author models.Author

	// Decoding
	err = json.Unmarshal(body, &author)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}
