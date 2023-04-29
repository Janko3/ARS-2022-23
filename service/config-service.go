package service

import (
	"log"
	"net/http"

	"github.com/XenZi/ARS-2022-23/data"
	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/utils"
)

/*
Kao sto ovde postoji ConfigService struktura, tako treba da postoji ConfigGroupService struktura koja ce biti namapirana na njen repository koji ce da koristi za handleovanja
*/

func CreateConfig(w http.ResponseWriter, r *http.Request) *model.Config {
	if utils.IsContentTypeJSON(w, r) == false {
		return nil
	}
	config, err := utils.DecodeBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	service := data.Data()
	service.AddConfig(config)
	log.Println(service.Data)
	return config
}

func GetAllConfigs() []*model.Config {
	service := data.Data()
	var listOfConfigs []*model.Config
	for _, val := range service.Data {
		for i := 0; i < len(val); i++ {
			listOfConfigs = append(listOfConfigs, val[i])
		}
	}
	return listOfConfigs
}
