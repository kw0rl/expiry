package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func handleHealth(c *gin.Context) {
	response := healthResponse{
		Status:  "ok",
		Service: "expiry-api",
	}
	c.JSON(http.StatusOK, response)
}

func main() {
	address := os.Getenv("API_ADDR")
	if address == "" {
		address = "127.0.0.1:8081"
	}

	router := gin.Default()
	router.GET("/api/health", handleHealth)

	server := &http.Server{
		Addr:              address,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Expiry API listening on http://%s", address)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
