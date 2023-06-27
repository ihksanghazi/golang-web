package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sayHello(writer http.ResponseWriter,request *http.Request){
	name:=request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(writer, "Hello")
	}else{
		fmt.Fprintf(writer, "Hello %s", name)
	}

}

func TestQueryParameter(t *testing.T){
	request:= httptest.NewRequest("GET","http://localhost:8080/hello?name=Sandy",nil)
	recorder:= httptest.NewRecorder()

	sayHello(recorder,request)

	response:= recorder.Result()
	body,_:=io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter,request *http.Request){
	firstname:=request.URL.Query().Get("firstname")
	lastname:=request.URL.Query().Get("lastname")

	fmt.Fprintf(writer,"hello %s %s",firstname,lastname)
}

func TestMultipleQueryParameter(t *testing.T) {
	request:= httptest.NewRequest("GET","http://localhost:8080/hello?firstname=Sandy&lastname=Ihksan",nil)
	recorder:= httptest.NewRecorder()

	MultipleQueryParameter(recorder,request)

	response:= recorder.Result()
	body,_:=io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter,request *http.Request){
	query:= request.URL.Query()
	names:=query["name"]
	fmt.Fprint(writer,strings.Join(names," "))
}

func TestMultipleParameterValues(t *testing.T) {
	request:= httptest.NewRequest("GET","http://localhost:8080/hello?name=Nur&name=Sandy&name=Ihksan",nil)
	recorder:= httptest.NewRecorder()

	MultipleParameterValues(recorder,request)

	response:= recorder.Result()
	body,_:=io.ReadAll(response.Body)

	fmt.Println(string(body))
}