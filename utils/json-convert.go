package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/tracing"
)

func DecodeBody(r io.Reader, cont context.Context) (*model.Config, error) {
	span := tracing.StartSpanFromContext(cont, "decodeBody")
	defer span.Finish()
	span.LogFields(
		tracing.LogString("requestUtility", fmt.Sprintf("decoding body")),
	)
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var config model.Config
	if err := dec.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func DecodeBodyForGroup(r io.Reader, cont context.Context) (*model.ConfigGroup, error) {
	span := tracing.StartSpanFromContext(cont, "decodeBodyForGroup")
	defer span.Finish()
	span.LogFields(
		tracing.LogString("requestUtility", fmt.Sprintf("decoding group")),
	)
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var configWithLabel model.ConfigGroup
	if err := dec.Decode(&configWithLabel); err != nil {
		return nil, err
	}
	return &configWithLabel, nil
}

func RenderJSON(w http.ResponseWriter, v interface{}, cont context.Context) {
	span := tracing.StartSpanFromContext(cont, "renderJSON")
	defer span.Finish()
	span.LogFields(
		tracing.LogString("requestUtility", fmt.Sprintf("rendering JSON")),
	)
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
