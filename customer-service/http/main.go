package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	init_app "customer-service/internal"
	pkg_config "customer-service/internal/pkg/config"
	util_db "customer-service/util/db"
	util_http "customer-service/util/http"
	util_http_middleware "customer-service/util/http/middleware"
	util_logger "customer-service/util/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := pkg_config.LoadConfig()
	pgDB := util_db.NewPostgresDB(config.PostgresHost, config.PostgresUser, config.PostgresPassword, config.PostgresDbName, config.PostgresPort)

	router := util_http.NewHTTPServer(config.AppEnv)
	util_logger.InitLogger(config.AppEnv, config.AppName, config.AppLogPath)

	router.Use(
		gin.Logger(),
		gin.Recovery(),
		util_http_middleware.CorsHandlerMiddleware(),
		util_http_middleware.ErrorHandlerMiddleware(),
	)

	init_app.InitializeApp(router, pgDB)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			util_logger.Fatal(ctx, err.Error())
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
