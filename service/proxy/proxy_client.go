package proxy

import (
	"dev/master/entity"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
)

const (
	MASTER_KEY_HEADER = "X-Master-Key"
	MASTER_KEY        = "VLADOS"
)

type Client interface {
	CreateKey(address string) (*entity.Key, error)
}

func NewClient(logger *zap.Logger) Client {
	return &clientImpl{logger: logger}
}

type clientImpl struct {
	logger *zap.Logger
}

type jsonKey struct {
	Id      string `json:"id"`
	Data    string `json:"key"`
	KeyType string `json:"key_type"`
}

func (c *clientImpl) CreateKey(address string) (*entity.Key, error) {
	req, err := http.NewRequest("POST", "http://"+address+"/keys", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(MASTER_KEY_HEADER, MASTER_KEY)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var key jsonKey
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.logger.Info("Received key", zap.String("key", string(data)))
	err = json.Unmarshal(data, &key)

	return &entity.Key{
		IdInProxy: key.Id,
		Data:      []byte(key.Data),
		KeyType:   entity.KeyTypeFromString(key.KeyType),
	}, nil
}
