// blockchainservice/main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/blockchain", func(c *gin.Context) {
		// Interagir com a blockchain
		c.JSON(http.StatusOK, gin.H{"blockchain": "interaction"})
	})

	router.Run(":8083")
}
