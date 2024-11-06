package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Julio-Cesar07/gobid/internal/api/dtos"
)

type Response struct {
	Error any `json:"error,omitempty"`
	Data  any `json:"data,omitempty"`
}

func EncodeJson(w http.ResponseWriter, resp Response, status int) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

func DecodeValidJson[T dtos.Validator](r *http.Request) (T, dtos.Evaluator, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("decode json %w", err)
	}
	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}

func DecodeJson[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return data, fmt.Errorf("decode json %w", err)
	}

	return data, nil
}
