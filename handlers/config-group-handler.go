package handlers

import (
	"errors"
	"github.com/XenZi/ARS-2022-23/repository"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type ConfigGroupHandler struct {
	Repo *repository.Repository
}

// swagger:route POST /api/group-config Groups AddConfigGroup
// Create new configuration group
//
// responses:
//
//	400: BadRequest
//		200: ConfigGroup
func (configGroupHandler *ConfigGroupHandler) AddConfigGroup(w http.ResponseWriter, r *http.Request) {
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

	group, err := utils.DecodeBodyForGroup(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdConfig, err := configGroupHandler.Repo.CreateNewGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, createdConfig)
}

// swagger:route GET /api/group-config Groups GetAllGroupConfigs
// Get all configurations groups
//
// responses:
//
//	400: BadRequest
//	200: []ConfigGroup
func (configGroupHandler *ConfigGroupHandler) GetAllGroupConfigs(w http.ResponseWriter, r *http.Request) {
	groups, err := configGroupHandler.Repo.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, groups)
}

// swagger:route GET /api/group-config/{id}/{version} Groups GetOneConfigGroup
// Get one configuration group
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func (configGroupHandler *ConfigGroupHandler) GetOneConfigGroup(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	group, err := configGroupHandler.Repo.GetGroupByID(groupID, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	utils.RenderJSON(w, group)
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
	id := mux.Vars(r)["id"]
	label := mux.Vars(r)["label"]
	version := mux.Vars(r)["version"]
	group, err := configGroupHandler.Repo.GetGroupConfigsByMatchingLabel(id, version, label)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	utils.RenderJSON(w, group)
}
