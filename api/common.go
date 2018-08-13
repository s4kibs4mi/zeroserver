package api

import (
	"net/http"
	"encoding/json"
	"errors"
	"io"
)

type (
	// object represents an object
	object map[string]interface{}

	// response represents a response
	response struct {
		code int
		Data interface{} `json:"data,omitempty"`
	}
)

func (r *response) ServeJSON(w http.ResponseWriter) error {
	if r.code == 0 {
		return errors.New("response status not defined")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		return err
	}
	return nil
}

func parseBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func forward(w http.ResponseWriter, resp *http.Response) {
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}
