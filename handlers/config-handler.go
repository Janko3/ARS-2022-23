package handlers

import (
	"github.com/XenZi/ARS-2022-23/service"
	"net/http"
)

type ConfigHandler struct {
	Service *service.ConfigService
}

func AddConfig(w http.ResponseWriter, r *http.Request) {

}
