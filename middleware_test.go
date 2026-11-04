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

type Middleware struct {
	 http.Handler
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recive request")
	m.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recoder := httptest.NewRecorder()
	middleware := &Middleware{router}
	middleware.ServeHTTP(recoder, request)
	response := recoder.Result()

	byte,_ := io.ReadAll(response.Body)
	assert.Equal(t,"Middleware",string(byte))



}