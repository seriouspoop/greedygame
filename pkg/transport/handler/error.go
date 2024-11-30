package handler

import (
	"errors"
	"net/http"
	"seriouspoop/greedygame/go-common/middleware"
	"seriouspoop/greedygame/pkg/svc"
)

func writeErrorResponse(err error, r *http.Request, w http.ResponseWriter) {
	ctx := r.Context()
	if errors.Is(err, svc.ErrNoData) {
		middleware.WriteJsonHttpErrorResponse(ctx, w, http.StatusNoContent, err)
	} else if errors.Is(err, svc.ErrBadInput) || errors.Is(err, svc.ErrImportantFieldMissing) {
		middleware.WriteJsonHttpErrorResponse(ctx, w, http.StatusBadRequest, err)
	}
}
