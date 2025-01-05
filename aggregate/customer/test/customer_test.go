package aggregate_test

import (
	"errors"
	aggregate "impl_DDD/aggregate/customer"
	"testing"
)


func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test string
		name string
		age uint
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "wrong age",
			name: "Kirill",
			age: 14,
			expectedErr: aggregate.ErrInvalidPerson,
		},

		{
			test: "wrong name",
			name: "",
			age: 25,
			expectedErr: aggregate.ErrInvalidPerson,
		},

		{
			test: "wrong name and age",
			name: "",
			age: 4,
			expectedErr: aggregate.ErrInvalidPerson,
		},

		{
			test: "correct customer",
			name: "Kirill",
			age: 18,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name, tc.age)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}