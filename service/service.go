package service

import (
	"go-atm-simulation/repository"
	"go-atm-simulation/types"
)

type Service interface {
	GetAccountInfoService(input types.GetAccountInfoServiceInput) (output types.GetAccountInfoServiceOutput, err error)
}

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAccountInfoService(input types.GetAccountInfoServiceInput) (output types.GetAccountInfoServiceOutput, err error) {
	getAccountInfoRepoInput := types.GetAccountInfoRepoInput(input)
	getAccountInfoRepoOutput, err := s.repository.GetAccountInfoRepo(getAccountInfoRepoInput)
	if err != nil {
		return output, err
	}
	output.AccountInformationDetail = getAccountInfoRepoOutput.AccountInformationDetail
	return output, err
}
