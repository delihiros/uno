package entities

import "encoding/json"

type Response struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}
