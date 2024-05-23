package api

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"go-bookmark/db/sqlc"
	"go-bookmark/mailjet"
	"go-bookmark/util"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type requestResetPasswordLinkRequest struct {
	Email string `json:"email"`
}

func (r requestResetPasswordLinkRequest) validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, is.Email))
}

func (h *BaseHandler) RequestResetPasswordLink(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)

	body.DisallowUnknownFields()

	var req requestResetPasswordLinkRequest

	err := body.Decode(&req)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("failed to decode request body with err: %v", e)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		} else {
			log.Printf("failed to decode request body with err: %v", err)
			util.Response(w, errors.New("something went wrong").Error(), http.StatusInternalServerError)
			return
		}
	}

	err = req.validate()
	if err != nil {
		log.Printf("request validation error: %v", err)
		util.Response(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := sqlc.New(h.db)
	account, err := q.GetAccountByEmail(r.Context(), req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("not found email in database")
			mail := mailjet.NewAccountNotFoundEmail(req.Email)
			mail.SendAccountNotFoundEmail()
			util.Response(w, "reset password link has been sent", http.StatusOK)
			return
		} else {
			log.Printf("could not get accout by email: %v", err)
			util.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	token := uuid.NewString()

	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	params := sqlc.CreatePasswordResetTokenParams{
		AccountID:   account.ID,
		TokenHash:   encodedToken,
		TokenExpiry: time.Now().UTC().Add(15 * time.Minute),
	}

	_, err = q.CreatePasswordResetToken(r.Context(), params)
	if err != nil {
		util.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	// create new password reset token mail and send password reset mail
	mail := mailjet.NewPasswordResetTokenMail(account.Fullname, account.Email, token)
	mail.SendPasswordResetEmail()
	util.Response(w, "reset password link has been sent", http.StatusOK)
}
