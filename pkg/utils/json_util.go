package utils

import (
	"encoding/json"
	"io"

	"xcylla.io/common/log"
)

var logging log.Logger = log.NewLogger("utils")

func Decode(body io.Reader, v any) error {
	err := json.NewDecoder(body).Decode(&v)
	if err != nil {
		logging.Error("Error decoding data: %s", err.Error())
	}
	return err
}

func Encode(v any) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		logging.Error("Error enconding data: %s", err.Error())
		return nil, err
	}

	return data, nil
}
