package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emmanuerl/vaultly/pkg/api"
	"github.com/emmanuerl/vaultly/pkg/api/middlewares"
	"github.com/emmanuerl/vaultly/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	db, err := config.ConnectDB(env)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := config.NewApp(db)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middlewares.HttpErrorHandler)

	r.Mount("/", api.WalletRoutes(app))

	addr := fmt.Sprintf(":%d", env.PORT)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Printf("%s listening on %s\n", env.ServiceName, addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
