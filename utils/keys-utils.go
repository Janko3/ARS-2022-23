package utils

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	config = "config/%s/%s"
	group  = "group/%s/%s/%s/config/%s"
)

func GenerateConfigKey(version string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(config, id, version), id
}

func ConstructConfigKey(id string, version string) string {
	return fmt.Sprintf(config, id, version)
}

func GenerateGroupKey(version string, labels string) (string, string, string) {
	idGroup := uuid.New().String()
	idConfig := uuid.New().String()
	return fmt.Sprintf(group, idGroup, version, labels, idConfig), idGroup, idConfig
}

func ConstructGroupKey(idGroup string, version string, labels string, idConfig string) string {
	return fmt.Sprintf(group, idGroup, version, labels, idConfig)
}
