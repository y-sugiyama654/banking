package service

import "github.com/y-sugiyama654/banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerRepository(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
