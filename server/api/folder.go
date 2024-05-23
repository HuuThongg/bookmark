package api

import (
	"database/sql"
	"go-bookmark/db/sqlc"
	"go-bookmark/middleware"
	"go-bookmark/util"
	"log"
	"net/http"
	"sync"
)

func (h *BaseHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {

	requestBody := r.Context().Value("createFolderRequest").(*middleware.CreateFolderRequestBody).Body

	authorizedPayload := r.Context().Value("createFolderRequest").(*middleware.CreateFolderRequestBody).PayLoad

	queries := sqlc.New(h.db)

	if requestBody.FolderID != "null" {
		log.Printf("foldeId %v", requestBody.FolderID)
		util.CreateChildFolder(queries, w, r, requestBody.FolderName, requestBody.FolderID, authorizedPayload.AccountID)
		return
	}
	stringChan := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.RandomStringGenerator(stringChan)
	}()
	folderLabelChan := make(chan string, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.GenFolderLabel(folderLabelChan)
	}()

	folderID := <-stringChan
	folderLabel := <-folderLabelChan
	folderParams := sqlc.CreateFolderParams{
		FolderID:    folderID,
		FolderName:  requestBody.FolderName,
		SubfolderOf: sql.NullString{},
		AccountID:   authorizedPayload.AccountID,
		Path:        folderLabel,
		Label:       folderLabel,
	}
	folder, err := queries.CreateFolder(r.Context(), folderParams)
	if err != nil {
		ErrorPgError(w, err)
		return
	}
	rf := newReturnedFolder(folder)
	util.JsonResponse(w, rf)
	wg.Wait()
}
