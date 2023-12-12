package client

import (
	"api-supermarket/internal/domain"
	"api-supermarket/pkg"
)

type ClientRepository struct {
	ClientsDB []domain.Client
}

func NewClientRepository() ClientRepository {

	slice := pkg.FullfilClientsDB("../../clients.json")

	return ClientRepository{slice}
}
