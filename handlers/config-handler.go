package handlers

import (
	"net/http"

	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
)

func AddConfig(w http.ResponseWriter, r *http.Request) {

	config := service.CreateConfig(w, r)
	utils.RenderJSON(w, config)
}

func GetAllConfigs(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, service.GetAllConfigs())
}
