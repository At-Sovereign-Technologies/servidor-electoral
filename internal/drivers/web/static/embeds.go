package static

// embeded filesystem for HTML templates
import (
	"embed"
)

//go:embed *
var StaticFS embed.FS
