package repository

import (
	"context"
	"dev/master/entity"
)

func (r *postgresRepository) FindCountryByName(ctx context.Context, name string) (*entity.Country, error) {
	query := `SELECT * FROM country WHERE name = $1;`

	var country entity.Country
	err := r.conn.QueryRowContext(ctx, query, name).Scan(&country.Id, &country.Name)
	if err != nil {
		return nil, err
	}

	return &country, nil
}
