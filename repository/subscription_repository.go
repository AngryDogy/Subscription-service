package repository

import (
	"context"
	"dev/master/entity"
	"errors"
)

func (r *postgresRepository) FindSubscriptionsByUserId(userId int64) ([]*entity.Subscription, error) {
	query := `SELECT * FROM subscription WHERE user_id = $1`

	rows, err := r.conn.Query(query, userId)
	if err != nil {
		return nil, err
	}

	subs := make([]*entity.Subscription, 0)
	for rows.Next() {
		var sub entity.Subscription
		err := rows.Scan(&sub.Id, &sub.UserId, &sub.CountryId, &sub.ExpirationTime)
		if err != nil {
			continue
		}

		subs = append(subs, &sub)
	}
	return subs, nil

}

func (r *postgresRepository) CreateSubscription(ctx context.Context, userId, countryId int64) (*entity.Subscription, error) {
	query := `INSERT INTO subscription (user_id, country_id, expiration_date) VALUES ($1, $2, $3) RETURNING id, user_id, country_id, expiration_date`

	var sub entity.Subscription
	err := r.conn.QueryRowContext(ctx, query, userId, countryId).Scan(&sub.Id, &sub.UserId, &sub.CountryId, &sub.ExpirationTime)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *postgresRepository) CreateTrialSubscription(ctx context.Context, userId, countryId int64) (*entity.Subscription, error) {
	user, err := r.FindUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user.HadTrial {
		return nil, errors.New("user already had trial")
	}

	tx, err := r.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `UPDATE "user" SET had_trial=$1 WHERE id=$2`
	_, err = tx.ExecContext(ctx, query, user.HadTrial, user.Id)
	if err != nil {
		return nil, err
	}

	query = `INSERT INTO subscription (user_id, country_id, expiration_date)
				VALUES($1, $2, $3) RETURNING id, user_id, country_id, expiration_date`
	var sub entity.Subscription
	err = tx.QueryRowContext(ctx, query, user.HadTrial, user.Id, countryId).Scan(
		&sub.Id,
		&sub.UserId,
		&sub.CountryId,
		&sub.ExpirationTime)

	if err != nil {
		return nil, err
	}
	return &sub, tx.Commit()
}
