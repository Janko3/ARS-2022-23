package utils

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	// config/config_id/config_version
	config = "config/%s/%s"
	// group/group_id/group_version/labels/config_id
	group = "group/%s/%s/%s/config/%s"
)

func GenerateConfigKey(version string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(config, id, version), id
}

func ConstructConfigKey(id string, version string) string {
	return fmt.Sprintf(config, id, version)
}

func ConstructGroupKey(idGroup string, version string, labels string, idConfig string) string {
	return fmt.Sprintf(group, idGroup, version, labels, idConfig)
}
