package model

import "encoding/json"

type Object struct {
	Type string          `json:"type"`
	List json.RawMessage `json:"list"`
}
