package handlers

import "github.com/labstack/echo/v4"

// isHTMX returns true when the request was issued by htmx (HX-Request header).
func isHTMX(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}
