package utils

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"

	"github.com/XenZi/ARS-2022-23/model"
	"github.com/google/uuid"
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

func RenderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func doesContentTypeExists(req *http.Request) *model.BadRequest {
	contentType := req.Header.Get("Content-Type")
	_, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		badRequest := model.BadRequest{Message: "Request is bad", StatusCode: http.StatusBadRequest}
		return &badRequest
	}
	return nil
}

func IsContentTypeJSON(w http.ResponseWriter, req *http.Request) bool {
	isRequestValid := doesContentTypeExists(req)
	if isRequestValid != nil {
		http.Error(w, isRequestValid.Message, isRequestValid.StatusCode)
		return false
	}
	return true
}
func CreateId() string {
	return uuid.New().String()
}

func Remove(slice []*model.Config, s int) []*model.Config {
	return append(slice[:s], slice[s+1:]...)
}
