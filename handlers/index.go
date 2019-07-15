package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetIndex render the index page of the app
func GetIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{})
}
