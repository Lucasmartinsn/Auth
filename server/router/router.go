package router

import (
	Login "authentication-system/handles"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func SecordMain() {
	router := gin.New()

	router.POST("/api/autentication", Login.Verificalogin)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowedMethods:   []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal((http.ListenAndServe("0.0.0.0:5000", handler)))
}
