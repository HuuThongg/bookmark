package router

import (
	"go-bookmark/api"
	"go-bookmark/db/connection"

	cm "go-bookmark/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)
	h := api.NewBaseHandler(connection.ConnectDB())

	// public routes go here
	r.Route("/public", func(r chi.Router) {
		r.Get("/", h.Example)
		r.Post("/checkIfIsAuthenticated", h.CheckIfIsAuthenticated)
		r.Post("/continueWithGoogle", h.ContinueWithGoogle)
		r.Get("/getUserMessages", h.GetAllUserMessages)
		r.Post("/refreshToken", h.RefreshToken)
		r.Post("/requestResetPasswordLink", h.RequestResetPasswordLink)
		r.Patch("/updatePassword", h.UpdatePassword)
		r.Post("/uploadHeroImage", h.UploadHeroImage)
		r.Route("/account", func(r chi.Router) {
			r.Post("/", h.ContinueWithGoogle)
			r.Post("/create", h.NewAccount)
			r.Get("/getAllAccounts", h.GetAllAccounts)
			r.Post("/signin", h.SignIn)
		})
	})
	r.Route("/private", func(r chi.Router) {
		r.Use(cm.AuthenticateRequest())
		r.Route("/getLinksAndFolders/{accountID}/{folderID}", func(r chi.Router) {
			r.Use(cm.AuthorizeReadRequestOnCollection())
			r.Get("/", h.GetLinksAndFolders)
		})
		r.Route("/folder", func(r chi.Router) {
			r.Route("/create", func(r chi.Router) {
				r.Use(cm.AuthorizeCreateFolderRequest())
				r.Post("/", h.CreateFolder)
			})
			r.Post("/new-child-folder", h.CreateChildFolder)
			r.Patch("/star", h.StarFolders)
			r.Patch("/unstar", h.UnstarFolders)
			r.Patch("/rename", h.RenameFolder)
			r.Patch("/moveFoldersToTrash", h.MoveFoldersToTrash)
			r.Patch("/moveFolders", h.MoveFolders)
			r.Patch("/moveFoldersToRoot", h.MoveFoldersToRoot)
			r.Patch("/toggle-folder-starred", h.ToggleFolderStarred)
			r.Patch("/restoreFoldersFromTrash", h.RestoreFoldersFromTrash)
			r.Delete("/deleteFoldersForever", h.DeleteFoldersForever)
			r.Get("/getRootFoldersByUserID", h.GetRootFolders)
			r.Get("/getFolderChildren/{folderID}/{accountID}", h.GetFolderChildren)
			r.Get("/getFolderAncestors/{folderID}", h.GetFolderAncestors)
			r.Get("/searchFolders/{query}", h.SearchFolders)
		})
		r.Route("/link", func(r chi.Router) {
			r.Post("/add", h.AddLink)
			r.Patch("/rename", h.RenameLink)
			r.Patch("/move", h.MoveLinks)
			r.Patch("/moveLinksToTrash", h.MoveLinksToTrash)
			r.Patch("/restoreLinksFromTrash", h.RestoreLinksFromTrash)
			r.Delete("/deleteLinksForever", h.DeleteLinksForever)
			r.Get("/getRootLinks/{accountID}", h.GetRootLinks)
			r.Get("/get_folder_links/{accountID}/{folderID}", h.GetFolderLinks)
			r.Get("/searchLinks/{query}", h.SearchLinks)
		})
	})
	return r
}
