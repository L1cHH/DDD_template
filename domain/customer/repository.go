package customer

import (
	"errors"
	aggregate "impl_DDD/aggregate/customer"

	"github.com/google/uuid"
)

var (
	ErrCustomerWasNotFound = errors.New("customer wasn't found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)


type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(*aggregate.Customer) error
	Update(uuid.UUID) error
}
