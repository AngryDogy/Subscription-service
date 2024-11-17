package entity

type Key struct {
	Id             int64   `db:"id"`
	Data           []byte  `db:"data"`
	KeyType        KeyType `db:"-"`
	SubscriptionId int64   `db:"subscription_id"`
	ProxyId        int64   `db:"proxy_id"`
	IdInProxy      int64   `db:"id_in_proxy"`
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

func (k KeyType) String() string {
	return [...]string{"TEXT", "FILE", "PHOTO"}[k]
}
