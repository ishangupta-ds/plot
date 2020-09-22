package pkg

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type logRt struct {
	transport http.RoundTripper
}

// queryResult contains result data for a query.
type queryResult struct {
	Type   model.ValueType       `json:"resultType"`
	Result []*model.SampleStream `json:"result"`

	// The decoded value.
	v model.Value
}

type apiResponse struct {
	Status    string          `json:"status"`
	Data      json.RawMessage `json:"data"`
	ErrorType v1.ErrorType    `json:"errorType"`
	Error     string          `json:"error"`
	Warnings  []string        `json:"warnings,omitempty"`
}

type validator struct {
	client    v1.API
	startTime time.Time
	values    map[string][]*model.SampleStream
	out       *os.File
}
