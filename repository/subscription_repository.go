package repository

import (
	"context"
	"dev/master/entity"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *postgresRepository) FindSubscriptions(ctx context.Context, userId int64, countryId int64, active bool) ([]*entity.Subscription, error) {
	query := `SELECT id, expiration_date, user_id, country_id, is_trial FROM subscription WHERE user_id = $1`

	if countryId >= 0 {
		query += fmt.Sprintf(` AND country_id = %d`, countryId)
	}

	if active {
		query += ` AND expiration_date > now()`
	}

	rows, err := r.conn.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[entity.Subscription])

}

func (r *postgresRepository) CreateSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	query := `INSERT INTO subscription (user_id, country_id, expiration_date, is_trial) VALUES ($1, $2, $3, $4) RETURNING id, user_id, country_id, expiration_date, is_trial`

	err := r.conn.
		QueryRow(ctx, query, subscription.UserId, subscription.CountryId, subscription.ExpirationDateTime.Truncate(time.Second), subscription.IsTrial).
		Scan(&subscription.Id, &subscription.UserId, &subscription.CountryId, &subscription.ExpirationDateTime, &subscription.IsTrial)

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *postgresRepository) CreateTrialSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	user, err := r.FindUserById(ctx, subscription.UserId)
	if err != nil {
		return nil, err
	}

	if user.HadTrial {
		return nil, errors.New("user already had trial")
	}

	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE "user" SET had_trial=$1 WHERE id=$2`
	_, err = tx.Exec(ctx, query, true, user.Id)
	if err != nil {
		return nil, err
	}

	sub, err := r.CreateSubscription(ctx, subscription)

	if err != nil {
		return nil, err
	}

	return sub, tx.Commit(ctx)
}

func (r *postgresRepository) GetAllSubscriptions(ctx context.Context) ([]*entity.Subscription, error) {
	query := `SELECT id, expiration_date, user_id, country_id, is_trial FROM subscription`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[entity.Subscription])
}
