package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/mattuttis/prd-project/customer-service/internal/adapter/inbound/handler"
	"github.com/mattuttis/prd-project/customer-service/internal/adapter/outbound/persistence"
	"github.com/mattuttis/prd-project/customer-service/internal/application"
	"github.com/mattuttis/prd-project/customer-service/internal/config"
	"github.com/mattuttis/prd-project/customer-service/pkg/logging"
	"github.com/mattuttis/prd-project/customer-service/pkg/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	repo := persistence.NewInMemoryCustomerRepository()
	service := application.NewCustomerService(repo)
	customerHandler := handler.NewCustomerHandler(service)

	logger := logging.NewLogger()
	router := handler.NewRouter(customerHandler,
		prometheus.PrometheusMiddleware(),
		logging.Middleware(logger),
	)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	cfg := config.Load()
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	logger.Info("Starting server", "port", cfg.ServerPort)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("listen error", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	}

	logger.Info("Server exiting")
}
