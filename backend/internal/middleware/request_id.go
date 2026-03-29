package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	RequestIDHeader = "X-Request-ID"				// for request header
	RequestIDKey = "request_id"						// for request context
)


func RequestID() echo.MiddlewareFunc{													// a middleware takes a handler and returns a handler and a handler takes a request context and returns error
	return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get(RequestIDHeader)

			if requestID == "" {
				requestID = uuid.New().String()
			}

			c.Set(RequestIDKey, requestID)
			c.Response().Header().Set(RequestIDHeader, requestID)

			return next(c)
		}
	}
}



func GetRequestID(c echo.Context) string{
	if requestID, ok := c.Get(RequestIDKey).(string) ; ok{
		return requestID
	}

	return ""
}

