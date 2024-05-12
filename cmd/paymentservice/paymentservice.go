// paymentservice/main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/payment", func(c *gin.Context) {
		// Processar pagamento
		c.JSON(http.StatusOK, gin.H{"payment": "processed"})
	})

	router.Run(":8082")
}
