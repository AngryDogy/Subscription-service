package entity

type Key struct {
	Id   int64
	Data []byte
	KeyType
	SubscriptionId int64
}

type KeyType int

const (
	Text KeyType = iota
	File
	Photo
)
