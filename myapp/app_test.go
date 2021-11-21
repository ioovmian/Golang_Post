package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("hello foo!! World", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/bar?name=bart", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("hello foo!! bart", string(data))
}
