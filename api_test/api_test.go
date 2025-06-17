package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	routes "github.com/Sirawithon11/Golang_Project/route_and_controller"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestUserRoute function for testing the /users route.
func TestCreateUserRoute(t *testing.T) {
	app := fiber.New()

	routes.SetupUserRoutes(app)
	// Define test cases
	tests := []struct {
		description  string
		requestBody  routes.Todo
		expectStatus int
	}{
		{
			description:  "Valid Body",
			requestBody:  routes.Todo{Id: 1, Success: false, Body: "test1"},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid Body",
			requestBody:  routes.Todo{Id: 2, Success: false, Body: ""},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "http://localhost:4000/users/", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
}
