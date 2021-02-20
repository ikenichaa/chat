package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"practice/golang/internal/config"
	"practice/golang/internal/domain/repository"
	"practice/golang/internal/domain/service"
	"practice/golang/internal/handler"
	"practice/golang/internal/logger"
	"practice/golang/internal/route"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	log = logger.Logger()
	e   = echo.New()
)

// var (
// 	workerService = service.NewWorkerService()
// 	hdlr          = handler.NewHandler(workerService)
// )

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	c, err := config.Init()
	if err != nil {
		log.Panicf("couldn't initialize config %v", err)
	}

	logger.InitLogger()
	gormLogger := logger.NewGormLogger()

	mySQLConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.MysqlUsername,
		c.MysqlPassword,
		c.MysqlAddress,
		c.MysqlPort,
		c.MysqlDB)

	mySQL, err := gorm.Open(mysql.Open(mySQLConn), &gorm.Config{
		Logger: gormLogger,
	})

	var (
		workerRepo    = repository.NewWorkerRepository(mySQL)
		workerService = service.NewWorkerService(service.WorkerConfig{
			Worker: workerRepo,
		})

		hdlr = handler.NewHandler(workerService)
	)

	if err != nil {
		log.Panicf("couldn't initialize database with connection %s: %v", mySQLConn, err)
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	route.InitRoute(e, hdlr)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", c.ServerPort)); err != nil {
			log.Info("shutting down the server")
		}
	}()

	sigWatcher := make(chan os.Signal)
	signal.Notify(sigWatcher, os.Interrupt, os.Kill)
	<-sigWatcher
	log.Info("catch signal, performing gracefully shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Error(err)
	}

}
