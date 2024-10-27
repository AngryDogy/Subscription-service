package repository

import "dev/master/entity"

func (r *postgresRepository) GetProxies() ([]*entity.Proxy, error) {
	query := `SELECT * FROM proxy`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}

	proxies := make([]*entity.Proxy, 0)
	for rows.Next() {
		var p entity.Proxy
		err := rows.Scan(&p.Id, &p.Address, &p.CountryId)
		if err != nil {
			continue
		}
		proxies = append(proxies, &p)
	}
	return proxies, nil
}
