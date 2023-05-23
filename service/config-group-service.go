package service

import (
	"fmt"

	"github.com/XenZi/ARS-2022-23/data"
	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/utils"
)

func CreateConfigGroup(group *model.ConfigGroup) *model.ConfigGroup {
	data := data.NewDataInstance()
	group.Id = utils.CreateId()
	for i := 0; i < len(group.Group); i++ {
		group.Group[i].Id = utils.CreateId()
	}
	data.AddGroupIntoDb(group)
	return group
}

func GetAllGroupConfigs() []*model.ConfigGroup {
	data := data.NewDataInstance().ConfigGroups
	return data
}

func GetGroupById(groupID, groupVersion string) *model.ConfigGroup {
	configGroups := data.DataInstance.ConfigGroups
	for i := 0; i < len(configGroups); i++ {
		if configGroups[i].Id == groupID && configGroups[i].Version == groupVersion {
			return configGroups[i]
		}
	}
	return nil
}

func GetGroupByIdAndLabel(groupId string, labelMatching string, groupVersion string) []*model.ConfigWithLabel {
	configGroups := data.DataInstance.ConfigGroups
	var configsWithLabels []*model.ConfigWithLabel
	for i := 0; i < len(configGroups); i++ {
		if configGroups[i].Id == groupId && configGroups[i].Version == groupVersion {
			for j := 0; j < len(configGroups[i].Group); j++ {
				label := ""
				for k, v := range configGroups[i].Group[j].Label {
					label += k + ":" + v + ";"
				}
				fmt.Println("....")
				fmt.Println(label)
				if label == (labelMatching + ";") {
					configsWithLabels = append(configsWithLabels, configGroups[i].Group[j])
				}
			}
		}
	}
	return configsWithLabels
}

func RemoveConfigGroup(groupID, version string) string {
	configGroups := data.DataInstance
	configGroups.RemoveGroupFromDb(groupID, version)
	return groupID
}
