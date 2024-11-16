package repository

import (
	"context"
	"dev/master/entity"
	"github.com/jackc/pgx/v5"
)

func (r *postgresRepository) FindActiveUsersKeys(ctx context.Context, userId int64) (map[string]*entity.Key, error) {
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
	var country, keyType string
	var data []byte
	var id, subscriptionId, proxyId int64
	_, err = pgx.ForEachRow(rows, []any{&country, &id, &keyType, &subscriptionId, &proxyId, &data}, func() error {
		keys[country] = &entity.Key{
			Id:             id,
			Data:           data,
			KeyType:        entity.KeyTypeFromString(keyType),
			SubscriptionId: subscriptionId,
			ProxyId:        proxyId,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *postgresRepository) getKeyBySubscription(ctx context.Context, subscriptionId int64) (*entity.Key, error) {
	query := `select k.id, k.key_type, k.proxy_id, k.data
				from key k
				where k.subscription_id = $1`

	var key entity.Key
	var keyType string
	err := r.conn.QueryRow(ctx, query, subscriptionId).Scan(&key.Id, &keyType, &key.ProxyId, &key.Data)
	if err != nil {
		return nil, err
	}

	key.KeyType = entity.KeyTypeFromString(keyType)
	key.SubscriptionId = subscriptionId

	return &key, nil
}
