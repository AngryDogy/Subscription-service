package repository

import (
	"context"
	"dev/master/entity"
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
