package handlers

import (
	"log"
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
	log.Println(configId)
	utils.RenderJSON(w, service.GetConfigById(configId))
}
