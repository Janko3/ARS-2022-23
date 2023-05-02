package service

import (
	"github.com/XenZi/ARS-2022-23/data"
	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/utils"
)

func CreateConfigGroup(request *model.ConfigGroupRequest) *model.ConfigGroup {
	configs := data.NewDataInstance().Service.Data
	newGroup := model.ConfigGroup{Id: utils.CreateId()}
	var newGroupConfigs []*model.Config

	for i := 0; i < len(request.Keys); i++ {
		for _, v := range configs {
			for j := 0; j < len(v); j++ {
				if request.Keys[i] == v[j].Id {
					newGroupConfigs = append(newGroupConfigs, v[j])
				}
			}
		}
	}
	newGroup.Group = newGroupConfigs
	data := data.NewDataInstance()
	data.AddGroupIntoDb(&newGroup)
	return &newGroup
}

func GetAllGroupConfigs() []*model.ConfigGroup {
	data := data.NewDataInstance().ConfigGroups
	return data
}

func GetGroupById(groupID string) *model.ConfigGroup {
	configGroups := data.DataInstance.ConfigGroups
	for i := 0; i < len(configGroups); i++ {
		if configGroups[i].Id == groupID {
			return configGroups[i]
		}
	}
	return nil
}

func AddConfigIntoGroup(groupID string, request *model.ConfigGroupRequest) *model.ConfigGroup {
	configs := data.NewDataInstance().Service.Data
	groups := data.NewDataInstance().ConfigGroups
	var group *model.ConfigGroup
	for i := 0; i < len(groups); i++ {
		if groups[i].Id == groupID {
			{
				group = groups[i]
			}
		}
	}
	for i := 0; i < len(request.Keys); i++ {
		for _, v := range configs {
			for j := 0; j < len(v); j++ {
				if request.Keys[i] == v[j].Id {
					group.AddConfigIntoGroup(v[j])
				}
			}
		}
	}
	return group
}

func RemoveConfigGroup(groupID string) string {
	configGroups := data.DataInstance
	configGroups.RemoveGroupFromDb(groupID)
	return groupID
}
