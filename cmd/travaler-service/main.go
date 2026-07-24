package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kadr/globe_express/config"
	api_handlers "github.com/kadr/globe_express/internal/travaler_service/application/handlers"
	trip_repository "github.com/kadr/globe_express/internal/travaler_service/domain/repository"
	trip_service "github.com/kadr/globe_express/internal/travaler_service/domain/service"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))
	cfg := config.MustLoad()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	db, err := sqlx.Connect("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	tripRepo := trip_repository.NewRepository(db, cfg.DBTimeout, logger)
	tripService := trip_service.NewService(tripRepo, logger)
	api := api_handlers.NewTravalerAPI(tripService, logger)
	logger.Info("register all handlers")
	api.RegisterHandlers()
	logger.Info("starting http server")
	api.Start(cfg.HttpAddress)
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	api.Shutdown(shutdownCtx)
}
