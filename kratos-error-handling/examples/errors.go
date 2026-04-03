```go
package errors

import (
	"errors"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
)

// Error is a status error.
type Error struct {
	Status
	cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v cause = %v", e.Code, e.Reason, e.Message, e.Metadata, e.cause)
}

// New returns an error object for the code, message.
func New(code int, reason, message string) *Error {
	return &Error{
		Status: Status{
			Code:    int32(code),
			Message: message,
			Reason:  reason,
		},
	}
}

// Code returns the http code for an error.
func Code(err error) int {
	if err == nil {
		return 200 //nolint:mnd
	}
	return int(FromError(err).Code)
}

// Reason returns the reason for a particular error.
func Reason(err error) string {
	if err == nil {
		return ""
	}
	return FromError(err).Reason
}

// FromError try to convert an error to *Error.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if !ok {
		return New(500, "", err.Error())
	}
	ret := New(
		httpstatus.FromGRPCCode(gs.Code()),
		"",
		gs.Message(),
	)
	for _, detail := range gs.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			ret.Reason = d.Reason
			return ret.WithMetadata(d.Metadata)
		}
	}
	return ret
}
```