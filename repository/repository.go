package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
	"strings"
)

type postgresRepository struct {
	logger *zap.Logger
	conn   *sql.DB
}

func NewPostgresRepository(logger *zap.Logger) Repository {
	return &postgresRepository{
		logger: logger,
	}
}

func (r *postgresRepository) Connect(source string) error {
	conn, err := sql.Open("postgres", source)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}
	r.conn = conn

	return nil
}

func (r *postgresRepository) Initialize(schemePath string) error {
	file, err := os.ReadFile(schemePath)
	if err != nil {
		return err
	}

	queries := strings.Split(string(file), ";")
	for _, query := range queries {
		_, err := r.conn.Exec(query)
		if err != nil {
			r.logger.Warn("failed to execute query", zap.String("query", query), zap.Error(err))
		} else {
			r.logger.Info("executed query", zap.String("query", query))
		}
	}

	return nil
}

func (r *postgresRepository) Close() error {
	return r.conn.Close()
}
