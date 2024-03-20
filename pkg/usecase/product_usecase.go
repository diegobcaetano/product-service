package usecase

import (
	"github.com/diegobcaetano/product-service/internal/common_errors"
	DiscountRule "github.com/diegobcaetano/product-service/pkg/domain/model/discount_rule"
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
	ri "github.com/diegobcaetano/product-service/pkg/domain/repository"
)

type ProductUseCase interface {
	GetProductByID(id string) (model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	ProductHasStock(id string, sellerID string) (bool, error)
}

type ProductService struct {
	repository ri.Repository[model.Product]
}

func NewProductService(r ri.Repository[model.Product]) *ProductService {
	return &ProductService{
		repository: r,
	}
}

func (s *ProductService) GetProductByID(id string) (model.Product, error) {
	product, err := s.repository.GetByID(id)
	if err != nil {
		return model.Product{}, err
	}
	return s.SetDiscount(product)
}

func (s *ProductService) ProductHasStock(id string, sellerID string) (bool, error) {
	result := false
	buyOption, err := s.repository.GetBuyOptionByProductIdAndSellerId(id, sellerID)
	if err != nil {
		return result, err
	}

	if buyOption.Stock > 0 {
		result = true
	}

	return result, nil
}

func (s *ProductService) CreateProduct(product model.Product) (model.Product, error) {

	if len(product.BuyOptions) < 1 {
		return model.Product{}, common_errors.NewCommonErrors(
			common_errors.ErrorCode(model.ErrBuyOptionNotFound),
			"The number of buy options should more than 0")
	}

	for _, bo := range product.BuyOptions {
		if bo.SellerID == "" {
			return model.Product{}, common_errors.NewCommonErrors(
				common_errors.ErrorCode(model.ErrSellerIdNotFound),
				"SellerID must be defined for each buy option")
		}
	}

	return s.repository.Create(product)
}

func (s *ProductService) SetDiscount(product model.Product) (model.Product, error) {
	for _, category := range product.Categories {
		switch category {
		case "book":
			product = DiscountRule.CalculateDiscount(product, DiscountRule.BookDiscountRule{})
		case "eletronic":
			product = DiscountRule.CalculateDiscount(product, DiscountRule.EletronicDiscountRule{})
		}
	}

	return product, nil
}
