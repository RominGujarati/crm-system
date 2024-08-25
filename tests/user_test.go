package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"crm-system/api/controllers"
	"crm-system/api/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func TestGetUsers(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateUser(t *testing.T) {
	r := setupRouter()
	user := map[string]string{
		"name":     "John Doe",
		"email":    "john@example.com",
		"company":  "ABC Corp",
		"status":   "active",
		"notes":    "New user",
	}
	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
