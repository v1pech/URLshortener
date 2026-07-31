package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/v1pech/URLshortener/config"
	"github.com/v1pech/URLshortener/internal/logs"
	"github.com/v1pech/URLshortener/internal/server/handlers"
	"github.com/v1pech/URLshortener/internal/server/mw_components"
	"github.com/v1pech/URLshortener/internal/storage"
)

func main() {
	config, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	logger, logFile := logs.Init(config.LogPath, config.LogLevel)
	logger.Info("starting url shortener")
	start_date := time.Now()
	defer func() {
		logFile.Close()
	}()

	db, err := storage.SetupStorage(config.Storage.Path)
	if err != nil {
		logger.Error("could not setup storage", "storage config: ",config.Storage, "error", err)
		panic(err)
	}
	defer db.Database.Close()

	logger.Info("initializing server", "adress", config.Server.Address)

	router := chi.NewRouter()
	router.Use(mw_components.InitLoggerMiddleware(logger))

	router.Get("/alias/{alias}", handlers.NewGetAlias(logger, db))
	router.Get("/{alias}", handlers.NewRedirect(logger, db))

	router.Route("/alias", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			config.Server.User: config.Server.Password}))
		r.Post("/", handlers.NewPostAlias(logger, db))
		r.Delete("/{alias}", handlers.NewDeleteAlias(logger, db))
	})

	server := &http.Server{
		Addr:         config.Server.Address,
		Handler:      router,
		ReadTimeout:  time.Duration(config.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.Timeout) * time.Second,
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start server!", "error", err.Error())
			panic(err)
		}
	}()

	logger.Info("server started successfully")

	<-done
	err = server.Shutdown(context.Background())
	if err != nil {
		logger.Error("failed to shutdown server!", "error", err.Error())
	} else {
		logger.Info("server stopped successfully")
	}
	logs.Stop(logger, config.LogPath, start_date)

}
