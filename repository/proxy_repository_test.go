package repository

import (
	"context"
	"dev/master/entity"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_postgresRepository_GetAllProxies(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	proxies, err := repository.GetAllProxies(context.TODO())

	require.Nil(t, err)

	require.Equal(t, 1, len(proxies))
	require.Equal(t, entity.Proxy{Id: 1, Address: "address", CountryId: 1}, *proxies[0])
}
