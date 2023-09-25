package endpoints

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Header struct {
    Title string
}

type Page struct {
	Header
}

type TemplateRenderer struct {
	tmpl *template.Template
}

func NewTemplateRenderer(tmpls *template.Template) TemplateRenderer {
	return TemplateRenderer{tmpls}
}

func (t TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func HandleIndex(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

func HandleHome(c echo.Context) error {
	return c.Render(200, "home.html", Page{
		Header: Header{
			Title: "TODO - Home",
		},
	})
}

func HandleSettings(c echo.Context) error {
	return c.Render(200, "settings.html", Page{
		Header: Header{
			Title: "TODO - Settings",
		},
	})
}

func HandleHelp(c echo.Context) error {
	return c.Render(200, "help.html", Page{
		Header: Header{
			Title: "TODO - Help",
		},
	})
}

