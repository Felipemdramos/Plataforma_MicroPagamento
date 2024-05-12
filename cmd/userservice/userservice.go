// userservice/main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		// Buscar e retornar informações do usuário
		c.JSON(http.StatusOK, gin.H{"users": "user data"})
	})

	router.Run(":8081")
}
