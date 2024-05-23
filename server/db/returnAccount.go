package db

import (
	"context"
	"go-bookmark/db/connection"
	"go-bookmark/db/sqlc"
)

func ReturnAccount(ctx context.Context, accountID int64) (*sqlc.Account, error) {
	q := sqlc.New(connection.ConnectDB())
	account, err := q.GetAccount(ctx, accountID)
	if err != nil {
		return &sqlc.Account{}, err
	}
	return &account, nil
}
