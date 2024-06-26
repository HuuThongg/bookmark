package api

import (
	"database/sql"
	"errors"
	"fmt"
	"go-bookmark/auth"
	"go-bookmark/db/sqlc"
	"go-bookmark/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgconn"
)

func (h *BaseHandler) SearchLinks(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "query")

	q := sqlc.New(h.db)

	payload := r.Context().Value("payload").(*auth.PayLoad)

	percent := "%"

	linkTitle := fmt.Sprintf("%s%s%s", percent, query, percent)

	arg := sqlc.SearchLinkzParams{
		LinkTitle: linkTitle,
		AccountID: payload.AccountID,
	}

	links, err := q.SearchLinkz(r.Context(), arg)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			log.Printf("failed to search links: %v", err)
			util.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			err := errors.New("no links found matching the search query")
			log.Println(err)
			util.Response(w, err.Error(), http.StatusNotFound)
			return
		} else {
			log.Printf("failed to search links: %v", err)
			util.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	util.JsonResponse(w, links)
}

func (h *BaseHandler) SearchFolders(w http.ResponseWriter, r *http.Request) {

	q := sqlc.New(h.db)

	query := chi.URLParam(r, "query")

	payload := r.Context().Value("payload").(*auth.PayLoad)

	percent := "%"

	folderName := fmt.Sprintf("%s%s%s", percent, query, percent)

	arg := sqlc.SearchFolderzParams{
		//PlaintoTsquery: query,
		FolderName: folderName,
		AccountID:  payload.AccountID,
	}

	folders, err := q.SearchFolderz(r.Context(), arg)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			log.Printf("failed to search folders: %v", err)
			util.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			err := errors.New("no folders found matching the search query")
			log.Println(err)
			util.Response(w, err.Error(), http.StatusNotFound)
			return
		} else {
			log.Printf("failed to search folders: %v", err)
			util.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	util.JsonResponse(w, folders)
}
