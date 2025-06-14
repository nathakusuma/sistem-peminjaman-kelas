package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/validator"
)

func SendError(w http.ResponseWriter, err error) {
	var respErr *errorpkg.ResponseError

	typePrefix := config.GetEnv().BackendURL + "/errors"
	w.Header().Set("Content-Type", "application/problem+json")

	if errors.As(err, &respErr) {
		w.WriteHeader(respErr.Status)
		json.NewEncoder(w).Encode(
			respErr.
				WithTypePrefix(typePrefix).
				WithInstance(config.GetEnv().BackendURL), // TODO: pass request URL as parameter
		)
		return
	}

	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(
			errorpkg.ErrValidation().
				WithValidationErrors(validationErr).
				WithTypePrefix(typePrefix).
				WithInstance(config.GetEnv().BackendURL), // TODO: pass request URL as parameter
		)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(
		errorpkg.ErrInternalServer().WithTypePrefix(typePrefix),
	)
}

func SendJSON(ctx context.Context, w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error(ctx).Err(err).Msg("failed to encode response to JSON")
		SendError(w, errorpkg.ErrInternalServer())
	}
}
