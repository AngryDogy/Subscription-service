package entity

type User struct {
	Id       int64 `db:"id"`
	HadTrial bool  `db:"had_trial"`
}
