package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	/*if err := router.SetTrustedProxies([]string{"colocar ip desejado"}); err != nil {
		panic(err)
	}*/
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Logged in!"})
	})

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Manejo de sinais de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	// Aqui você pode adicionar qualquer limpeza ou fechamento de recursos antes de sair
	os.Exit(0)
	router.Run(":8080")
}
