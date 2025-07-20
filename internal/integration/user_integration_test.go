package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rezalaal/coral/internal/integration"
	"github.com/rezalaal/coral/internal/models"
)

func TestUserIntegration(t *testing.T) {
	server, teardown := integration.SetupTestServer(t)
	defer teardown()

	// ساخت کاربر
	userPayload := map[string]string{
		"name":          "Integration User",
		"mobile":        "09121234567",
		"password_hash": "hashedpass",
	}
	payloadBytes, _ := json.Marshal(userPayload)

	resp, err := http.Post(server.URL+"/users/create", "application/json", bytes.NewReader(payloadBytes))
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdUser models.User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	assert.NoError(t, err)
	assert.Equal(t, "Integration User", createdUser.Name)
	resp.Body.Close()

	// دریافت لیست کاربران
	resp, err = http.Get(server.URL + "/users")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(users), 1)
}
