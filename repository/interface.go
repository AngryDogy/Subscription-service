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
	FindUserKeys(ctx context.Context, userId int64) (map[string]*entity.Key, error)
}

type CountryRepository interface {
	FindCountryByName(ctx context.Context, name string) (*entity.Country, error)
}

type ProxyRepository interface {
	GetProxies() ([]*entity.Proxy, error)
}

type SubscriptionRepository interface {
	FindSubscriptionsByUserId(userId int64) ([]*entity.Subscription, error)
	CreateSubscription(ctx context.Context, userId, countryId int64) (*entity.Subscription, error)
	CreateTrialSubscription(ctx context.Context, userId, countryId int64) (*entity.Subscription, error)
}
