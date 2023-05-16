package handlers

import (
	"net/http"

	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
)

// swagger:route POST /api/config Configuration AddConfig
// Create new configuration
//
// responses:
//
//	400: BadRequest
//		200: Config

func AddConfig(w http.ResponseWriter, r *http.Request) {

	config := service.CreateConfig(w, r)
	utils.RenderJSON(w, config)
}

// swagger:route GET /api/config Configuration GetAllConfigs
// Get all configurations
//
// responses:
//
//	400: BadRequest
//	200: []Config
func GetAllConfigs(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, service.GetAllConfigs())
}

// swagger:route GET /api/config/{id}/{version} Configuration GetOneConfig
// Get one configuration
//
// responses:
//
//	 400: BadRequest
//		200: Config
func GetOneConfig(w http.ResponseWriter, r *http.Request) {
	configId := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	utils.RenderJSON(w, service.GetConfigById(configId, version))
}

// swagger:route DELETE /api/config/{id}/{version} Configuration DeleteOneConfig
// Delete one configuration
//
// responses:
//
// 400: BadRequest
// 200: Config
func DeleteOneConfig(w http.ResponseWriter, r *http.Request) {
	configId := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	firstValue, _ := service.DeleteConfigById(configId, version)
	if firstValue == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		utils.RenderJSON(w, firstValue)
	}

}
