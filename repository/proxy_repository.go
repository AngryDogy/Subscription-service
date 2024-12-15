package repository

import (
	"context"
	"dev/master/entity"
	"fmt"

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

func (r *postgresRepository) GetProxyById(ctx context.Context, id int64) (*entity.Proxy, error) {
	query := `SELECT id, address, country_id FROM proxy WHERE id=$1`

	var proxy entity.Proxy
	err := r.conn.QueryRow(ctx, query, fmt.Sprintf("%d", id)).Scan(&proxy.Id, &proxy.Address, &proxy.CountryId)
	if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (r *postgresRepository) CreateProxy(ctx context.Context, proxy entity.Proxy) (*entity.Proxy, error) {
	query := `INSERT INTO proxy(address, country_id) VALUES($1, $2) RETURNING id, address, country_id`

	var createdProxy entity.Proxy
	err := r.conn.QueryRow(ctx, query, proxy.Address, proxy.CountryId).Scan(&createdProxy.Id, &createdProxy.Address, &createdProxy.CountryId)
	if err != nil {
		return nil, err
	}

	return &createdProxy, nil
}

func (r *postgresRepository) GetRandomProxyByCountry(ctx context.Context, countryId int64) (*entity.Proxy, error) {
	query := `SELECT id, address, country_id FROM proxy WHERE country_id=$1`

	rows, err := r.conn.Query(ctx, query, countryId)
	if err != nil {
		return nil, err
	}

	proxies, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[entity.Proxy])
	if err != nil {
		return nil, err
	}

	return proxies[0], nil
}
