package product

import (
	aggregate "impl_DDD/aggregate/product"

	"github.com/google/uuid"
)

type ProductRepository interface {
	GetAll(uuid.UUID) ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Update(aggregate.Product) error
	Add(aggregate.Product) error
	Delete(uuid.UUID) error
}