package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
)

func TestMadlibHandlerWithExternalCalls(t *testing.T) {
	r := gin.Default()
	r.GET("/madlib", madlibHandler)

	req, err := http.NewRequest("GET", "/madlib", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body := rr.Body.String()
	assert.Contains(t, body, "It was a ")
	assert.Contains(t, body, "day. I went downstairs to see if I could ")
	assert.Contains(t, body, "dinner. I asked, 'Does the stew need fresh ")
}

func TestMadlibHandlerWithoutExternalCalls(t *testing.T) {
	fetchWordOrigin := fetchWord
	fetchWord = func(part string) (string, error) {
		return "test", nil
	}

	r := gin.Default()
	r.GET("/madlib", madlibHandler)

	req, err := http.NewRequest("GET", "/madlib", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Contains(t, rr.Body.String(), "It was a test day. I went downstairs to see if I could test dinner. I asked, 'Does the stew need fresh test?'")

	fetchWord = fetchWordOrigin
}

func TestMadlibHandlerNegative(t *testing.T) {
	fetchWordOrigin := fetchWord
	fetchWord = func(part string) (string, error) {
		return "", errors.New("test error")
	}

	r := gin.Default()
	r.GET("/madlib", madlibHandler)

	req, err := http.NewRequest("GET", "/madlib", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	assert.Contains(t, rr.Body.String(), "Failed to fetch word(s)")

	fetchWord = fetchWordOrigin
}
