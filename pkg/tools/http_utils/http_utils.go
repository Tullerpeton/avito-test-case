package http_utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/avito-test-case/internal/app/errors/handler_errors"
)

func SetJSONResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(handler_errors.HttpIncorrectRequestBody.Error())); err != nil {
			log.Fatal(err)
		}
		return
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(result); err != nil {
		log.Fatal(err)
	}
}

func ParseBody(r *http.Request, savedStruct interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, savedStruct)
	if err != nil {
		return err
	}

	return nil
}
