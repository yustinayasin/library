package utils

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGrpcError(err error) error {
	if err == nil {
		return nil
	}

	var grpcStatusCode codes.Code

	switch {
	case IsNotFoundError(err):
		grpcStatusCode = codes.NotFound
	case IsBadRequestError(err):
		grpcStatusCode = codes.InvalidArgument
	case IsInternalServerError(err):
		grpcStatusCode = codes.Internal
	default:
		grpcStatusCode = codes.Unknown
	}

	return status.Errorf(grpcStatusCode, "Internal error: %v", err.Error())
}

func IsNotFoundError(err error) bool {
	return fmt.Sprintf("%v", err) == "Not Found"
}

func IsBadRequestError(err error) bool {
	return fmt.Sprintf("%v", err) == "Bad Request"
}

func IsInternalServerError(err error) bool {
	return fmt.Sprintf("%v", err) == "Internal Error"
}
