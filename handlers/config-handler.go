package handlers

import (
	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
	"log"
	"net/http"
)

func AddConfig(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	service.CreateConfig()
}

func GetAllConfigs(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, service.GetAllConfigs())
}
