package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"seriouspoop/greedygame/go-common/logging"

	"go.uber.org/zap"
)

const (
	AcceptHeader      = "Accept"
	ContentTypeHeader = "Content-Type"
	ContentTypeJson   = "application/json"
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteJsonHttpErrorResponse(ctx context.Context, w http.ResponseWriter, statusCode int, err error) {
	logger := logging.New(zap.DebugLevel)
	if err == nil {
		w.WriteHeader(statusCode)
		return
	}

	errMsg := ErrorMessage{Code: statusCode, Message: err.Error()}
	response, err := json.Marshal(errMsg)
	if err != nil {
		logger.Error("JSON marshal failed", zap.Error(err))
	}
	err = BuildHTTPResponse(w, response, ContentTypeJson, statusCode)

	if err != nil {
		// error while building http response
		// maybe writer was closed
		logger.Error("HTTP reponse writing failed", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func WriteJsonHttpResponse(ctx context.Context, w http.ResponseWriter, statusCode int, serviceRes interface{}) {
	if serviceRes != nil {
		response, err := json.Marshal(serviceRes)
		if err != nil {
			WriteJsonHttpErrorResponse(ctx, w, http.StatusInternalServerError, errors.New("marshal error"))
			return
		}

		err = BuildHTTPResponse(w, response, ContentTypeJson, statusCode)
		if err != nil {
			WriteJsonHttpErrorResponse(ctx, w, http.StatusInternalServerError, errors.New("write error"))
			return
		}
	} else {
		w.WriteHeader(statusCode)
	}
}

func BuildHTTPResponse(w http.ResponseWriter, data []byte, contentType string, statusCode int) error {
	w.Header().Set(ContentTypeHeader, contentType)
	w.WriteHeader(statusCode)

	// Write to response
	_, err := w.Write(data)
	return err
}
