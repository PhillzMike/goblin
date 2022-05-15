package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var ac AuthController

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ac = NewAuthController("test")

	r := gin.Default()
	r.POST("/v1/auth/register", ac.Register)

	params := url.Values{}
	params.Add("first_name", "Chinonso")
	params.Add("last_name", "Okoli")
	params.Add("email", "leokingluthers@gmail.com")
	params.Add("password", "password")
	params.Add("confirm_password", "password")

	encoded := params.Encode()

	req, err := http.NewRequest(http.MethodPost, "/v1/auth/register", strings.NewReader(encoded))
	if err != nil {
		t.Fatalf("couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
