package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/mpgarate/justread/controllers"
	"html/template"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Funcs: []template.FuncMap{
			{
				"add": func(a, b int) int {
					return a + b
				},
			},
		},
	}))

	m.Use(sessions.Sessions("reader_session", nil))

	m.Get("/", controllers.HomeController)

	m.Get("/read/*", controllers.ReadController)

	fmt.Println("Listening on port 8989")
	err := http.ListenAndServe(":8989", m)
	if err != nil {
		fmt.Println(err)
	}
}
