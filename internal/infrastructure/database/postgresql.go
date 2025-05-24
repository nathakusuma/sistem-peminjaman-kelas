package database

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

var (
	db   *pgx.Conn
	once sync.Once
)

func NewPostgreSQLConn(host, port, user, pass, dbName, sslMode string) *pgx.Conn {
	once.Do(func() {
		url := "postgresql://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbName + "?sslmode=" + sslMode
		conn, err := pgx.Connect(context.Background(), url)
		if err != nil {
			log.Fatal(context.Background()).Err(err).Msg("Fail to connect to PostgreSQL")
			return
		}

		log.Info(context.Background()).Msg("Connected to PostgreSQL")
		db = conn
	})

	return db
}
