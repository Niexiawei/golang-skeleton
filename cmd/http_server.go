package cmd

import (
	"context"
	"fmt"
	"github.com/Niexiawei/golang-skeleton/config"
	"github.com/Niexiawei/golang-skeleton/database"
	"github.com/Niexiawei/golang-skeleton/internal/cvalid"
	"github.com/Niexiawei/golang-skeleton/logger"
	"github.com/Niexiawei/golang-skeleton/router"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var httServer = &cli.Command{
	Name:  "server",
	Usage: "启动http服务器",
	Action: func(context *cli.Context) error {
		httpServerBootstrap()
		httServerRun()
		return nil
	},
}

func httpServerBootstrap() {
	cvalid.SetupValid()
	database.SetupDatabase()
	database.SetupRedis()
}

func httServerRun() {
	ginEngine := router.SetupRouter()
	stopHttpServer := startHttpServer(ginEngine)
	logger.Logger.Info("server start successful ...")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	stopHttpServer()
}

func startHttpServer(r *gin.Engine) func() {
	addr := fmt.Sprintf(":%d", config.Instance.HttpServer.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	logger.Logger.Debug("http server address " + addr)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	shutdown := func() {
		logger.Logger.Info("Shutdown Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Logger.Error("Server Shutdown:", err)
		}
		logger.Logger.Info("Server exiting ...")
	}
	return shutdown
}
