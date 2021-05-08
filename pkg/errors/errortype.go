package errors

import (
	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InvalidArgument invalid params
func InvalidArgument(message string, details map[string]string) error {
	st := status.New(codes.InvalidArgument, message)
	fields := make([]*errdetails.BadRequest_FieldViolation, 0)
	for k, v := range details {
		field := &errdetails.BadRequest_FieldViolation{
			Field:       k,
			Description: v,
		}
		fields = append(fields, field)
	}
	st, _ = st.WithDetails(&errdetails.BadRequest{
		FieldViolations: fields,
	})
	return st.Err()
}

// Unauthenticated not Unauthenticated
func Unauthenticated(message string) error {
	st := status.New(codes.Unauthenticated, message)
	return st.Err()
}

//PermissionDenied not access permission
func PermissionDenied(message string) error {
	st := status.New(codes.PermissionDenied, message)
	return st.Err()
}

// NotFound 404
func NotFound(message string) error {
	var (
		st *status.Status
	)
	st = status.New(codes.NotFound, message)
	return st.Err()
}

// ResourceExhausted 429
func ResourceExhausted(message string, details map[string]string, seconds int64) error {
	st := status.New(codes.ResourceExhausted, message)
	violations := make([]*errdetails.QuotaFailure_Violation, 0)
	for k, v := range details {
		violation := &errdetails.QuotaFailure_Violation{
			Subject:     k,
			Description: v,
		}
		violations = append(violations, violation)
	}
	st, _ = st.WithDetails(&errdetails.QuotaFailure{
		Violations: violations,
	}, &errdetails.RetryInfo{
		RetryDelay: &duration.Duration{Seconds: seconds},
	})
	return st.Err()
}

// Internal server internal error 500
func Internal(message string) error {
	st := status.New(codes.Internal, message)
	return st.Err()
}

// Unavailable service unavailable 503
func Unavailable(message string) error {
	st := status.New(codes.Unavailable, message)
	return st.Err()
}
