package memory

import (
	"fmt"
	aggregate "impl_DDD/aggregate/customer"
	"impl_DDD/domain/customer"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	} 
	
	return aggregate.Customer{}, customer.ErrCustomerWasNotFound

}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {

	customerId, err := c.GetID()

	if err != nil {
		return fmt.Errorf("%v, tryind to add empty customer", customer.ErrFailedToAddCustomer)
	}

	if _, ok := mr.customers[customerId]; ok {
		return fmt.Errorf("%v, customer with that id already exists", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[customerId] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(id uuid.UUID) error {

	if customer, ok := mr.customers[id]; ok {
		
		mr.Lock()
		mr.customers[id] = customer
		mr.Unlock()
	} 

	return fmt.Errorf("%v, customer with that id doesn't exist", customer.ErrFailedToUpdateCustomer)
}