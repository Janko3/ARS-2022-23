package model

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
	// Id of the config
	// in: string
	Id string `json: "id"`
	// Entries map of configs
	// in: map[string]string
	Entries map[string]string `json:"entries"`
}

type configServer struct {
	Keys map[string]string
}
