package repository

import (
	"context"
	"dev/master/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userId int64) (*entity.User, error)
	FindUserById(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

func (r *postgresRepository) CreateUser(ctx context.Context, userId int64) (*entity.User, error) {
	query := `INSERT INTO "user"(id, had_trial) VALUES($1, $2) RETURNING id, had_trial`

	var user entity.User
	err := r.conn.QueryRow(ctx, query, userId, false).Scan(&user.Id, &user.HadTrial)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgresRepository) FindUserById(ctx context.Context, userId int64) (*entity.User, error) {
	query := `SELECT id, had_trial FROM "user" WHERE id = $1`
	var user entity.User
	err := r.conn.QueryRow(ctx, query, userId).Scan(&user.Id, &user.HadTrial)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *postgresRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `UPDATE "user" SET had_trial = $1 WHERE id = $2 RETURNING id, had_trial`
	var newUser entity.User
	err := r.conn.QueryRow(ctx, query, user.HadTrial, user.Id).Scan(&newUser.Id, &newUser.HadTrial)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
