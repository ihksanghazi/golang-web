package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter,request *http.Request){
	cookie:= new(http.Cookie)
	cookie.Name = "Sandy-Jigong"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer,cookie)
	fmt.Fprint(writer,"Success Create Cookie")
}

func getCookie(writer http.ResponseWriter,request *http.Request){
	cookie,err:= request.Cookie("Sandy-Jigong")
	if err != nil {
		fmt.Fprint(writer,"No Cookie")
	}else{
		fmt.Fprintf(writer, "Hello %s",cookie.Value)
	}
}

func TestCookie(t *testing.T){
	mux:= http.NewServeMux()
	mux.HandleFunc("/set-cookie",SetCookie)
	mux.HandleFunc("/get-cookie",getCookie)

	server:= http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err:=server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080/?name=sandy",nil)
	recorder:= httptest.NewRecorder()

	SetCookie(recorder,request)

	cookies:= recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s %s",cookie.Name,cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080/",nil)
	cookie:= new(http.Cookie)
	cookie.Name = "Sandy-Jigong"
	cookie.Value = "Sandy"
	request.AddCookie(cookie)
	
	recorder:= httptest.NewRecorder()

	getCookie(recorder,request)

	body,_:= io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}