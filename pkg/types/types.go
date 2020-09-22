package types

import (
	"encoding/json"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

// QueryResult contains result data for a query.
type QueryResult struct {
	Type   model.ValueType       `json:"resultType"`
	Result []*model.SampleStream `json:"result"`

	// The decoded value.
	v model.Value
}

// ApiResponse contains API response for a query.
type ApiResponse struct {
	Status    string          `json:"status"`
	Data      json.RawMessage `json:"data"`
	ErrorType v1.ErrorType    `json:"errorType"`
	Error     string          `json:"error"`
	Warnings  []string        `json:"warnings,omitempty"`
}
