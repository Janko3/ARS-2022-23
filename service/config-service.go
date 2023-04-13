package service

import "github.com/XenZi/ARS-2022-23/repository"

/*
Kao sto ovde postoji ConfigService struktura, tako treba da postoji ConfigGroupService struktura koja ce biti namapirana na njen repository koji ce da koristi za handleovanja
*/
type ConfigService struct {
	Repository *repository.ConfigRepository
}
