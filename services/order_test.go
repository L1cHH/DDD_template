package services

import (
	aggregateCustomer "impl_DDD/aggregate/customer"
	aggregateProduct "impl_DDD/aggregate/product"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) (products []aggregateProduct.Product) {
	chips, err := aggregateProduct.NewProduct("Lays", "Unhealthy chips", 120.0, 15)

	if err != nil {
		t.Fatal(err)
	}

	chocolate, err := aggregateProduct.NewProduct("Alpen GOld", "chocolate", 60, 55)

	if err != nil {
		t.Fatal(err)
	}

	pepsi, err := aggregateProduct.NewProduct("Pepsi", "fizzy water", 120.0, 15)

	if err != nil {
		t.Fatal(err)
	}

	return []aggregateProduct.Product{
		chips, chocolate, pepsi,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryProductRepository(products),
		WithMemoryCustomerRepository(),
	)

	if err != nil {
		t.Fatal(err)
	}
		
	cust, err := aggregateCustomer.NewCustomer("Lucy", 20)

	if err != nil {
		t.Fatal(err)
	}

	creatingErr := os.customers.Add(cust)

	if creatingErr != nil {
		t.Fatal(creatingErr)
	}

	firstProductId, err := products[0].GetID()
	if err != nil {
		t.Fatal(err)
	}

	secondProductId, err := products[1].GetID()
	if err != nil {
		t.Fatal(err)
	}

	thirdProductId, err := products[2].GetID()
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID {
		firstProductId, secondProductId, thirdProductId,
	}

	custID, err := cust.GetID()
	if err != nil {
		t.Fatal(err)
	}

	err = os.CreateOrder(custID, order)
	if err != nil {
		t.Fatal(err)
	}
	
}

