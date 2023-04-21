package data

import "github.com/XenZi/ARS-2022-23/model"

func data() *model.Service {
	s1 := model.Service{}

	cf1 := model.Config{Entries: map[string]string{
		"test":  "test",
		"test2": "test2",
	}}
	cf2 := model.Config{Entries: map[string]string{
		"test3": "test3",
		"test4": "test4",
	}}
	cf1p := &cf1
	cf2p := &cf2
	cslice := []*model.Config{cf1p, cf2p}
	s1.Data = map[string][]*model.Config{
		"1": cslice,
	}
	return &s1
}
