package main

import (
	"context"
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/database"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/repository"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/handler"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/interface/http/router"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/bcrypt"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/jwt"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/mail"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/validator"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/usecases/proposal"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/usecases/user"
)

func main() {
	// Initialize configuration and logger
	config.GetEnv()
	log.NewLogger()

	// Initialize database connection
	db := database.NewPostgreSQLConn(
		config.GetEnv().DBHost,
		config.GetEnv().DBPort,
		config.GetEnv().DBUser,
		config.GetEnv().DBPass,
		config.GetEnv().DBName,
		config.GetEnv().DBSSLMode,
	)
	defer db.Close(context.Background())

	// Initialize dependencies
	bcryptInstance := bcrypt.GetBcrypt()
	jwtInstance := jwt.NewJwt(config.GetEnv().JwtExpireDuration, config.GetEnv().JwtSecretKey)
	mailer := mail.NewMailDialer()
	validatorInstance := validator.NewValidator()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	proposalRepo := repository.NewProposalRepository(db)

	// Initialize services
	userService := user.NewUserService(userRepo, bcryptInstance, jwtInstance)
	proposalService := proposal.NewProposalService(proposalRepo, mailer)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService, validatorInstance)
	proposalHandler := handler.NewProposalHandler(proposalService, validatorInstance)

	// Initialize routes
	mux := router.NewRouter(userHandler, proposalHandler, jwtInstance)

	log.Info(context.Background()).Msgf("Server is running on port %s", config.GetEnv().BackendPort)

	err := http.ListenAndServe(":"+config.GetEnv().BackendPort, mux)
	if err != nil {
		log.Fatal(context.Background()).Err(err).Msg("Failed to start server")
	}
}
