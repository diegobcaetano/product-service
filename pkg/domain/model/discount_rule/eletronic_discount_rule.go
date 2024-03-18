package model

import (
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
)

type EletronicDiscountRule struct{}

func (e EletronicDiscountRule) Apply(product model.Product) model.Product {

	for i := range product.BuyOptions {
		product.BuyOptions[i].PromotionalDiscount += 10
	}

	return product
}
