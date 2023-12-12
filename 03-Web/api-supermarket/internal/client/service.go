package client

type ClientService struct {
	repo ClientRepository
}

func NewClientService() ClientService {
	repo := NewClientRepository()

	return ClientService{repo: repo}
}
