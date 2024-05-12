// Exemplo de teste unitário para o serviço de autenticação (authservice/auth_test.go)
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginEndpoint(t *testing.T) {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Logged in!"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Logged in!")
}
