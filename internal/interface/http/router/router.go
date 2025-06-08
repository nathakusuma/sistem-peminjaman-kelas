package router

import (
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/handler"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/middleware"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/jwt"
)

func NewRouter(
	userHandler *handler.UserHandler,
	proposalHandler *handler.ProposalHandler,
	jwtService jwt.IJwt,
) http.Handler {
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
	mux.Handle("POST /api/v1/proposals",
		middleware.Chain(
			http.HandlerFunc(proposalHandler.CreateProposal),
			middleware.RequireAuth(jwtService),
			middleware.RequireRole(enum.UserRoleStudent),
		),
	)
	mux.Handle("GET /api/v1/proposals",
		middleware.Chain(
			http.HandlerFunc(proposalHandler.GetProposals),
			middleware.RequireAuth(jwtService),
		),
	)
	mux.Handle("GET /api/v1/proposals/{id}",
		middleware.Chain(
			http.HandlerFunc(proposalHandler.GetProposalDetail),
			middleware.RequireAuth(jwtService),
		),
	)
	mux.Handle("POST /api/v1/proposals/{id}/replies",
		middleware.Chain(
			http.HandlerFunc(proposalHandler.CreateReply),
			middleware.RequireAuth(jwtService),
			middleware.RequireRole(enum.UserRoleAdmin),
		),
	)

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

	// Apply global middleware
	return middleware.Chain(mux, middleware.Logging, middleware.CORS)
}
