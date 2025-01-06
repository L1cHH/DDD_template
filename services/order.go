package services

import (
	"impl_DDD/domain/customer"
	"impl_DDD/domain/customer/memory"
)

type OrderService struct {
	customers customer.CustomerRepository
}

type OrderConfiguration func(*OrderService) error

func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfgFunc := range configs {
		error := cfgFunc(os)
		if error != nil {
			return nil, error
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	memory_repository := memory.New()

	return WithCustomerRepository(memory_repository)
}