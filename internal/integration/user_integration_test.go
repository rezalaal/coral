package integration_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/rezalaal/coral/internal/db"
    "github.com/rezalaal/coral/internal/models"
    "github.com/rezalaal/coral/internal/repository/postgres"
    "github.com/rezalaal/coral/internal/router"
)

func setupTestServer(t *testing.T) (*httptest.Server, func()) {
    dbConn, err := db.Connect()
    assert.NoError(t, err)

    // پاکسازی دیتابیس قبل از شروع تست
    _, err = dbConn.Exec("DELETE FROM users")
    assert.NoError(t, err)

    userRepo := postgres.NewUserPG(dbConn)

    // اگر رستوران هم بخوای تست کنی، مشابه همین بساز

    r := router.NewRouter(userRepo, nil)

    server := httptest.NewServer(r)

    return server, func() {
        dbConn.Close()
        server.Close()
    }
}

func TestUserIntegration(t *testing.T) {
    server, teardown := setupTestServer(t)
    defer teardown()

    // ساخت payload برای ایجاد کاربر
    userPayload := map[string]string{
        "name":          "Integration User",
        "mobile":        "09121234567",
        "password_hash": "hashedpass",
    }
    payloadBytes, _ := json.Marshal(userPayload)

    // ارسال درخواست POST به مسیر ایجاد کاربر
    resp, err := http.Post(server.URL+"/users/create", "application/json", bytes.NewReader(payloadBytes))
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)

    var createdUser models.User
    err = json.NewDecoder(resp.Body).Decode(&createdUser)
    assert.NoError(t, err)
    assert.Equal(t, "Integration User", createdUser.Name)
    resp.Body.Close()

    // درخواست لیست کاربران
    resp, err = http.Get(server.URL + "/users")
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    var users []models.User
    err = json.NewDecoder(resp.Body).Decode(&users)
    assert.NoError(t, err)
    assert.GreaterOrEqual(t, len(users), 1)
}

func TestCreateUser_InvalidJSON(t *testing.T) {
	server := test.NewTestServer()
	defer server.Close()

	body := `{"name": "علی", "mobile": "0935...` // JSON ناقص
	resp, err := http.Post(server.URL+"/users", "application/json", strings.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var msg string
	json.NewDecoder(resp.Body).Decode(&msg)
	assert.Contains(t, msg, "نامعتبر")
}

