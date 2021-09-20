package entities

import (
	"encoding/json"
	"strconv"
)

type AbilityCasts struct {
	CCast int `json:"c_cast"`
	QCast int `json:"q_cast"`
	ECast int `json:"e_cast"`
	XCast int `json:"x_cast"`
}

func (ac *AbilityCasts) UnmarshalJSON(data []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	ac.CCast, err = safeStringInt(m["c_cast"])
	if err != nil {
		return err
	}
	ac.QCast, err = safeStringInt(m["q_cast"])
	if err != nil {
		return err
	}
	ac.ECast, err = safeStringInt(m["e_cast"])
	if err != nil {
		return err
	}
	ac.XCast, err = safeStringInt(m["x_cast"])
	if err != nil {
		return err
	}
	return nil
}

func safeStringInt(m *json.RawMessage) (int, error) {
	if m == nil {
		return 0, nil
	}
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
