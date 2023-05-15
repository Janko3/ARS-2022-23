package handlers

import (
	"fmt"
	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func AddConfigGroup(w http.ResponseWriter, r *http.Request) {
	group, _ := utils.DecodeBodyForGroup(r.Body)
	utils.RenderJSON(w, service.CreateConfigGroup(group))
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

func GetAllConfigsInGroupByLabel(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	label := mux.Vars(r)["label"]
	fmt.Println(label)
	fmt.Println(groupID)
	utils.RenderJSON(w, service.GetGroupByIdAndLabel(groupID, label))
}
