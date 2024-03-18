package model

import "time"

//bbcb2990-4074-489c-a9dc-8428f0c1bbf7
//135b7dc3-d6e4-43bf-960d-689c8e5da26f
type Product struct {
    ID               string                 `json:"id"`
    Name             string                 `json:"name"`
    BuyOptions       []BuyOption            `json:"buy_options"`
    Categories       []string               `json:"categories"`       
    CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
    CreatedAt        time.Time              `json:"created_at"`
    IsBlocked        bool                   `json:"is_blocked"`
    Images           []string               `json:"images"`
}

type BuyOption struct {
    SellerID            string `json:"seller_id"`
    Price               uint32 `json:"price"`
    PromotionalDiscount uint8  `json:"promotional_discount"`
    Stock               uint16 `json:"stock"`
    ProductID           string `json:"product_id"`
}