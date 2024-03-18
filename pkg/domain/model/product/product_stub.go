package model

import (
	"time"

	"github.com/google/uuid"
)

var ValidProduct Product = Product{
	Name: "Test Product",
	ID: "3c32517c-146c-4258-999a-ec04ab9deb66",
	IsBlocked: false,
	CreatedAt: time.Date(2024, time.March, 10, 15, 30, 0, 0, time.UTC),
	Images: []string{
		"http://fakeimage.com/14-20.png", 
		"http://fakeimage.com/14-20.png"},
	Categories: []string{
		"dummy",
		"category",
	},
	CustomAttributes: map[string]interface{} { 
		"key_1" : "value_1", 
	},
	BuyOptions: []BuyOption {
		{
			SellerID:            "seller1",
			Price:               100,
			PromotionalDiscount: 10,
			Stock:               20,
			ProductID:           "product1",
		},
		{
			SellerID:            "seller2",
			Price:               200,
			PromotionalDiscount: 20,
			Stock:               0,
			ProductID:           "product2",
		},
	},
}

func StubProductWithoutBuyOptions(p Product) Product {
	p.BuyOptions = nil
	return p
}

func StubProduct(p Product) Product {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	return p
}

func StubProductWithCustomSellerID(p Product, sellerID string, index int) Product {
	
    buyOptions := make([]BuyOption, len(p.BuyOptions))
    copy(buyOptions, p.BuyOptions)
	
    if index < len(buyOptions) {
        buyOptions[index].SellerID = sellerID
    }

    p.BuyOptions = buyOptions
    return p
}

func StubProductWithCustomCategory(p Product, category string) Product {
	
	categories := make([]string, len(p.Categories))
	copy(categories, p.Categories)
	categories = append(categories, category)
	p.Categories = categories
	return p
}