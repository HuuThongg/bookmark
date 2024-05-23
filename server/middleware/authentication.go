package middleware

import (
	"context"
	"errors"
	"go-bookmark/auth"
	"go-bookmark/db"
	"go-bookmark/util"
	"log"
	"net/http"
	"strings"
)

func AuthenticateRequest() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			payload, err := getAndVerifyToken(r)
			if err != nil {
				log.Println(err)
				util.Response(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if payload != nil {
				account, err := db.ReturnAccount(r.Context(), payload.AccountID)
				if err != nil {
					log.Println(err)
					util.Response(w, errors.New("unauthorized").Error(), http.StatusUnauthorized)
					return
				}
				if payload.IssuedAt.Unix() != account.LastLogin.Unix() {
					err := errors.New("invalid token")
					log.Println(err)
					util.Response(w, err.Error(), http.StatusUnauthorized)
					return
				}
			}
			ctx := context.WithValue(r.Context(), "payload", payload)
			next.ServeHTTP(w, r.WithContext(ctx))

		}
		return http.HandlerFunc(fn)
	}
}

func getAndVerifyToken(r *http.Request) (*auth.PayLoad, error) {
	token := r.Header.Get("authorization")
	log.Printf("token embeded in Header %v", token)
	if token == "" {
		log.Println("token is empty!")
		return nil, errors.New("token is empty")
	}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		log.Println("bearer not in proper format")
		err := errors.New("bearer token is not in proper format")
		return nil, err
	}
	token = splitToken[1]
	refreshTokenCookie, err := r.Cookie("refreshTokenCookie")
	// if err != nil {
	// 	return nil, errors.New("refreshTokenCookie is not available")
	// }
	log.Printf("refreshTokenCookie: %v", refreshTokenCookie)

	payload, err := auth.VerifyToken(token)
	if err != nil {
		log.Println("payload is not verified")
		return nil, err
	}
	log.Printf("payload %v", payload)
	return payload, nil
}
