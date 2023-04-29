package model

type Config struct {
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
