package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(string(data), err)
	}
	return nil
}

func unmarshalUseNumber(decoder *json.Decoder, v interface{}) error {
	decoder.UseNumber()
	return decoder.Decode(v)
}

func formatError(v string, err error) error {
	return fmt.Errorf("string:`%s`,error:`%s`", v, err.Error())
}
