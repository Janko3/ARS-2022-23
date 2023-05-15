package data

import (
	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/utils"
)

var DataInstance *model.DbConfig

func NewDataInstance() *model.DbConfig {
	if DataInstance == nil {
		DataInstance = &model.DbConfig{}
		cf1 := model.Config{Id: utils.CreateId(), Entries: map[string]string{
			"test":  "test",
			"test2": "test2",
		}}
		cf2 := model.Config{Id: utils.CreateId(), Entries: map[string]string{
			"test3": "test3",
			"test4": "test4",
		}}
		DataInstance.Service = model.Service{}
		DataInstance.ConfigGroups = []*model.ConfigGroup{
			{
				Id:    utils.CreateId(),
				Group: []*model.Config{&cf1},
			},
		}
		DataInstance.Service.Data = map[string][]*model.Config{
			"1": {&cf1, &cf2},
		}
	}
	return DataInstance
}
