package service

import (
	"net/http"
	"net/http/httptest"
	"server/data_access"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestGetAllShortenedURLs(t *testing.T) {
	router := gin.Default()
	sessionStore, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", sessionStore))
	store := data_access.NewURLStore()
	userStore := data_access.NewUserStore()

	storedUser, err := userStore.GetByUsername("testuser")

	assert.NoError(t, err)
	
	router.GET("/stats", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Set("user", storedUser.ID)
		getAllShortenUrl(ctx, store)
	})

	req, _ := http.NewRequest("GET", "/stats", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}