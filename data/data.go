package data

import (
	"github.com/XenZi/ARS-2022-23/model"
)

var DataInstance *model.DbConfig

func NewDataInstance() *model.DbConfig {
	if DataInstance == nil {
		DataInstance = &model.DbConfig{}

	}
	return DataInstance
}
