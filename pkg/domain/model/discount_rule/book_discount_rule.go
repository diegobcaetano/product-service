package model

import (
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
)

const BookDiscountValue = 30

type BookDiscountRule struct {
	// includedSellers []string
}

func (e BookDiscountRule) Apply(product model.Product) model.Product {

	buyOptions := make([]model.BuyOption, len(product.BuyOptions))
	copy(buyOptions, product.BuyOptions)
	for i := range buyOptions {
		buyOptions[i].PromotionalDiscount += BookDiscountValue
	}
	product.BuyOptions = buyOptions
	return product
}

// func (e BookDiscountRule) defineParticipants() {
// 	//retrieve from db...
// 	e.includedSellers = []string{ "10", "11", "12", "13" }
// }
