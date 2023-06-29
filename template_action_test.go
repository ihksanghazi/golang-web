package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter,request *http.Request){
	t:= template.Must(template.ParseFiles(`./templates/if.gohtml`))
	t.ExecuteTemplate(writer,"if.gohtml", map[string]interface{}{
		"Title":"Judul",
		// "Name":"Sandy",
	})
}

func TestTemplateActionIf(t *testing.T){
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder:= httptest.NewRecorder()

	TemplateActionIf(recorder,request)

	body,_:= io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter,request *http.Request){
	t:= template.Must(template.ParseFiles(`./templates/comparator.gohtml`))
	t.ExecuteTemplate(writer,"comparator.gohtml", map[string]interface{}{
		"Title":"Judul",
		"FinalValue": 70,
	})
}

func TestTemplateActionOperator(t *testing.T){
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder:= httptest.NewRecorder()

	TemplateActionOperator(recorder,request)

	body,_:= io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter,request *http.Request){
	t:= template.Must(template.ParseFiles(`./templates/range.gohtml`))
	t.ExecuteTemplate(writer,"range.gohtml", map[string]interface{}{
		"Title":"Judul",
		"Hobbies": []string{
			"coding", "reading", "gaming",
		},
	})
}

func TestTemplateActionRange(t *testing.T){
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder:= httptest.NewRecorder()

	TemplateActionRange(recorder,request)

	body,_:= io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter,request *http.Request){
	t:= template.Must(template.ParseFiles(`./templates/address.gohtml`))
	t.ExecuteTemplate(writer,"address.gohtml", map[string]interface{}{
		"Title":"Judul",
		"Name": "Sandy",
		"Address": map[string]interface{}{
			"Street": "Jl Panjang Banget",
			"City": "Jakarta Selatan",
		},
	})
}

func TestTemplateActionWith(t *testing.T){
	request:= httptest.NewRequest(http.MethodGet,"http://localhost:8080",nil)
	recorder:= httptest.NewRecorder()

	TemplateActionWith(recorder,request)

	body,_:= io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}