package service

import (
	"github.com/Rafael-Lopez/GO/RESTAPI/banking/domain"
	"github.com/Rafael-Lopez/GO/RESTAPI/banking/dto"
	"github.com/Rafael-Lopez/GO/RESTAPI/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)

	for _, c := range customers {
		response = append(response, c.ToDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
