package product

import (
	"errors"
	aggregate "impl_DDD/aggregate/product"

	"github.com/google/uuid"
)

var (
	ErrProductWasNotFound = errors.New("product with this id was not found")
	ErrFailedToAddProduct = errors.New("failed to add new product")
	ErrFailedToUpdateProduct = errors.New("failed to update product")
	ErrFailedToDeleteProduct = errors.New("failed to delete product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Update(aggregate.Product) error
	Add(aggregate.Product) error
	Delete(uuid.UUID) error
}