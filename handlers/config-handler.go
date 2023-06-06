package handlers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/XenZi/ARS-2022-23/repository"
	"github.com/XenZi/ARS-2022-23/tracing"
	"github.com/opentracing/opentracing-go"

	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
)

type ConfigHandler struct {
	Repo   *repository.Repository
	Tracer opentracing.Tracer
	Closer io.Closer
}

// swagger:route POST /api/config Configuration AddConfig
// Create new configuration
//
// responses:
//
//	400: BadRequest
//		200: Config

func (configHandler *ConfigHandler) AddConfig(w http.ResponseWriter, req *http.Request) {
	span := tracing.StartSpanFromRequest("addConfig", configHandler.Tracer, req)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("creating config at: %s\n", req.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	config, err := utils.DecodeBody(req.Body, cont)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdConfig, err := configHandler.Repo.CreateConfig(config)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, createdConfig, cont)

}

// swagger:route GET /api/config/{id}/{version} Configuration GetOneConfig
// Get one configuration
//
// responses:
//
//	 400: BadRequest
//		200: Config
func (configHandler *ConfigHandler) GetOneConfig(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest("getOneConfig", configHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("get one config at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	configId := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	config, err := configHandler.Repo.GetConfigById(configId, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, config, cont)
}

// swagger:route DELETE /api/config/{id}/{version} Configuration DeleteOneConfig
// Delete one configuration
//
// responses:
//
// 400: BadRequest
// 200: Config
func (configHandler *ConfigHandler) DeleteOneConfig(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest("deleteOneConfig", configHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("delete config at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	configId := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	config, err := configHandler.Repo.DeleteConfig(configId, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, config, cont)
}

// swagger:route GET /api/config Configuration GetAllConfigs
// Get all configurations
//
// responses:
//
//	400: BadRequest
//	200: []Config

func (configHandler *ConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest("getAllConfigs", configHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("get all configs at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	configs, err := configHandler.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, configs, cont)
}
