package xerr

import (
	"github.com/asjard/asjard/core/status"
	"google.golang.org/grpc/codes"
)

var (
	ErrInvalidName = func() error {
		return status.Error(codes.InvalidArgument, "invalid name, name is must and max length 20")
	}
)
