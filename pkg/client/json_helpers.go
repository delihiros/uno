package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func safeStringInt(m *json.RawMessage) (int, error) {
	var i int
	err := json.Unmarshal(*m, &i)
	if err != nil {
		var s string
		err = json.Unmarshal(*m, &s)
		if err != nil {
			return 0, err
		}
		if s == "N.A" {
			return 0, nil
		}
		i, err = strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
	}
	return i, nil
}

func PrintJSON(v interface{}, prettify bool) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if prettify {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, b, "", "  ")
		if err != nil {
			return err
		}
		b = prettyJSON.Bytes()
	}
	fmt.Println(string(b))
	return nil
}
