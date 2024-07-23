package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BadRequestError bad request
var BadRequestError = status.Errorf(
	codes.InvalidArgument,
	"Bad request error",
)

// InternalError internal error
var InternalError = status.Errorf(
	codes.Internal,
	"Internal server error",
)

// NotFoundError object is not found
var NotFoundError = status.Errorf(
	codes.NotFound,
	"User not found",
)
