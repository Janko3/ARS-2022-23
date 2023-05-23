package repository

import (
	"encoding/json"
	"fmt"
	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/hashicorp/consul/api"
	"os"
)

type Repository struct {
	cli *api.Client
}

func New() (*Repository, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Repository{
		cli: client,
	}, nil
}

func (repo *Repository) CreateConfig(config *model.Config) (*model.Config, error) {
	kv := repo.cli.KV()

	dbID, entityID := utils.GenerateConfigKey(config.Version)
	config.Id = entityID

	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	d := &api.KVPair{Key: dbID, Value: data}
	_, err = kv.Put(d, nil)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (repo *Repository) GetConfigById(id string, version string) (*model.Config, error) {
	kv := repo.cli.KV()

	pair, _, err := kv.Get(utils.ConstructConfigKey(id, version), nil)
	if err != nil {
		return nil, err
	}
	config := &model.Config{}
	err = json.Unmarshal(pair.Value, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (repo *Repository) DeleteConfig(id string, version string) (*model.Config, error) {
	kv := repo.cli.KV()
	config, _ := repo.GetConfigById(id, version)
	_, err := kv.Delete(utils.ConstructConfigKey(id, version), nil)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (repo *Repository) GetAll() ([]*model.Config, error) {
	kv := repo.cli.KV()
	data, _, err := kv.List("config/", nil)
	if err != nil {
		return nil, err
	}

	configs := []*model.Config{}
	for _, pair := range data {
		config := &model.Config{}
		err = json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}
