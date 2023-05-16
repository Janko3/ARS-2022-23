package model

type Service struct {
	Data map[string][]*Config
}

// Config swagger: model Config
type Config struct {
	// Id of the config
	// in: string
	Id string `json: "id"`
	// Entries map of configs
	// in: map[string]string
	Entries map[string]string `json:"entries"`
	// Version of the config
	// in: string
	Version string `json:"version"`
}

// ConfigGroup swagger: model ConfigGroup
type ConfigGroup struct {
	// Id of the group
	// in: string
	Id string `json:"id"`
	// Array of configs with label
	// in: []*ConfigWithLabel
	Group []*ConfigWithLabel `json:"group"`
	// Version of the group
	// in: string
	Version string `json:"version"`
}

// ConfigWithLabel swagger: model ConfigWithLabel
type ConfigWithLabel struct {
	// Label
	// in: map[string]string
	Label map[string]string `json: "label"`
	// Config
	// in: Config
	Config Config `json: "config"`
}

type DbConfig struct {
	Service      Service
	ConfigGroups []*ConfigGroup
}

type ConfigGroupRequest struct {
	Keys []string `json: "keys"`
}

func (service *Service) AddConfig(config *Config) {
	service.Data["1"] = append(service.Data["1"], config)
}

func (group *DbConfig) AddGroupIntoDb(configGroup *ConfigGroup) {
	group.ConfigGroups = append(group.ConfigGroups, configGroup)
}

func (group *DbConfig) RemoveGroupFromDb(configGroupID, groupVersion string) {
	for i := 0; i < len(group.ConfigGroups); i++ {
		if group.ConfigGroups[i].Id == configGroupID && group.ConfigGroups[i].Version == groupVersion {
			group.ConfigGroups = append(group.ConfigGroups[:i], group.ConfigGroups[i+1:]...)
		}
	}
}

func (group *ConfigGroup) AddConfigIntoGroup(config *ConfigWithLabel) {
	group.Group = append(group.Group, config)
}
