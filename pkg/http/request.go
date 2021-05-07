package http

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/module/apmhttp"
)

// Header struct for customizable header
type Header struct {
	Key   string
	Value string
}

// RequestWithTimeout create request to url with timeout
func RequestWithTimeout(ctx context.Context, method string, url string, bytesData []byte, duration time.Duration, header ...Header) (*http.Response, error) {
	var client = apmhttp.WrapClient(&http.Client{Timeout: duration})
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "application/json")
	for _, head := range header {
		req.Header.Set(head.Key, head.Value)
	}
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	return res, err
}

// Request create request to url
func Request(ctx context.Context, method string, url string, bytesData []byte, header ...Header) (*http.Response, error) {
	var client = apmhttp.WrapClient(http.DefaultClient)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "application/json")
	for _, head := range header {
		req.Header.Set(head.Key, head.Value)
	}
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	return res, err
}

func BindAndValidate(c echo.Context, r interface{}) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.Validate(r)
}
