package internal

import (
	"io"
	"net/http"
)

// Custom Error instance to be used for reporting business constraint errors
type ApiErr struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func (a *ApiErr) Error() string {
	return a.Message
}

func ParseAndValidate(r *http.Request, v validatable) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	DecodeJSON(b, v)
	if err = Validate(v); err != nil {
		panic(err)
	}
}

func HttpRespond(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := EncodeJSON(v)

	w.WriteHeader(statusCode)
	w.Write(json)
}
