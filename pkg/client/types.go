package client

import "encoding/json"

type HenrikdevResponse struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type UnofficialValorantResponse struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data"`
}
