package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type ErrorDetails struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
type Error struct {
	ErrorDetails `json:"error"`
}

func Convert(err error) (httpStatus int, errorData *Error) {
	errorData = new(Error)
	s := status.Convert(err)
	httpStatus = HTTPStatusFromCode(s.Code())
	errorData.Code = int(s.Code())
	errorData.Message = s.Message()
	errorData.Status = s.Code().String()
	errorData.Details = s.Details()
	return
}

func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
