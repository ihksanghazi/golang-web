package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter,request *http.Request){
	name:= request.URL.Query().Get("name")

	if name == ""{
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer,"Name is Empty")
	}else{
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer,"Hello %s",name)
	}
}

func TestResponseCode(t *testing.T) {
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080/?name=Sandy",nil)
	recorder:= httptest.NewRecorder()
	
	ResponseCode(recorder,request)

	response:= recorder.Result()
	body,_:= io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}