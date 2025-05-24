package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/database"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

func main() {
	config.GetEnv()
	log.NewLogger()

	db := database.NewPostgreSQLConn(
		config.GetEnv().DBHost,
		config.GetEnv().DBPort,
		config.GetEnv().DBUser,
		config.GetEnv().DBPass,
		config.GetEnv().DBName,
		config.GetEnv().DBSSLMode,
	)
	defer db.Close(context.Background())

	// Register the handler function for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Info(context.Background()).Msgf("Server is running on port %s", config.GetEnv().AppPort)

	err := http.ListenAndServe(":"+config.GetEnv().AppPort, nil)
	if err != nil {
		log.Fatal(context.Background()).Err(err).Msg("Failed to start server")
	}
}
