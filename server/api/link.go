package api

import (
	"encoding/json"
	"net/http"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type URL struct {
	URL      string `json:"url"`
	FolderId string `json:"folder_id"`
}

func (u URL) validate(requestValidateChan chan error) error {
	validationError := validation.ValidateStruct(&u,
		validation.Field(&u.URL, validation.Required.Error("url is required"), is.URL.Error("url must be a valid url")),
	)

	requestValidateChan <- validationError
	return validationError
}

func (h *BaseHandler) AddLink(w http.ResponseWriter, r *http.Request) {
	rBody := json.NewDecoder(r.Body)
	rBody.DisallowUnknownFields()

	var req URL

	if err := rBody.Decode(&req); err != nil {
		ErrorDecodingRequest(w, err)
		return
	}
	requestValidationChan := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		req.validate(requestValidationChan)
	}()
	validationError := <-requestValidationChan

	if validationError != nil {
		ErrorInvalidRequest(w, validationError)
		return
	}

	var host string
	var urlToOpen string

}
