package handlers

import (
	"net/http"

	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
)

func AddConfig(w http.ResponseWriter, r *http.Request) {

	config := service.CreateConfig(w, r)
	utils.RenderJSON(w, config)
}

func GetAllConfigs(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, service.GetAllConfigs())
}

func GetOneConfig(w http.ResponseWriter, r *http.Request) {
	configId := mux.Vars(r)["id"]
	utils.RenderJSON(w, service.GetConfigById(configId))
}

func DeleteOneConfig(w http.ResponseWriter, r *http.Request) {
	configId := mux.Vars(r)["id"]
	firstValue, _ := service.DeleteConfigById(configId)
	if firstValue == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		utils.RenderJSON(w, firstValue)
	}

}
