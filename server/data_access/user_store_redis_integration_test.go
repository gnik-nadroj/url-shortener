// Import packages
package data_access

import (
	"context"
	"os"
	"os/exec"
	"server/model"
	"testing"

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

var user = &model.User{
	ID:       "123",
	Username: "alice",
	Password: "secret",
}

func TestInsertUser(t *testing.T) {
	store := NewUserStore()

	t.Run("test user", func(t *testing.T) {
		err := store.Insert(user)

		assert.NoError(t, err)

		ctx := context.Background()
		result, err := store.client.HGetAll(ctx, usersPrefix+user.ID).Result()

		assert.NoError(t, err)

		assert.Equal(t, user.Username, result[usernameKey])
		assert.Equal(t, user.Password, result[passwordKey])
	})
}

func TestGetByID(t *testing.T) {
	store := NewUserStore()
	t.Run("valid user", func(t *testing.T) {
		u, err := store.GetByID(user.ID)

		assert.NoError(t, err)

		assert.Equal(t, user.Username, u.Username)
		assert.Equal(t, user.Password, u.Password)
	})
}
