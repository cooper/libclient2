package libclient2

import (
	"encoding/json"
)

func genericSend(client Client, command string, params map[string]interface{}) bool {
	b, err := json.Marshal([]interface{}{command, params})
	if err != nil {
		return false
	}
	b = append(b, '\n')
	if _, err = client.write(b); err != nil {
		return false
	}
	return true
}

func genericWrite(client Client, b []byte) (int, error) {
	return client.socket().Write(b)
}
