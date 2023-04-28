package service

import (
	"github.com/XenZi/ARS-2022-23/data"
	"github.com/XenZi/ARS-2022-23/model"
	"log"
)

/*
Kao sto ovde postoji ConfigService struktura, tako treba da postoji ConfigGroupService struktura koja ce biti namapirana na njen repository koji ce da koristi za handleovanja
*/

func CreateConfig() {
	log.Println("test")
}

func GetAllConfigs() []*model.Config {
	service := data.Data().Data
	var listOfConfigs []*model.Config
	for _, val := range service {
		for i := 0; i < len(val); i++ {
			listOfConfigs = append(listOfConfigs, val[i])
		}
	}
	return listOfConfigs
}
