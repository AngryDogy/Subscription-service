package repository

import (
	"context"
	"dev/master/entity"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_postgresRepository_FindActiveUsersKeys(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	keys, err := repository.FindActiveUsersKeys(context.TODO(), 1)

	require.Nil(t, err)

	require.Equal(t, 1, len(keys))
	require.Equal(t,
		entity.Key{
			Id:             1,
			Data:           []byte("data"),
			KeyType:        entity.Text,
			SubscriptionId: 1,
			ProxyId:        1},
		*keys["Russia"],
	)
}

func Test_postgresRepository_getKeyBySubscription(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	key, err := repository.GetKeyBySubscription(context.TODO(), 1)

	require.Nil(t, err)
	require.NotNil(t, key)

	require.Equal(t,
		entity.Key{
			Id:             1,
			Data:           []byte("data"),
			KeyType:        entity.Text,
			SubscriptionId: 1,
			ProxyId:        1},
		*key,
	)
}

func Test_addKey(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	key, err := repository.InsertKey(context.TODO(), entity.Key{
		Data:           []byte("ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTptMElrR3FCOGREalZqRFFNZjdENVdY@158.160.6.79:1416/?outline=1"),
		KeyType:        entity.Text,
		SubscriptionId: 1,
		ProxyId:        1,
		IdInProxy:      "1",
	})

	require.Nil(t, err)
	require.Equal(t, entity.Key{
		Id:             2,
		Data:           []byte("ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTptMElrR3FCOGREalZqRFFNZjdENVdY@158.160.6.79:1416/?outline=1"),
		KeyType:        entity.Text,
		SubscriptionId: 1,
		ProxyId:        1,
		IdInProxy:      "1",
	}, *key)
}

func Test_getKey(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	key, err := repository.GetKeyBySubscription(context.TODO(), 1)

	require.Nil(t, err)
	require.NotNil(t, key)

	require.Equal(t, entity.Key{
		Id:             2,
		Data:           []byte("ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTptMElrR3FCOGREalZqRFFNZjdENVdY@158.160.6.79:1416/?outline=1"),
		KeyType:        entity.Text,
		SubscriptionId: 1,
		ProxyId:        1,
	}, *key)

}
