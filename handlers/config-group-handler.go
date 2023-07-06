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
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
)

type ConfigGroupHandler struct {
	Repo   *repository.Repository
	Tracer opentracing.Tracer
	Closer io.Closer
}

// swagger:route POST /api/group-config Groups AddConfigGroup
// Create new configuration group
//
// responses:
//
//	400: BadRequest
//		200: ConfigGroup
func (configGroupHandler *ConfigGroupHandler) AddConfigGroup(w http.ResponseWriter, r *http.Request) {
	_, err := utils.DoesKeyExistInTheCurrentSessionOfRequests(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	span := tracing.StartSpanFromRequest("addConfigGroup", configGroupHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("add one config Group at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	contentType := r.Header.Get("Content-Type")
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

	group, err := utils.DecodeBodyForGroup(r.Body, cont)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdConfig, err := configGroupHandler.Repo.CreateNewGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, createdConfig, cont)
}

// swagger:route GET /api/group-config Groups GetAllGroupConfigs
// Get all configurations groups
//
// responses:
//
//	400: BadRequest
//	200: []ConfigGroup
func (configGroupHandler *ConfigGroupHandler) GetAllGroupConfigs(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest("getAllGroupConfigs", configGroupHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("get all config Groups at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	groups, err := configGroupHandler.Repo.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, groups, cont)
}

// swagger:route GET /api/group-config/{id}/{version}/ Groups GetOneConfigGroup
// Get one configuration group
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func (configGroupHandler *ConfigGroupHandler) GetOneConfigGroup(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest("getOneConfigGroup", configGroupHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("get one config Group at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	groupID := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	group, err := configGroupHandler.Repo.GetGroupByID(groupID, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	utils.RenderJSON(w, group, cont)
}

// swagger:route DELETE /api/group-config/{id}/{version} Groups RemoveConfigGroup
// Remove one configuration group
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func RemoveConfigGroup(w http.ResponseWriter, r *http.Request) {
	//groupID := mux.Vars(r)["id"]
	//version := mux.Vars(r)["version"]
	//utils.RenderJSON(w, service.RemoveConfigGroup(groupID, version))
}

// swagger:route GET /api/group-config/{id}/{version}/{label} Groups GetAllConfigsInGroupByLabel
// Get one configuration group by label
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func (configGroupHandler *ConfigGroupHandler) GetAllConfigsInGroupByLabel(w http.ResponseWriter, r *http.Request) {
	span := tracing.StartSpanFromRequest("getAllConfigsInGroupByLabel", configGroupHandler.Tracer, r)
	defer span.Finish()

	span.LogFields(
		tracing.LogString("handler", fmt.Sprintf("get all config Groups by label at: %s\n", r.URL.Path)),
	)
	cont := tracing.ContextWithSpan(context.Background(), span)
	id := mux.Vars(r)["id"]
	label := mux.Vars(r)["label"]
	version := mux.Vars(r)["version"]

	group, err := configGroupHandler.Repo.GetGroupConfigsByMatchingLabel(id, version, label, cont)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	utils.RenderJSON(w, group, cont)
}
