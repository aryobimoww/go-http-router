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

func TestNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"tidak boleh")
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// fmt.Fprint(w, "Hello http Router")
		panic("internal server error")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	byte,_ := io.ReadAll(response.Body)
	assert.Equal(t,"tidak boleh",string(byte))

}