package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"tidak tersedia")
	})
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// fmt.Fprint(w, "Hello http Router")
		panic("internal server error")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/1", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	byte,_ := io.ReadAll(response.Body)
	assert.Equal(t,"tidak tersedia",string(byte))

}