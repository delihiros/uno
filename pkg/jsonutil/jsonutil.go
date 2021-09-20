package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func FormatJSON(v interface{}, prettify bool) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	if prettify {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, b, "", "  ")
		if err != nil {
			return "", err
		}
		b = prettyJSON.Bytes()
	}
	return string(b), nil
}

func PrintJSON(v interface{}, prettify bool) error {
	s, err := FormatJSON(v, prettify)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
