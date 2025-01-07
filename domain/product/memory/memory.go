package memory

import (
	aggregate "impl_DDD/aggregate/product"
	"impl_DDD/domain/product"
	"sync"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id] ;ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductWasNotFound
}

func (mpr *MemoryProductRepository) Add(p aggregate.Product) error {
	productId, err := p.GetID()

	if err != nil {
		return product.ErrFailedToAddProduct
	}

	if _, ok := mpr.products[productId]; !ok {

		mpr.Lock()
		mpr.products[productId] = p
		mpr.Unlock()
		return nil
	}

	return product.ErrFailedToAddProduct
}

func (mpr *MemoryProductRepository) Update(p aggregate.Product) error {
	productId, err := p.GetID()

	if err != nil {
		return product.ErrFailedToUpdateProduct
	}

	if _, ok := mpr.products[productId]; ok {
		mpr.Lock()
		mpr.products[productId] = p
		mpr.Unlock()
		return nil
	}

	return product.ErrFailedToUpdateProduct
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {

	if _, ok := mpr.products[id]; ok {
		mpr.Lock()
		delete(mpr.products, id)
		mpr.Unlock()
		return nil
	}

	return product.ErrFailedToDeleteProduct
}