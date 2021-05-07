package http

import (
	"context"
	"fmt"

	"github.com/hidayatullahap/evermos/pkg/grpc/codes"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm"
	grpcCode "google.golang.org/grpc/codes"
)

// Response struct
type Response struct {
	Code       grpcCode.Code          `json:"code"`
	Message    string                 `json:"message,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
	Pagination *Pagination            `json:"pagination,omitempty"`
	Errors     []string               `json:"errors,omitempty"`
	TraceID    string                 `json:"trace_id,omitempty"`
	Header     map[string]interface{} `json:"-"`
}

// Pagination struct
type Pagination struct {
	CurrentPage int32  `json:"current_page"`
	PageSize    int32  `json:"page_size"`
	TotalPage   int32  `json:"total_page"`
	TotalResult int32  `json:"total_result"`
	Next        string `json:"next,omitempty"`
	Prev        string `json:"prev,omitempty"`
}

// WithPagination set response with pagination
func (r *Response) WithPagination(c echo.Context, pagination Pagination) *Response {
	r.Pagination = &pagination
	page := r.Pagination.CurrentPage
	u := c.Request().URL
	if r.Pagination.NextPage() {
		q := u.Query()
		q.Set("page", fmt.Sprintf("%d", page+1))
		u.RawQuery = q.Encode()
		r.Pagination.Next = u.String()
	}
	if page > 1 {
		q := u.Query()
		q.Set("page", fmt.Sprintf("%d", page-1))
		u.RawQuery = q.Encode()
		r.Pagination.Prev = u.String()
	}
	return r
}

// JSON render response as JSON
func (r *Response) JSON(c echo.Context) error {
	for k, v := range r.Header {
		c.Response().Header().Set(k, fmt.Sprintf("%v,%v", c.Response().Header().Get(k), v))
	}
	return c.JSON(codes.HTTPStatusFromCode(r.Code), r)
}

// NextPage determine is available next page
func (p *Pagination) NextPage() bool {
	return (p.CurrentPage * p.PageSize) < p.TotalResult
}

// GetTraceID get trace ID from current context
func GetTraceID(ctx context.Context) string {
	trx := apm.TransactionFromContext(ctx)
	return trx.TraceContext().Trace.String()
}
