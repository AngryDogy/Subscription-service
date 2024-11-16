package repository

import (
	"context"
	"dev/master/entity"
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
	CreateUser(ctx context.Context, userId int64) (*entity.User, error)
	FindUserById(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type KeyRepository interface {
	FindActiveUsersKeys(ctx context.Context, userId int64) (map[string]*entity.Key, error)
}

type CountryRepository interface {
	FindCountryByName(ctx context.Context, name string) (*entity.Country, error)
	GetAllCountries(ctx context.Context) ([]*entity.Country, error)
}

type ProxyRepository interface {
	GetAllProxies(ctx context.Context) ([]*entity.Proxy, error)
}

type SubscriptionRepository interface {
	FindSubscriptions(ctx context.Context, userId int64, countryId int64, active bool) ([]*entity.Subscription, error)
	CreateSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	CreateTrialSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
}
