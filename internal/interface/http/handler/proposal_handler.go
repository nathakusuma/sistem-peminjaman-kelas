package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/service"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/dto"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/validator"
)

type ProposalHandler struct {
	svc service.IProposalService
	val validator.IValidator
}

func NewProposalHandler(svc service.IProposalService, val validator.IValidator) *ProposalHandler {
	return &ProposalHandler{
		svc: svc,
		val: val,
	}
}

func (h *ProposalHandler) CreateProposal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateProposalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, errorpkg.ErrFailParseRequest())
		return
	}

	if err := h.val.ValidateStruct(req); err != nil {
		SendError(w, err)
		return
	}

	err := h.svc.CreateProposal(ctx, &req)
	if err != nil {
		SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProposalHandler) GetProposals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var pageReq dto.PaginationRequest
	if err := json.NewDecoder(r.Body).Decode(&pageReq); err != nil {
		SendError(w, errorpkg.ErrFailParseRequest())
		return
	}

	if err := h.val.ValidateStruct(pageReq); err != nil {
		SendError(w, err)
		return
	}

	proposals, pagination, err := h.svc.GetProposals(ctx, &pageReq)
	if err != nil {
		SendError(w, err)
		return
	}

	var resp struct {
		Proposals  []*dto.ProposalResponse `json:"proposals"`
		Pagination *dto.PaginationResponse `json:"pagination"`
	}

	resp.Proposals = proposals
	resp.Pagination = pagination

	SendJSON(ctx, w, http.StatusOK, resp)
}

func (h *ProposalHandler) GetProposalDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		SendError(w, errorpkg.ErrFailParseRequest())
		return
	}

	proposal, err := h.svc.GetProposalDetail(ctx, id)
	if err != nil {
		SendError(w, err)
		return
	}

	SendJSON(ctx, w, http.StatusOK, proposal)
}

func (h *ProposalHandler) CreateReply(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateReplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, errorpkg.ErrFailParseRequest())
		return
	}

	if err := h.val.ValidateStruct(req); err != nil {
		SendError(w, err)
		return
	}

	err := h.svc.CreateReply(ctx, &req)
	if err != nil {
		SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
