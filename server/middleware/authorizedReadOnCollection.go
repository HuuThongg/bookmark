package middleware

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-bookmark/auth"
	"go-bookmark/db"
	"go-bookmark/util"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgconn"
)

type ReadRequsetOnCollectionDetails struct {
	FolderID  string
	AccountID int64
	Payload   auth.PayLoad
}

func newReadRequestOnCollectionDetails(folderID string, accountID int64, payload auth.PayLoad) *ReadRequsetOnCollectionDetails {
	return &ReadRequsetOnCollectionDetails{
		FolderID:  folderID,
		AccountID: accountID,
		Payload:   payload,
	}
}

func AuthorizeReadRequestOnCollection() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			payload := r.Context().Value("payload").(*auth.PayLoad)

			folderID := chi.URLParam(r, "folderID")

			accountID := chi.URLParam(r, "accountID")
			log.Printf("folderID: %v", folderID)
			log.Printf("accountID: %v", accountID)

			fmt.Printf("accountID: %v", accountID)
			log.Printf("accountid from payload: %v", payload.ID)
			account_id, err := strconv.Atoi(accountID)

			if err != nil {
				log.Printf("could not convert account id from url to int64 at authorizReadRequestOnCollectin.go: %v", err)
				util.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
			if int64(account_id) != payload.AccountID {
				log.Println("account IDs do not match")
				util.Response(w, errors.New("account IDs do not match").Error(), http.StatusUnauthorized)
				return
			}

			if folderID == "null" {
				// users onws folder hence is allowed to read from it hence no futher checks hence pss request details to context;
				log.Println("user onws folder")
				body := newReadRequestOnCollectionDetails(folderID, int64(account_id), *payload)
				log.Printf("body in folder;, %v", body)
				ctx := context.WithValue(r.Context(), "readRequestOnCollectionDetails", body)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			collection, err := db.ReturnFolder(r.Context(), folderID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					err := errors.New("collection not found")
					log.Println(err)
					util.Response(w, err.Error(), http.StatusUnauthorized)
					return
				}

				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					log.Printf("could not get collection at AuthorizeReadRequestOnCollection.go with pgErr: %v", pgErr)
					util.Response(w, "something went wrong", http.StatusInternalServerError)
					return
				}
				log.Printf("could not get collection at AuthorizeReadRequestOnCollection.go with pgErr: %v", pgErr)
				util.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}

			// check if collection belong to user
			if collection.AccountID == payload.AccountID {
				// collection belongs to user hence user is allowed to read from it no further check hence pass request details to context
				body := newReadRequestOnCollectionDetails(folderID, int64(account_id), *payload)
				ctx := context.WithValue(r.Context(), "readRequestOnCollectionDetails", body)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			// user does not own folder hence check if folder has been shared with them
			collectionMember, err := db.ReturnCollectionMemberByCollectionandMemberIDs(r.Context(), collection.FolderID, payload.AccountID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					//This specific folder has not been shared with them
					// check if any of it's ancestor has been shared with user
					ancestorsOfFolders, err := db.ReturnAncestorsOfFolder(r.Context(), collection.FolderID)
					if err != nil {
						if errors.Is(err, sql.ErrNoRows) {
							err := errors.New("collection not found")
							log.Println(err)
							util.Response(w, err.Error(), http.StatusUnauthorized)
							return
						}
						var pgErr *pgconn.PgError
						if errors.As(err, &pgErr) {
							log.Printf("authorizeReadRequestOnCollection: could not get ancestors of folder: pgErr: %v", pgErr)
							util.Response(w, "something went wrong", http.StatusInternalServerError)
							return
						}
						log.Printf("authorizeReadRequestOnCollection: could not get ancestors of folder: pgErr: %v", pgErr)
						util.Response(w, "something went wrong", http.StatusInternalServerError)
						return
					}

					for _, acancestorsOfFolder := range ancestorsOfFolders {
						val, err := db.CheckIfCollectionMemberExists(r.Context(), acancestorsOfFolder.FolderID, payload.AccountID)
						if err != nil {
							var pgErr *pgconn.PgError

							if errors.As(err, &pgErr) {
								log.Printf("authorizeReadRequestOnCollecton.go: could not check of collection member exists: %v", pgErr)
								util.Response(w, "something went wrong", http.StatusInternalServerError)
								return
							}
							log.Printf("authorizeReadRequestOnCollecton.go: could not check of collection member exists: %v", err)
							util.Response(w, "something went wrong", http.StatusInternalServerError)
							return
						}
						if val {
							body := newReadRequestOnCollectionDetails(collection.FolderID, int64(account_id), *payload)
							ctx := context.WithValue(r.Context(), "readRequestOnCollectionDetails", body)
							next.ServeHTTP(w, r.WithContext(ctx))
							return
						}
					}
					return
				}

				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					log.Printf("could not get collection member in authorizeReadRequestOnCollection.go... pgErr: %v", pgErr)
					util.Response(w, "something went wrong", http.StatusInternalServerError)
					return
				}

				log.Printf("could not get collection member in authorizeReadRequestOnCollection.go... err: %v", err)
				util.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
			// collection has been shared with user hence pass request detail
			body := newReadRequestOnCollectionDetails(collectionMember.CollectionID, int64(account_id), *payload)

			ctx := context.WithValue(r.Context(), "readRequestOnCollectionDetails", body)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
