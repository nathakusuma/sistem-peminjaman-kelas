package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/service"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/validator"
)

type UserHandler struct {
	svc service.IUserService
	val validator.IValidator
}

func NewUserHandler(svc service.IUserService, val validator.IValidator) *UserHandler {
	return &UserHandler{
		svc: svc,
		val: val,
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, errorpkg.ErrFailParseRequest())
		return
	}

	if err := h.val.ValidateStruct(req); err != nil {
		SendError(w, err)
		return
	}

	token, user, err := h.svc.Login(ctx, req.Username, req.Password)
	if err != nil {
		SendError(w, err)
		return
	}

	response := dto.LoginResponse{
		Token: token,
		User:  user,
	}

	SendJSON(ctx, w, http.StatusOK, response)
}
