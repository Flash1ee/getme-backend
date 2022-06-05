package echo_adapter

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	hf "getme-backend/internal/pkg/handler/handler_interfaces"
)

type Renderer struct {
	templates *template.Template
	debug     bool
	location  string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug
	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.templates = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func WrapMiddleware(m func(http.Handler) http.Handler) hf.HMiddlewareFunc {
	return func(next hf.Handler) hf.Handler {
		return hf.HandlerFunc(func(c echo.Context) (err error) {
			m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.SetRequest(r)
				c.SetResponse(echo.NewResponse(w, c.Echo()))
				err = next.ServeHTTP(c)
			})).ServeHTTP(c.Response(), c.Request())
			return
		})
	}
}
func WrapMiddlewareToFunc(m func(http.Handler) http.Handler) hf.HFMiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.SetRequest(r)
				c.SetResponse(echo.NewResponse(w, c.Echo()))
				err = next(c)
			})).ServeHTTP(c.Response(), c.Request())
			return
		}
	}
}
