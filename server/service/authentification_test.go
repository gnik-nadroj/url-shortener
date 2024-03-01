package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"server/data_access"
	"server/model"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    cmd := exec.Command("docker", "run", "--name", "test-redis", "-p", "6379:6379", "-d", "redis")
    err := cmd.Run()
    if err != nil {
        panic(err)
    }

    code := m.Run()

    cmd = exec.Command("docker", "rm", "-f", "test-redis")
    err = cmd.Run()
    if err != nil {
        panic(err)
    }

    os.Exit(code)
}

func TestSignUp(t *testing.T) {
	router := gin.Default()
	store := data_access.NewUserStore()
	
	router.POST("/signup", func(ctx *gin.Context) {
		signUp(ctx, store)
	})

	user := model.User{
		Username: "testuser",
		Password: "testpassword",
	}

	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "User registered successfully")
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	store := data_access.NewUserStore()

	sessionStore, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", sessionStore))
	
	router.POST("/login", func(ctx *gin.Context) {
		login(ctx, store)
	})

	user := model.User{
		Username: "testuser",
		Password: "testpassword",
	}

	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "You are logged in")
}

