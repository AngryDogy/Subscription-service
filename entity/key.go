package entity

type Key struct {
	Id             int64   `db:"id"`
	Data           []byte  `db:"data"`
	KeyType        KeyType `db:"-"`
	SubscriptionId int64   `db:"subscription_id"`
	ProxyId        int64   `db:"proxy_id"`
}

type KeyType int

const (
	Text KeyType = iota
	File
	Photo
)

func KeyTypeFromString(keyType string) KeyType {
	switch keyType {
	case "TEXT":
		return Text
	case "FILE":
		return File
	case "PHOTO":
		return Photo
	default:
		return Text
	}
}
