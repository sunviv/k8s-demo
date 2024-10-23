package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewHttpServer(server *gin.Engine) *http.Server {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	log.Println("http server listening on", httpServer.Addr)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return httpServer
}
