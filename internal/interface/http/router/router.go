package router

import (
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/handler"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/middleware"
)

func NewRouter(userHandler *handler.UserHandler, proposalHandler *handler.ProposalHandler) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK", "message": "Service is healthy"}`))
	})

	// User routes
	mux.HandleFunc("POST /api/v1/auth/login", userHandler.Login)
	mux.HandleFunc("POST /api/v1/auth/register", userHandler.Register)

	// Proposal routes
	mux.HandleFunc("POST /api/v1/proposals", proposalHandler.CreateProposal)
	mux.HandleFunc("GET /api/v1/proposals", proposalHandler.GetProposals)
	mux.HandleFunc("GET /api/v1/proposals/{id}", proposalHandler.GetProposalDetail)
	mux.HandleFunc("POST /api/v1/proposals/{id}/reply", proposalHandler.CreateReply)

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Sistem Peminjaman Kelas API", "version": "v1"}`))
	})

	// Apply middleware
	return middleware.Chain(mux, middleware.Logging, middleware.CORS)
}
