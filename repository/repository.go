package repository

import (
	"context"
	"dev/master/entity"
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

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserById(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type KeyRepository interface {
	FindActiveUsersKeys(ctx context.Context, userId int64) (map[string]*entity.Key, error)
	GetKeyBySubscription(ctx context.Context, subscriptionId int64) (*entity.Key, error)
	InsertKey(ctx context.Context, key entity.Key) (*entity.Key, error)
}

type CountryRepository interface {
	FindCountryByName(ctx context.Context, name string) (*entity.Country, error)
	GetAllCountries(ctx context.Context) ([]*entity.Country, error)
	CreateCountry(ctx context.Context, country entity.Country) (*entity.Country, error)
}

type SubscriptionRepository interface {
	FindSubscriptions(ctx context.Context, userId int64, countryId int64, active bool) ([]*entity.Subscription, error)
	CreateSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	CreateTrialSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	GetAllSubscriptions(ctx context.Context) ([]*entity.Subscription, error)
}

type ProxyRepository interface {
	GetAllProxies(ctx context.Context) ([]*entity.Proxy, error)
	GetProxyById(ctx context.Context, id int64) (*entity.Proxy, error)
	CreateProxy(ctx context.Context, proxy entity.Proxy) (*entity.Proxy, error)
	GetRandomProxyByCountry(ctx context.Context, country_id int64) (*entity.Proxy, error)
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
