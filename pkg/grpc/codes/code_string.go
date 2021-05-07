package codes

import (
	"google.golang.org/grpc/codes"
)

// StatusMessage represent string message for code
var StatusMessage = map[codes.Code]string{
	Success:          "Success",
	SuccessCreated:   "Success Create",
	SuccessNoContent: "Success",
	InvalidArgument:  "Invalid Parameter",
	Unauthorized:     "Unauthorized",
	Forbidden:        "Forbidden access",
	NotFound:         "Data not found",
	Cancelled:        "Request canceled",
	RequestTimeout:   "Request Timeout",
	InvalidToken:     "Invalid or expired token",
	InternalError:    "Error dari server",
}
