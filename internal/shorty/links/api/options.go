package api

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func HealthCheck(c *routing.Context) error {
	if c.Request.URL.Path == "/health-check" {
		c.Abort()
	}
	return nil
}

func Encoding(c *routing.Context) error {
	c.Response.Header().Add("Content-Type", "application/json; charset=utf-8")
	return nil
}
