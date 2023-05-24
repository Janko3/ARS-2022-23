package handlers

import (
	"errors"
	"github.com/XenZi/ARS-2022-23/repository"
	"mime"
	"net/http"

	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
)

type ConfigHandler struct {
	Repo *repository.Repository
}

// swagger:route POST /api/config Configuration AddConfig
// Create new configuration
//
// responses:
//
//	400: BadRequest
//		200: Config

func (configHandler *ConfigHandler) AddConfig(w http.ResponseWriter, req *http.Request) {
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

	config, err := utils.DecodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdConfig, err := configHandler.Repo.CreateConfig(config)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, createdConfig)
}

// swagger:route GET /api/config/{id}/{version} Configuration GetOneConfig
// Get one configuration
//
// responses:
//
//	 400: BadRequest
//		200: Config
func (configHandler *ConfigHandler) GetOneConfig(w http.ResponseWriter, r *http.Request) {
	configId := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	config, err := configHandler.Repo.GetConfigById(configId, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, config)
}

// swagger:route DELETE /api/config/{id}/{version} Configuration DeleteOneConfig
// Delete one configuration
//
// responses:
//
// 400: BadRequest
// 200: Config
func (configHandler *ConfigHandler) DeleteOneConfig(w http.ResponseWriter, r *http.Request) {
	configId := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	config, err := configHandler.Repo.DeleteConfig(configId, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, config)
}

// swagger:route GET /api/config Configuration GetAllConfigs
// Get all configurations
//
// responses:
//
//	400: BadRequest
//	200: []Config

func (configHandler *ConfigHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	configs, err := configHandler.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RenderJSON(w, configs)
}
