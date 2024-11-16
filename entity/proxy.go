package entity

type Proxy struct {
	Id        int64  `db:"id"`
	Address   string `db:"address"`
	CountryId int64  `db:"country_id"`
}
