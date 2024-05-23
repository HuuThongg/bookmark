package api

import (
	"encoding/json"
	"errors"
	"go-bookmark/util"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgconn"
)

var (
	internalServerError = "something went wrong"
	badRequest          = "bad request"
)

func ErrorPgError(w http.ResponseWriter, err error) {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		log.Println(pgErr)
		util.Response(w, internalServerError, http.StatusInternalServerError)
	} else {
		log.Print(err)
		util.Response(w, internalServerError, http.StatusInternalServerError)
	}
}
func ErrorInternalServerError(w http.ResponseWriter, err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Println(pgErr)
		util.Response(w, internalServerError, http.StatusInternalServerError)
	} else {
		log.Println(err)
		util.Response(w, internalServerError, http.StatusInternalServerError)
	}
}
func ErrorInternalServer(w http.ResponseWriter, err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Println(pgErr)
		util.Response(w, internalServerError, http.StatusInternalServerError)
	} else {
		log.Println(err)
		util.Response(w, internalServerError, http.StatusInternalServerError)
	}
}
func ErrorDecodingRequest(w http.ResponseWriter, err error) {
	if e, ok := err.(*json.SyntaxError); ok {
		log.Printf("JSON syntax error occurred at offset byte: %d", e.Offset)
		util.Response(w, "something went wrong", http.StatusInternalServerError)
	} else {
		log.Printf("error decoding request body to struct: %v", err)
		util.Response(w, "bad request", http.StatusBadRequest)
	}
}

func ErrorInvalidRequest(w http.ResponseWriter, err error) {
	if e, ok := err.(validation.InternalError); ok {
		log.Println(e)
		util.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	} else {
		log.Println(err)
		util.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
