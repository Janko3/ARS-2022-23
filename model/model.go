package model

type Config struct {
	Entries map[string]string `json:"entries"`
}

type Service struct {
	Data map[string][]*Config `json:"data"`
}
