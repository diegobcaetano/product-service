package model

import (
    pb "github.com/diegobcaetano/product-service/pkg/domain/grpc"
    "time"
)


//bbcb2990-4074-489c-a9dc-8428f0c1bbf7
//135b7dc3-d6e4-43bf-960d-689c8e5da26f
// 8f1e6c49-7810-4e8f-9698-3bb9ca3d6e25
//  ba67cd35-c6db-42c4-bc57-70a2e5aa46ed
//  16bd1f3c-be0c-4cd7-9154-1288efca01cc
//  a231f5ba-45d8-407e-87f4-c0eeafcfb868
//  5e697955-cbed-493c-ac52-01ef05951126
//  391b2888-d011-4fc9-a604-42a805358c28
//  188c2e49-79c2-499d-bf61-21cc557bfe2f
//  9159a6c6-e8c4-477e-998a-f9ec188cb75e
//  99d8fcba-5d72-45d8-b9e1-f436bc25a4fa
//  55a34d3d-cb82-475f-98d9-8d79c54ceebf
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

func ProtoToProduct(protoProduct *pb.Product) (Product, error) {
    product := Product{
        ID:               protoProduct.GetId(),
        Name:             protoProduct.GetName(),
        Categories:       protoProduct.GetCategories(),
        CustomAttributes: make(map[string]interface{}),
        CreatedAt:        time.Unix(protoProduct.GetCreatedAt(), 0),
        IsBlocked:        protoProduct.GetIsBlocked(),
        Images:           protoProduct.GetImages(),
    }

    for key, value := range protoProduct.GetCustomAttributes() {
        product.CustomAttributes[key] = value
    }

    for _, protoBuyOption := range protoProduct.GetBuyOptions() {
        buyOption := BuyOption{
            SellerID:            protoBuyOption.GetSellerId(),
            Price:               protoBuyOption.GetPrice(),
            PromotionalDiscount: uint8(protoBuyOption.GetPromotionalDiscount()),
            Stock:               uint16(protoBuyOption.GetStock()),
            ProductID:           protoBuyOption.GetProductId(),
        }
        product.BuyOptions = append(product.BuyOptions, buyOption)
    }

    return product, nil
}

func ProductToProto(product Product) (*pb.Product, error) {

    copyImages := make([]string, len(product.Images))
    copy(copyImages, product.Images)

    copyCategories := make([]string, len(product.Categories))
    copy(copyCategories, product.Categories)
    
    protoProduct := &pb.Product{
        Id:               product.ID,
        Name:             product.Name,
        CustomAttributes: make(map[string]string),
        CreatedAt:        product.CreatedAt.Unix(),
        IsBlocked:        product.IsBlocked,  
        Categories:       copyCategories,
        Images:           copyImages,        
    }

    for key, value := range product.CustomAttributes {
        if stringValue, ok := value.(string); ok {
            protoProduct.CustomAttributes[key] = stringValue
        } 
    }

    for _, buyOption := range product.BuyOptions {
        protoBuyOption := &pb.BuyOption{
            SellerId:            buyOption.SellerID,
            Price:               buyOption.Price,
            PromotionalDiscount: uint32(buyOption.PromotionalDiscount),
            Stock:               uint32(buyOption.Stock),
            ProductId:           buyOption.ProductID,
        }
        protoProduct.BuyOptions = append(protoProduct.BuyOptions, protoBuyOption)
    }

    return protoProduct, nil
}
