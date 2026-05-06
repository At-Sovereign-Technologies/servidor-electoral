package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	root            *template.Template
	pageTemplates   map[string]*template.Template
	sharedTemplates map[string]struct{}
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	basePath, err := filepath.Abs("templates")
	if err != nil {
		return nil, err
	}

	baseLayout := filepath.Join(basePath, "base.html")
	partialsPattern := filepath.Join(basePath, "partials", "*.html")
	pagesPattern := filepath.Join(basePath, "pages", "*.html")

	root := template.New("root")
	if _, err := root.ParseFiles(baseLayout); err != nil {
		return nil, err
	}

	if _, err := root.ParseGlob(partialsPattern); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	pageFiles, err := filepath.Glob(pagesPattern)
	if err != nil {
		return nil, err
	}

	pageTemplates := make(map[string]*template.Template, len(pageFiles))
	for _, page := range pageFiles {
		name := filepath.Base(page)
		clone, err := root.Clone()
		if err != nil {
			return nil, err
		}

		if _, err := clone.ParseFiles(page); err != nil {
			return nil, err
		}

		pageTemplates[fmt.Sprintf("pages/%s", name)] = clone
	}

	sharedTemplates := make(map[string]struct{})
	for _, tmpl := range root.Templates() {
		sharedTemplates[tmpl.Name()] = struct{}{}
	}

	return &TemplateRenderer{
		root:            root,
		pageTemplates:   pageTemplates,
		sharedTemplates: sharedTemplates,
	}, nil
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if pageTmpl, ok := t.pageTemplates[name]; ok {
		return pageTmpl.ExecuteTemplate(w, name, data)
	}

	if _, ok := t.sharedTemplates[name]; ok {
		return t.root.ExecuteTemplate(w, name, data)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, "Template not found: "+name)
}
