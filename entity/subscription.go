package entity

import (
	"time"
)

type Subscription struct {
	Id                 int64     `db:"id"`
	UserId             int64     `db:"user_id"`
	CountryId          int64     `db:"country_id"`
	ExpirationDateTime time.Time `db:"expiration_date"`
}
