package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-bookmark/auth"
	"go-bookmark/db"
	"go-bookmark/util"
	"log"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgconn"
)

type createFolderRequest struct {
	FolderName string `json:"folder_name"`
	FolderID   string `json:"parent_folder_id"`
}

func (s createFolderRequest) validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.FolderName, validation.Required.When(s.FolderName == "").Error("folder name is required"), validation.Match(regexp.MustCompile("^[^?[\\]{}|\\\\`./!@#$%^&*()_-]+$")).Error("folder name must not have special characters"), validation.Length(1, 100).Error("folder name must be at least 1 character long")),
	)
}

type CreateFolderRequestBody struct {
	PayLoad *auth.PayLoad
	Body    *createFolderRequest
}

func newCreateFolderRequestBody(p auth.PayLoad, b createFolderRequest) *CreateFolderRequestBody {
	return &CreateFolderRequestBody{
		PayLoad: &p,
		Body:    &b,
	}
}

func AuthorizeCreateFolderRequest() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			body := json.NewDecoder(r.Body)
			body.DisallowUnknownFields()

			var req createFolderRequest
			if err := body.Decode(&req); err != nil {
				if e, ok := err.(*json.SyntaxError); ok {
					log.Printf("syntax error at byte offset %d", e.Offset)
					util.Response(w, "something went wrong", http.StatusInternalServerError)
					return
				} else {
					log.Printf("error decoding request body to struct: %v", err)
					util.Response(w, "bad requset", http.StatusBadRequest)
					return
				}
			}
			// validate request
			if err := req.validate(); err != nil {
				log.Println(err)
				util.Response(w, err.Error(), http.StatusBadRequest)
				return
			}

			// check if use wants to create a root folder
			if req.FolderID == "null" {
				log.Println("AuthorizeCreateFolderRequest(), folderID is null")
				payload := r.Context().Value("payload").(*auth.PayLoad)
				rB := newCreateFolderRequestBody(*payload, req)
				log.Printf("rb : %v", rB)
				ctx := context.WithValue(r.Context(), "createFolderRequest", rB)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			log.Println("folder ID is not null")

			// get user payload from request
			payload := r.Context().Value("payload").(*auth.PayLoad)

			// get parent folder
			folder, err := db.ReturnFolder(r.Context(), req.FolderID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					log.Println("collection not found")
					util.Response(w, "folder not found", http.StatusUnauthorized)
					return
				}
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					log.Println(pgErr)
					util.Response(w, "something went wrong", http.StatusInternalServerError)
					return
				}

			}

			// check if user owns parent folder
			if folder.AccountID == payload.AccountID {
				rB := newCreateFolderRequestBody(*payload, req)
				ctx := context.WithValue(r.Context(), "createFolderRequest", rB)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}
