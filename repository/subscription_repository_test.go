package repository

import (
	"context"
	"dev/master/entity"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func Test_postgresRepository_CreateSubscription(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	expirationDateTime := time.Now().Add(time.Hour * 168).UTC().Truncate(time.Second)

	subscription, err := repository.CreateSubscription(context.TODO(), entity.Subscription{
		UserId:             1,
		CountryId:          1,
		ExpirationDateTime: expirationDateTime,
	})

	require.Nil(t, err)
	require.NotNil(t, subscription)

	require.Equal(t,
		entity.Subscription{
			Id:                 subscription.Id,
			UserId:             1,
			CountryId:          1,
			ExpirationDateTime: expirationDateTime,
		},
		*subscription)
}

func Test_postgresRepository_FindSubscriptions(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	subscriptions, err := repository.FindSubscriptions(context.TODO(), 1, 1, true)

	require.Nil(t, err)
	require.Equal(t, 5, len(subscriptions))

	subscriptions, err = repository.FindSubscriptions(context.TODO(), 1, -1, true)

	require.Nil(t, err)
	require.Equal(t, 6, len(subscriptions))

	subscriptions, err = repository.FindSubscriptions(context.TODO(), 1, -1, false)

	require.Nil(t, err)
	require.Equal(t, 7, len(subscriptions))
}

func Test_postgresRepository_CreateTrialSubscription(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	expirationDateTime := time.Now().Add(time.Hour * 168).UTC().Truncate(time.Second)

	subscription, err := repository.CreateTrialSubscription(context.TODO(), entity.Subscription{
		UserId:             1,
		CountryId:          1,
		ExpirationDateTime: expirationDateTime,
	})

	require.Nil(t, err)
	require.NotNil(t, subscription)

	require.Equal(t,
		entity.Subscription{
			Id:                 subscription.Id,
			UserId:             1,
			CountryId:          1,
			ExpirationDateTime: expirationDateTime,
		},
		*subscription)

	user, err := repository.FindUserById(context.TODO(), 1)

	require.Nil(t, err)
	require.NotNil(t, user)
	require.True(t, user.HadTrial)
}
