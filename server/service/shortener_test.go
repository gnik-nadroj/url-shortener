package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/data_access"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestShortener(t *testing.T) {
	router := gin.Default()
	sessionStore, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", sessionStore))
	store := data_access.NewURLStore()
	userStore := data_access.NewUserStore()

	storedUser, err := userStore.GetByUsername("testuser")

	assert.NoError(t, err)
	
	router.POST("/shortener", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Set("user", storedUser.ID)
		shortener(ctx, store)
	})

	reqBody := struct {
		URL string `json:"url" binding:"required"`
	}{
		URL: "https://www.google.com",
	}

	jsonBody, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/shortener", bytes.NewBuffer(jsonBody))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}