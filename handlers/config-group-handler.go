package handlers

import (
	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func AddConfigGroup(w http.ResponseWriter, r *http.Request) {
	keys, _ := utils.DecodeBodyForKeys(r.Body)
	utils.RenderJSON(w, service.CreateConfigGroup(keys))
}

func GetAllGroupConfigs(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, service.GetAllGroupConfigs())
}

func GetOneConfigGroup(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	utils.RenderJSON(w, service.GetGroupById(groupID))
}

func RemoveConfigGroup(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	utils.RenderJSON(w, service.RemoveConfigGroup(groupID))
}

func AddConfigIntoGroup(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	keys, _ := utils.DecodeBodyForKeys(r.Body)
	utils.RenderJSON(w, service.AddConfigIntoGroup(groupID, keys))
}
