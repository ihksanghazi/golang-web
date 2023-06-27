package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T){
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server:= http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err:=server.ListenAndServe()
	
	if err!= nil{
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	mux:= http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hellow World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Images")
	})

	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Thumbnail")
	})

	server:= http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err:= server.ListenAndServe()
	if err!= nil{
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.URL)
	}

	server:= http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err:=server.ListenAndServe()
	if err!= nil{
		panic(err)
	}
}