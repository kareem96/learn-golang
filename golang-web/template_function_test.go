package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct{
	Name string
}

func (myPage MyPage) SayHello(name string) string  {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Abdul"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage {
		Name: "Karem",
	})
}

func TestTemplateFunction(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage {
		Name: "Karem",
	})
}

func TestTemplateFunctionGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request)  {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage {
		Name: "Kareem",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateFunctionGlobalPipeline(writer http.ResponseWriter, request *http.Request)  {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func (value string) string {
			return "Hello " + value
		},
	})
	t = t.Funcs(map[string]interface{}{
		"upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage {
		Name: "Kareem",
	})
}

func TestTemplateFunctionGlobalPipeline(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}