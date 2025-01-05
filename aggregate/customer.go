package aggregate

import (
	"impl_DDD/entity"
	"impl_DDD/valueobject"
)

type Customer struct {
	//person is the root entity of Customer
	//which means that Person.ID is the main identifier for the customer
	person *entity.Person
	items []*entity.Item

	transactions []valueobject.Transaction
}