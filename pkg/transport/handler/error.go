package handler

import (
	"errors"
	"net/http"
	"seriouspoop/greedygame/go-common/middleware"
	"seriouspoop/greedygame/pkg/svc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func writeErrorResponse(err error, r *http.Request, w http.ResponseWriter) {
	ctx := r.Context()
	if errors.Is(err, svc.ErrNoData) {
		middleware.WriteJsonHttpErrorResponse(ctx, w, http.StatusNoContent, err)
	} else if errors.Is(err, svc.ErrBadInput) || errors.Is(err, svc.ErrImportantFieldMissing) {
		middleware.WriteJsonHttpErrorResponse(ctx, w, http.StatusBadRequest, err)
	}
}

func grpcError(err error) error {
	if errors.Is(err, svc.ErrBadInput) || errors.Is(err, svc.ErrImportantFieldMissing) {
		return status.Error(codes.InvalidArgument, "invalid or empty arguments")
	} else if errors.Is(err, svc.ErrUnexpected) {
		return status.Error(codes.Unknown, "unexpected error at service")
	} else if errors.Is(err, svc.ErrDuplicateData) {
		return status.Error(codes.AlreadyExists, "already exists")
	}
	return status.Error(codes.Unknown, "unexpected error at service")
}
