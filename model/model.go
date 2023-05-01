package model

type Config struct {
	Id      string            `json: "id"`
	Entries map[string]string `json:"entries"`
}

type Service struct {
	Data map[string][]*Config `json:"data"`
}

func (service *Service) AddConfig(config *Config) {
	service.Data["1"] = append(service.Data["1"], config)
}

type BadRequest struct {
	Message    string
	StatusCode int
}

type DbConfig struct {
	Service Service
}
