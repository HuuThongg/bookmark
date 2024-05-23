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
		r.Post("/refreshToken", h.RefreshToken)
		r.Post("/requestResetPasswordLink", h.RequestResetPasswordLink)
		r.Patch("/updatePassword", h.UpdatePassword)
		r.Route("/account", func(r chi.Router) {
			r.Post("/", h.ContinueWithGoogle)
			r.Post("/create", h.NewAccount)
			//r.Get("/getAllAccounts",h.GetAllAcounts)
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
		})
	})
	return r
}
