package repository

import (
	"context"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type Repository interface {
	Connect(source string) error
	Close() error
	Initialize(schemePath string) error
	UserRepository
	KeyRepository
	CountryRepository
	ProxyRepository
	SubscriptionRepository
}

type postgresRepository struct {
	logger *zap.Logger
	conn   *pgx.Conn
}

func NewPostgresRepository(logger *zap.Logger) Repository {
	return &postgresRepository{
		logger: logger,
	}
}

func (r *postgresRepository) Connect(source string) error {
	conn, err := pgx.Connect(context.TODO(), source)
	if err != nil {
		return err
	}

	if err := conn.Ping(context.TODO()); err != nil {
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
		_, err := r.conn.Exec(context.TODO(), query)
		if err != nil {
			r.logger.Warn("failed to execute query", zap.String("query", query), zap.Error(err))
		} else {
			r.logger.Info("executed query", zap.String("query", query))
		}
	}

	return nil
}

func (r *postgresRepository) Close() error {
	return r.conn.Close(context.TODO())
}
