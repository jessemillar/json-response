package jsonresp

import "github.com/labstack/echo"

// Response makes RESTful responses JSON friendly
type Response struct {
	Response string `json:"message"`
}

// Create returns an Echo error message in proper JSON format
func Create(c echo.Context, httpStatus int, message string) error {
	response := &Response{
		Response: message,
	}

	return c.JSON(httpStatus, *response)
}
