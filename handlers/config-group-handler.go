package handlers

import (
	"fmt"
	"net/http"

	"github.com/XenZi/ARS-2022-23/service"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/gorilla/mux"
)

// swagger:route POST /api/group-config Groups AddConfigGroup
// Create new configuration group
//
// responses:
//
//	400: BadRequest
//		200: ConfigGroup
func AddConfigGroup(w http.ResponseWriter, r *http.Request) {
	group, _ := utils.DecodeBodyForGroup(r.Body)
	utils.RenderJSON(w, service.CreateConfigGroup(group))
}

// swagger:route GET /api/group-config Groups GetAllGroupConfigs
// Get all configurations groups
//
// responses:
//
//	400: BadRequest
//	200: []ConfigGroup
func GetAllGroupConfigs(w http.ResponseWriter, r *http.Request) {
	utils.RenderJSON(w, service.GetAllGroupConfigs())
}

// swagger:route GET /api/group-config/{id}/{version} Groups GetOneConfigGroup
// Get one configuration group
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func GetOneConfigGroup(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	utils.RenderJSON(w, service.GetGroupById(groupID, version))
}

// swagger:route DELETE /api/group-config/{id}/{version} Groups RemoveConfigGroup
// Remove one configuration group
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func RemoveConfigGroup(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	version := mux.Vars(r)["version"]
	utils.RenderJSON(w, service.RemoveConfigGroup(groupID, version))
}

// swagger:route GET /api/group-config/{id}/{version}/{label} Groups GetAllConfigsInGroupByLabel
// Get one configuration group by label
//
// responses:
//
//	 400: BadRequest
//		200: ConfigGroup
func GetAllConfigsInGroupByLabel(w http.ResponseWriter, r *http.Request) {
	groupID := mux.Vars(r)["id"]
	label := mux.Vars(r)["label"]
	version := mux.Vars(r)["version"]
	fmt.Println(label)
	fmt.Println(groupID)
	utils.RenderJSON(w, service.GetGroupByIdAndLabel(groupID, label, version))
}
