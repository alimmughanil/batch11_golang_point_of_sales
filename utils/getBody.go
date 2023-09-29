package utils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func GetBody(c echo.Context) map[string]interface{} {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return nil
	}

	return jsonBody
}
