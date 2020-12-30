package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"os"

	"github.com/labstack/echo/v4"
)

var config = func() map[string]string {
	config, err := func() (map[string]string, error) {
		configfile, err := os.Open("config.json")
		if err != nil {
			return map[string]string{}, err
		}
		jsondata, err := ioutil.ReadAll(configfile)
		if err != nil {
			return map[string]string{}, err
		}
		data := make(map[string]string)
		err = json.Unmarshal(jsondata, &data)
		if err != nil {
			return map[string]string{}, err
		}
		return data, nil
	}()
	if err != nil {
		os.Exit(-1)
	}
	return config
}()

//Template wrapper
type Template struct {
	templates *template.Template
}

//Render template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.Static("/static", "public")
	e.GET("/", func(c echo.Context) error {
		return c.Render(
			200,
			"index.html",
			map[string]interface{}{
				"PageTitle": config["site.title"],
			},
		)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
