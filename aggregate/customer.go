package aggregate

import (
	"impl_DDD/entity"
	"impl_DDD/valueobject"

	"github.com/google/uuid"
)

type Customer struct {
	//person is the root entity of Customer
	//which means that Person.ID is the main identifier for the customer
	person *entity.Person
	items []*entity.Item

	transactions []valueobject.Transaction
}



func NewCustomer(name string, age uint) (Customer, error) {
	if name == "" || age < 18 {
		return Customer{}, nil
	}

	person := &entity.Person {
		Name: name,
		ID: uuid.New(),
	}

	return Customer{
		person: person,
		items: make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}