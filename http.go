package main

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"cork/go13/action"
	"cork/go13/echoGinLogger"
)

type Index struct {
	templates *template.Template
}

func (t *Index) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var index = Index{
	templates: template.Must(template.New("index").Parse(`<!doctype html>
<html>
  <body>
	<pre>
	 {{range .files}}<a href="{{.path}}">{{.name}}</a>{{end}}
	</pre>
  </body>
</html>`)),
}

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

func startWebServer(config *action.Config) {
	router := echo.New()
	router.HideBanner = true
	router.Use(echoGinLogger.EchoLogger)
	router.Use(middleware.SecureWithConfig(middleware.SecureConfig{XFrameOptions: ""}))
	router.Use(middleware.Recover())

	router.Renderer = &index

	fs := http.Dir(config.Folder)

	router.GET("/:config", func(c echo.Context) error {
		file, err := fs.Open(c.Param("config"))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		return c.Stream(http.StatusOK, "text/plain", file)
	})
	router.GET("/", func(c echo.Context) error {
		f, err := fs.Open("")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error opening directory")
			return errors.New("Error opening directory")
		}

		dirs, err := f.Readdir(-1)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error reading directory")
			return errors.New("Error reading directory")
		}
		sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })

		var files []map[string]interface{}
		for _, d := range dirs {
			name := d.Name()
			if d.IsDir() {
				name += "/"
			}
			// name may contain '?' or '#', which must be escaped to remain
			// part of the URL path, and not indicate the start of a query
			// string or fragment.
			url := url.URL{Path: name}
			files = append(files, map[string]interface{}{"path": url.String(), "name": htmlReplacer.Replace(name)})
		}

		return c.Render(http.StatusOK, "index", map[string]interface{}{"files": files})
	})

	router.Start(":9587")
}
