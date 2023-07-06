package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/XenZi/ARS-2022-23/model"
	"github.com/XenZi/ARS-2022-23/tracing"
	"github.com/XenZi/ARS-2022-23/utils"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
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

func (repo *Repository) CreateNewGroup(group *model.ConfigGroup) (*model.ConfigGroup, error) {
	kv := repo.cli.KV()

	group.Id = uuid.New().String()
	for i := 0; i < len(group.Group); i++ {
		group.Group[i].Id = uuid.New().String()
		generatedLabels := utils.GetLabelAsStringWithSeparator(group.Group[i].Label)
		data, err := json.Marshal(group.Group[i])
		if err != nil {
			return nil, err
		}
		placedValue := &api.KVPair{Key: utils.ConstructGroupKey(group.Id, group.Version, generatedLabels, group.Group[i].Id), Value: data}
		_, err = kv.Put(placedValue, nil)
		if err != nil {
			return nil, err
		}
	}
	return group, nil
}

func (repo *Repository) GetGroupByID(id string, version string) (*model.ConfigGroup, error) {
	kv := repo.cli.KV()
	data, _, err := kv.List("group/"+id+"/"+version, nil)
	if err != nil {
		return nil, err
	}
	var groupList []*model.ConfigWithLabel
	for _, key := range data {
		config := &model.ConfigWithLabel{}
		err := json.Unmarshal(key.Value, config)
		if err != nil {
			return nil, err
		}
		groupList = append(groupList, config)
	}
	return &model.ConfigGroup{Group: groupList, Id: id, Version: version}, nil
}

func (repo *Repository) GetAllGroups() ([]*model.ConfigGroup, error) {
	kv := repo.cli.KV()

	data, _, err := kv.List("group/", nil)

	if err != nil {
		return nil, err
	}

	groupMap := make(map[string]*model.ConfigGroup)
	groupList := []*model.ConfigGroup{}

	for _, key := range data {
		config := &model.ConfigWithLabel{}
		err := json.Unmarshal(key.Value, config)
		if err != nil {
			return nil, err
		}

		groupVersion := utils.GetKeyIndexInfo("groupVersion", key.Key)
		groupID := utils.GetKeyIndexInfo("groupID", key.Key)
		if groupMap[groupID+groupVersion] == nil {
			newGroup := &model.ConfigGroup{}
			newGroup.Group = []*model.ConfigWithLabel{}
			newGroup.Version = groupVersion
			newGroup.Id = groupID
			newGroup.Group = append(newGroup.Group, config)
			groupMap[groupID+groupVersion] = newGroup
		} else {
			groupMap[groupID+groupVersion].Group = append(groupMap[groupID+groupVersion].Group, config)
		}
	}

	for _, g := range groupMap {
		groupList = append(groupList, g)
	}
	return groupList, nil
}

func (repo *Repository) GetGroupConfigsByMatchingLabel(id, version, label string, cont context.Context) (*model.ConfigGroup, error) {
	span := tracing.StartSpanFromContext(cont, "GetGroupConfigsByMatchingLabel")
	defer span.Finish()

	span.LogFields(
		tracing.LogString("repo", fmt.Sprintf("get all config Groups by label:\n")),
	)
	kv := repo.cli.KV()
	data, _, err := kv.List("group/"+id+"/"+version+"/"+label+"/", nil)
	if err != nil {
		return nil, err
	}
	var groupList []*model.ConfigWithLabel
	for _, key := range data {
		config := &model.ConfigWithLabel{}
		err := json.Unmarshal(key.Value, config)
		if err != nil {
			return nil, err
		}
		groupList = append(groupList, config)
	}
	return &model.ConfigGroup{Group: groupList, Id: id, Version: version}, nil

}
