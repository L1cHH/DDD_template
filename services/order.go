package services

import (
	_ "impl_DDD/aggregate/customer"
	aggregateProduct "impl_DDD/aggregate/product"
	"impl_DDD/domain/customer"
	mcr "impl_DDD/domain/customer/memory" //MemoryCustomerRepository
	"impl_DDD/domain/product"
	mpr "impl_DDD/domain/product/memory" //MemoryProductRepository

	"github.com/google/uuid"
)

type OrderService struct {
	customers customer.CustomerRepository
	products product.ProductRepository
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

func WithProductRepository(products []aggregateProduct.Product, pr product.ProductRepository) OrderConfiguration {
	return func(os *OrderService) error {

		os.products = pr

		for _, p := range products {
			err := os.products.Add(p)
			if err != nil {
				return err
			}
		} 

		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	mcr := mcr.New()

	return WithCustomerRepository(mcr)
}

func WithMemoryProductRepository(products []aggregateProduct.Product) OrderConfiguration {
	mpr := mpr.New()

	return WithProductRepository(products, mpr)
}



func (os *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	_, err := os.customers.Get(customerID)

	if err != nil {
		return err
	}

	var products []aggregateProduct.Product
	var price float64

	for _, productID := range productIDs {
		product, err := os.products.GetByID(productID)

		if err != nil {
			return err
		}
		price += product.GetPrice()
		products = append(products, product)
	}


	return nil
}