package entity

type Country struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
