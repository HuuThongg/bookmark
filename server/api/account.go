package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"go-bookmark/auth"
	"go-bookmark/db/sqlc"
	"go-bookmark/util"
	"log"
	"net/http"
	"sync"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgconn"
)

type signUp struct {
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type session struct {
	Account      sqlc.Account
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func newSession(account sqlc.Account, accessToken, refreshToken string) session {
	return session{
		Account:      account,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
func (s signUp) Validate(requestValidationChan chan error) error {
	err := validation.ValidateStruct(&s,
		validation.Field(&s.FullName, validation.Required.Error("name required"), validation.Length(1, 255).Error("name must be between 1 and 255 characters long")),
		validation.Field(&s.EmailAddress, validation.Required.Error("email address is equired"), is.Email.Error("email must be valid email address")),
		validation.Field(&s.Password, validation.Required.Error("password is required"), validation.Length(6, 1000).Error("Password must be at least 6 characters")),
	)
	requestValidationChan <- err
	return err
}
func (h *BaseHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	var req signUp

	if err := body.Decode(&req); err != nil {
		log.Printf("failed to decode request with error %v", err)
		ErrorDecodingRequest(w, err)
		return
	}

	requestValidationChan := make(chan error, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		req.Validate(requestValidationChan)
	}()
	wg.Wait()

	q := sqlc.New(h.db)
	emailExists, err := q.EmailExists(r.Context(), req.EmailAddress)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			log.Println(pgErr)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		} else {
			log.Println(err)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		}
	}

	if emailExists {
		log.Println("email already exists")
		util.Response(w, errors.New("email already exists").Error(), http.StatusConflict)
		return
	}

	var p string

	p, err = util.HashPassword(req.Password)
	if err != nil {
		log.Printf("failed to hash password with error: %s", err)
		util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
		return
	}

	arg := sqlc.NewAccountParams{
		Fullname:        req.FullName,
		Email:           req.EmailAddress,
		AccountPassword: p,
	}

	account, err := q.NewAccount(r.Context(), arg)
	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			log.Println(pgErr)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		} else {
			log.Println(err)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		}
	}

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Printf("failed to load config file with err: %v", err)
		util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
		return
	}

	accessToken, accessTokenPayload, err := auth.CreateToken(account.ID, time.Now().UTC(), config.Access_Token_Duration)
	if err != nil {
		log.Printf("failed to create access token with err: %v", err)
		util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, refreshTokenPayload, err := auth.CreateToken(account.ID, accessTokenPayload.IssuedAt, config.Refresh_Token_Duration)
	if err != nil {
		log.Printf("failed to create refresh token with err: %v", err)
		util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
		return
	}

	createAccountSessionParams := sqlc.CreateAccountSessionParams{
		RefreshTokenID: refreshTokenPayload.ID,
		AccountID:      account.ID,
		IssuedAt:       refreshTokenPayload.IssuedAt,
		Expiry:         refreshTokenPayload.Expiry,
		UserAgent:      "",
		ClientIp:       "",
	}
	_, err = q.CreateAccountSession(r.Context(), createAccountSessionParams)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			log.Printf("failed co create session with err: %s", pgErr.Message)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		} else {
			log.Printf("failed to create session with error: %s", err.Error())
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		}
	}

	account, err = q.GetAccount(r.Context(), account.ID)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			log.Printf("failed to get created account with err: %v", pgErr)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			log.Println("account not found")
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		} else {
			log.Printf("failed to get created account with err: %v", err)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		}
	}

	res := newSession(account, accessToken, refreshToken)

	util.JsonResponse(w, res)

}

// sign up with google

type continueWithGoogle struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

func (c continueWithGoogle) Validate(errChan chan error) error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error("name required")),
		validation.Field(&c.Email, validation.Required.Error("email required"), is.Email),
		validation.Field(&c.Picture, validation.Required.Error("profile picture required"), is.URL),
	)

	errChan <- err

	return err
}

func (h *BaseHandler) ContinueWithGoogle(w http.ResponseWriter, r *http.Request) {
	rBody := json.NewDecoder(r.Body)
	rBody.DisallowUnknownFields()

	var req continueWithGoogle

	if err := rBody.Decode(&req); err != nil {
		ErrorDecodingRequest(w, err)
		return
	}

	errChan := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		req.Validate(errChan)
	}()
	wg.Wait()

	q := sqlc.New(h.db)

	email := req.Email
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println(err)
		return
	}

	exisitingAccount, err := q.GetAccountByEmail(r.Context(), email)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.As(err, &pgErr):
			ErrorInternalServer(w, pgErr)
			return
		case errors.Is(err, sql.ErrNoRows):
			createAccountParams := sqlc.NewAccountParams{
				Fullname: req.Name,
				Email:    req.Email,
			}
			createAccount(createAccountParams, q, w, h, config, r.Context())
			return
		default:
			ErrorInternalServer(w, err)
			return
		}
	}
	if exisitingAccount != (sqlc.Account{}) {
		loginUser(exisitingAccount, w, h, config, r.Context())
	}
}

func createAccount(arg sqlc.NewAccountParams, q *sqlc.Queries, w http.ResponseWriter, h *BaseHandler, config util.Config, ctx context.Context) {
	account, err := q.NewAccount(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError

		switch {
		case errors.As(err, &pgErr):

			ErrorInternalServerError(w, pgErr)
			return
		default:
			ErrorInternalServerError(w, err)
			return
		}
	}

	if err := q.UpdateAccountEmailVerificationStatus(ctx, arg.Email); err != nil {
		log.Println(err)
		util.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	loginUser(account, w, h, config, ctx)
}
func loginUser(account sqlc.Account, w http.ResponseWriter, h *BaseHandler, config util.Config, ctx context.Context) {

	accessToken, accessTokenPayload, err := auth.CreateToken(account.ID, time.Now().UTC(), config.Access_Token_Duration)
	if err != nil {
		ErrorInternalServerError(w, err)
		return
	}

	q := sqlc.New(h.db)
	refreshToken, refreshTokenPayload, err := auth.CreateToken(account.ID, accessTokenPayload.IssuedAt, config.Refresh_Token_Duration)
	if err != nil {
		ErrorInternalServerError(w, err)
		return
	}

	createAccountSessionParams := sqlc.CreateAccountSessionParams{
		RefreshTokenID: refreshTokenPayload.ID,
		AccountID:      account.ID,
		IssuedAt:       refreshTokenPayload.IssuedAt,
		Expiry:         refreshTokenPayload.Expiry,
		UserAgent:      "",
		ClientIp:       "",
	}

	if _, err := q.CreateAccountSession(ctx, createAccountSessionParams); err != nil {
		ErrorInternalServerError(w, err)
		return
	}

	refreshTokenCookie := http.Cookie{
		Name:    "refreshTokenCookie",
		Value:   refreshToken,
		Path:    "/",
		Expires: refreshTokenPayload.Expiry,
		// Secure:   true,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}

	http.SetCookie(w, &refreshTokenCookie)

	newAccount, err := q.GetAccount(ctx, account.ID)
	if err != nil {
		ErrorInternalServerError(w, err)
		return
	}

	res := newSession(newAccount, accessToken, refreshToken)

	util.JsonResponse(w, res)
}

func (h *BaseHandler) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.db)

	accounts, err := q.GetAllAccounts(r.Context())
	if err != nil {
		log.Println(err)
		util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
		return
	}

	if len(accounts) == 0 {
		util.Response(w, errors.New("no accounts found").Error(), http.StatusNotFound)
		return
	}

	util.JsonResponse(w, accounts)
}
