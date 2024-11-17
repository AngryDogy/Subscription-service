package proxy

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

const j = `{ "id": "2", "key": "ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTptMElrR3FCOGREalZqRFFNZjdENVdY@158.160.6.79:1416/?outline=1", "key_type": "TEXT"}`

func Test_keyConversion(t *testing.T) {
	var key jsonKey
	err := json.Unmarshal([]byte(j), &key)
	require.Nil(t, err)
	require.Equal(t, jsonKey{
		Id:      "2",
		Data:    "ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTptMElrR3FCOGREalZqRFFNZjdENVdY@158.160.6.79:1416/?outline=1",
		KeyType: "TEXT",
	}, key)
}
