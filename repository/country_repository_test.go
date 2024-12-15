package repository

import (
	"context"
	"dev/master/entity"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_postgresRepository_GetAllCountries(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	countries, err := repository.GetAllCountries(context.TODO())

	require.Nil(t, err)

	require.Equal(t, 1, len(countries))
	require.Equal(t, entity.Country{Id: 1, Name: "Russia"}, *countries[0])
}

func Test_postgresRepository_FindCountryByName(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	country, err := repository.FindCountryByName(context.TODO(), "Russia")

	require.Nil(t, err)

	require.Equal(t, entity.Country{Id: 1, Name: "Russia"}, *country)
}

func Test_postgresRepository_CreateCountry(t *testing.T) {
	repository := &postgresRepository{logger: nil}
	err := repository.Connect(os.Getenv("DATABASE_URL"))

	require.Nil(t, err)
	defer repository.Close()

	country, err := repository.CreateCountry(context.TODO(), entity.Country{
		Id:   2,
		Name: "Netherlands",
	})

	require.Nil(t, err)

	require.Equal(t, entity.Country{Id: 2, Name: "Netherlands"}, *country)
}
