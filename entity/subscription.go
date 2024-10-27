package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type Subscription struct {
	Id             int64
	UserId         int64
	CountryId      int64
	ExpirationTime timestamp.Timestamp
}
