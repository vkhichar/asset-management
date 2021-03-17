package contract

import "encoding/json"

type UpdateRequest struct {
	Status         *string         `json:"status"`
	Specifications json.RawMessage `json:"specifications"`
}
