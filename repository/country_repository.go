package repository

import (
	"context"
	"dev/master/entity"

	"github.com/jackc/pgx/v5"
)

func (r *postgresRepository) FindCountryByName(ctx context.Context, name string) (*entity.Country, error) {
	query := `SELECT id, name FROM country WHERE name = $1;`

	var country entity.Country
	err := r.conn.QueryRow(ctx, query, name).Scan(&country.Id, &country.Name)
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *postgresRepository) GetAllCountries(ctx context.Context) ([]*entity.Country, error) {
	query := `SELECT id, name FROM country`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[entity.Country])
}

func (r *postgresRepository) CreateCountry(ctx context.Context, country entity.Country) (*entity.Country, error) {
	query := `INSERT INTO country (id, name) VALUES ($1, $2) RETURNING id, name`

	var createdCountry entity.Country
	err := r.conn.
		QueryRow(ctx, query, country.Id, country.Name).
		Scan(&createdCountry.Id, &createdCountry.Name)

	if err != nil {
		return nil, err
	}

	return &createdCountry, nil
}
