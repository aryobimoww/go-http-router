package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServerFile(t *testing.T) {
	router := httprouter.New()
	directory,err := fs.Sub(resources,"resources")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	byte,_ := io.ReadAll(response.Body)
	assert.Equal(t,"Hello txt",string(byte))



}