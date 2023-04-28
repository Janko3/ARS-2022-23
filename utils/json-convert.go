package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/XenZi/ARS-2022-23/model"
	"github.com/google/uuid"
)

func DecodeBody(r io.Reader) (*model.Service, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var s model.Service
	if err := dec.Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func RenderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func CreateId() string {
	return uuid.New().String()
}
