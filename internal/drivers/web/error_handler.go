package web

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTMLHTTPErrorHandler(err error, c echo.Context) {
	var code = http.StatusInternalServerError
	var message = "Error interno del servidor"

	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code

		if msg, ok := he.Message.(string); ok {
			message = msg
		}
	}

	log.Printf("[ERROR] %d %s | path=%s | err=%v",
		code,
		message,
		c.Request().URL.Path,
		err,
	)

	// Avoid double write
	if c.Response().Committed {
		return
	}

	data := map[string]interface{}{
		"Title":   fmt.Sprintf("Error %d", code),
		"Code":    code,
		"Message": message,
	}

	// If HTMX request → return partial
	if isHTMX(c) {
		_ = c.Render(code, "partials/error.html", data)
		return
	}

	// Full page
	if err := c.Render(code, "pages/error.html", data); err != nil {
		c.String(http.StatusInternalServerError, message)
	}
}
