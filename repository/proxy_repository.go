package repository

import (
	"context"
	"dev/master/entity"
	"github.com/jackc/pgx/v5"
)

func (r *postgresRepository) GetAllProxies(ctx context.Context) ([]*entity.Proxy, error) {
	query := `SELECT id, address, country_id FROM proxy`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[entity.Proxy])
}
