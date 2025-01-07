package product

import (
	"errors"
	"impl_DDD/entity"

	"github.com/google/uuid"
)

var (
	ErrMissingPriceOrQuantity = errors.New("price or quantity are 0")
)

type Product struct {
	item *entity.Item
	price float64
	quantity uint
}

func NewProduct(name string, description string, price float64, quantity uint) (Product, error) {

	if price == 0.0 || quantity == 0 {
		return Product{}, ErrMissingPriceOrQuantity
	} 

	item := &entity.Item{ 
		ID: uuid.New(),
		Name: name,
		Description: description,
	}
	
	product := Product {
		item: item,
		price: price,
		quantity: quantity,
	}

	return product, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}