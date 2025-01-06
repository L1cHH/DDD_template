package memory

import (
	"errors"
	aggregate "impl_DDD/aggregate/customer"
	"impl_DDD/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		testName string
		id uuid.UUID
		expectedErr error
	}

	firstCustomer, _ := aggregate.NewCustomer("Alex", 20)
	secondCustomer, _ := aggregate.NewCustomer("Ron", 25)
	
	firstCustomerId, _ := firstCustomer.GetID()
	secondCustomerId, _ := secondCustomer.GetID()
	
	memoryRepo := MemoryRepository {
		customers: map[uuid.UUID]aggregate.Customer {
			firstCustomerId: firstCustomer,
		},
	}

	testCases := []testCase {
		{
			testName: "Correctly get Customer",
			id: firstCustomerId,
			expectedErr: nil,
		},

		{
			testName: "Getting customer with error occured",
			id: secondCustomerId,
			expectedErr: customer.ErrCustomerWasNotFound,
		},

	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			_, err := memoryRepo.Get(test.id)

			if !errors.Is(err, test.expectedErr) {
				t.Errorf("expected error: %v, got: %v", test.expectedErr, err)
			}
		})
	}
}