package entity

type User struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	HadTrial bool   `db:"had_trial"`
}
