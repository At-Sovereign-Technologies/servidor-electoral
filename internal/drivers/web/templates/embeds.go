package templates

// embeded filesystem for HTML templates
import (
	"embed"
	"html/template"
)

//go:embed *.html
var templateFS embed.FS

func LoadTemplate(patterns ...string) (*template.Template, error) {
	template, err := template.ParseFS(templateFS, patterns...)
	if err != nil {
		return nil, err
	}

	return template, nil
}
