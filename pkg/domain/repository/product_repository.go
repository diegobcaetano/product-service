package model

import model "github.com/diegobcaetano/product-service/pkg/domain/model/product"

type Repository[T any] interface {
	// GetAll() ([]*Product, error)
	GetByID(id string) (T, error)
	Create(entity T) (T, error)
	GetBuyOptionByProductIdAndSellerId(id string, sellerID string) (model.BuyOption, error)
}
