package repository

import (
	"context"
	"dev/master/entity"
	"github.com/jackc/pgx/v5"
)

func (r *postgresRepository) FindUserKey(ctx context.Context, userId int64) (map[string]*entity.Key, error) {
	query := `select c.name, k.id, k.key_type, s.id, k.proxy_id, k.data
				from key k join subscription s 
				on k.subscription_id = s.id join country c
				on s.country_id = c.id
				where s.user_id = $1 and s.expiration_date > now()`

	rows, err := r.conn.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]*entity.Key)
	var country, keyType, data string
	var id, subscriptionId, proxyId int64
	_, err = pgx.ForEachRow(rows, []any{&country, &id, &keyType, &subscriptionId, &proxyId, &data}, func() error {
		keys[country] = &entity.Key{
			Id:             id,
			Data:           []byte(data),
			KeyType:        entity.KeyTypeFromString(keyType),
			SubscriptionId: subscriptionId,
			ProxyId:        proxyId,
		}
		return nil
	})

	return keys, nil
}
