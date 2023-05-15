package model

type Service struct {
	Data map[string][]*Config
}

type Config struct {
	Id      string            `json: "id"`
	Entries map[string]string `json:"entries"`
	Version string            `json:"version"`
}

type ConfigGroup struct {
	Id      string             `json: "id"`
	Group   []*ConfigWithLabel `json: "group"`
	Version string             `json:"version"`
}

type ConfigWithLabel struct {
	Label  map[string]string `json: "label"`
	Config Config            `json: "config"`
}
type BadRequest struct {
	Message    string
	StatusCode int
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
