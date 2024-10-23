package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sunviv/k8s-demo/config"
	"github.com/sunviv/k8s-demo/internal/repository"
	"github.com/sunviv/k8s-demo/internal/repository/dao"
	"github.com/sunviv/k8s-demo/internal/service"
	"github.com/sunviv/k8s-demo/internal/web"
	"github.com/sunviv/k8s-demo/internal/web/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := gin.Default()
	server.Use(sessions.Sessions("ssid", cookie.NewStore([]byte("secret"))))
	db := initDB()
	server.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to k8s-demo!")
	})
	server.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	userHdl := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(dao.NewUserDao(db))))
	userHdl.RegisterRoutes(server)
	httpServer := web.NewHttpServer(server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalln(fmt.Errorf("http server shutdown: %v", err))
	}
	log.Println("Server exiting")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DB.DSN))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
