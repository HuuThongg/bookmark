package db

import (
	"context"
	"go-bookmark/db/connection"
	"go-bookmark/db/sqlc"
)

func CheckIfCollectionMemberExists(ctx context.Context, folderID string, accountID int64) (bool, error) {
	value, err := sqlc.New(connection.ConnectDB()).CheckIfCollectionMemberWithCollectionAndMemberIdsExists(ctx, sqlc.CheckIfCollectionMemberWithCollectionAndMemberIdsExistsParams{CollectionID: folderID, MemberID: accountID})
	if err != nil {
		return false, err
	}
	return value, nil
}
