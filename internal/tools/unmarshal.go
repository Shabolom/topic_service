package tools

import (
	"encoding/json"
	"io"
)

func ShortUnmarshal(body io.Reader, obj interface{}) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	return nil
}
