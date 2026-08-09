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
	api_handlers "github.com/kadr/globe_express/internal/order_service/application/handlers"
	repository "github.com/kadr/globe_express/internal/order_service/domain/repository"
	order_service "github.com/kadr/globe_express/internal/order_service/domain/service/order"
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
	middleware "github.com/kadr/globe_express/internal/shared/auth"
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
	orderRepo := repository.NewOrderRepository(db, cfg.DBTimeout, logger)
	productRepo := repository.NewProductRepository(db, cfg.DBTimeout, logger)

	productService := product_service.NewService(productRepo, logger)
	orderService := order_service.NewService(orderRepo, productService, logger)
	api := api_handlers.NewProductAPI(orderService, productService, logger)
	logger.Info("register all handlers")
	api.RegisterMiddleware(
		middleware.RequestID,
		middleware.Logger(logger),
		middleware.Recovery,
		middleware.GetUserFromReq(cfg),
		middleware.ExecuteResponse,
	)
	api.RegisterHandlers()
	logger.Info("starting http server")
	api.Start(cfg.HttpAddressOrder)
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	api.Shutdown(shutdownCtx)
}
