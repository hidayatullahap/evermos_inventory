package http

import (
	netHttp "net/http"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/hidayatullahap/evermos/pkg"
	"github.com/hidayatullahap/evermos/pkg/grpc/codes"
	"github.com/hidayatullahap/evermos/pkg/http"
	"github.com/labstack/echo/v4"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
	idTrans "gopkg.in/go-playground/validator.v9/translations/id"
)

var translator ut.Translator

func setValidator(c *echo.Echo) {
	i := id.New()
	uni := ut.New(i, i)

	translator, _ = uni.GetTranslator("id")
	v := validator.New()
	_ = idTrans.RegisterDefaultTranslations(v, translator)
	c.Validator = &pkg.CustomValidator{Validator: v}
}

func ErrorHandler(c *echo.Echo) {
	setValidator(c)
	c.HTTPErrorHandler = func(err error, c echo.Context) {
		traceID := http.GetTraceID(c.Request().Context())

		// Validation Error
		if errs, ok := err.(validator.ValidationErrors); ok {
			var message []string

			translated := errs.Translate(translator)
			for _, v := range translated {
				message = append(message, v)
			}

			resp := http.Response{
				Code:    codes.InvalidArgument,
				Message: codes.StatusMessage[codes.InvalidArgument],
				Errors:  message,
				TraceID: traceID,
			}
			_ = resp.JSON(c)
			return
		}

		// GRPC Error
		if st, ok := status.FromError(err); ok {
			resp := http.Response{
				Code:    st.Code(),
				Message: codes.StatusMessage[st.Code()],
				Errors:  []string{st.Message()},
				TraceID: traceID,
			}

			// Handle Cancelled context
			if st.Code() == grpcCodes.Canceled {
				resp = http.Response{
					Code:    codes.RequestTimeout,
					Message: codes.StatusMessage[codes.RequestTimeout],
					Errors:  []string{codes.StatusMessage[codes.RequestTimeout]},
					TraceID: traceID,
				}
			}

			var errDetails []string
			if len(st.Details()) > 0 {
				for _, detail := range st.Details() {
					switch t := detail.(type) {
					case *errdetails.BadRequest:
						for _, violation := range t.GetFieldViolations() {
							errDetails = append(errDetails, violation.GetDescription())
						}
					}
				}

				resp.Errors = errDetails
			}
			_ = resp.JSON(c)
			return
		}

		if httpErr, ok := err.(*echo.HTTPError); ok {
			if httpErr.Code != int(codes.InternalError) {
				code := grpcCodes.Code(httpErr.Code)
				if code == netHttp.StatusMethodNotAllowed {
					code = codes.Forbidden
				}

				res := http.Response{
					Code:    code,
					Message: httpErr.Message.(string),
					Errors:  []string{err.Error()},
					TraceID: traceID,
				}

				_ = res.JSON(c)
				return
			}
		}

		res := http.Response{
			Code:    codes.InternalError,
			Message: codes.StatusMessage[codes.InternalError],
			Errors:  []string{err.Error()},
			TraceID: traceID,
		}

		_ = res.JSON(c)
		return
	}
}
