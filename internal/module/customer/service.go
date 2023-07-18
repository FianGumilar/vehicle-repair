package customer

import "github.com/FianGumilar/vehicle-repair/domain"

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{customerRepository: customerRepository}
}
