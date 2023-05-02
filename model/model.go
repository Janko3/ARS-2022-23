package model

type Config struct {
	Id      string            `json: "id"`
	Entries map[string]string `json:"entries"`
}

type Service struct {
	Data map[string][]*Config
}

type ConfigGroup struct {
	Id    string    `json: "id"`
	Group []*Config `json: "group"`
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

func (group *DbConfig) RemoveGroupFromDb(configGroupID string) {
	for i := 0; i < len(group.ConfigGroups); i++ {
		if group.ConfigGroups[i].Id == configGroupID {
			group.ConfigGroups = append(group.ConfigGroups[:i], group.ConfigGroups[i+1:]...)
		}
	}
}

func (group *ConfigGroup) AddConfigIntoGroup(config *Config) {
	group.Group = append(group.Group, config)
}
