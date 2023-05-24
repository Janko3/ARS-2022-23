package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/XenZi/ARS-2022-23/model"
)

func DecodeBody(r io.Reader) (*model.Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var config model.Config
	if err := dec.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func DecodeBodyForGroup(r io.Reader) (*model.ConfigGroup, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var configWithLabel model.ConfigGroup
	if err := dec.Decode(&configWithLabel); err != nil {
		return nil, err
	}
	return &configWithLabel, nil
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
