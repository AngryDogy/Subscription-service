package repository

import (
	"context"
	"dev/master/entity"
)

func (r *postgresRepository) FindUserKey(ctx context.Context, userID, countryID int64) (*entity.Key, error) {
	query := `SELECT * FROM key WHERE subscription_id=
                        (SELECT id FROM subscription WHERE user_id=$1 AND country_id=$2)`

	var key entity.Key
	err := r.conn.QueryRowContext(ctx, query, userID, countryID).Scan(
		&key.Id,
		&key.Data,
		&key.KeyType,
		&key.SubscriptionId)
	if err != nil {
		return nil, err
	}

	return &key, nil
}
