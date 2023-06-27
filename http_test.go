package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "hello World")
}

func TestHttp (t *testing.T){
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/hello",nil)
	recorder := httptest.NewRecorder()

	helloHandler(recorder,request)

	response:= recorder.Result()
	body,_:=io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}