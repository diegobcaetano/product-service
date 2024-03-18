package usecase

import (
	"testing"

	"github.com/diegobcaetano/product-service/internal/common_errors"
	modelDiscount "github.com/diegobcaetano/product-service/pkg/domain/model/discount_rule"
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
	repository "github.com/diegobcaetano/product-service/pkg/infrastructure/db"

	"github.com/google/go-cmp/cmp"
)

func TestCreateProduct(t *testing.T) {
	data := []struct {
		name          string
		product       model.Product
		expectedError model.ErrorCode
	}{
		{
			"Created a product",
			model.ValidProduct,
			"",
		},
		{
			"Failed to create a product, missing buyoption",
			model.StubProductWithoutBuyOptions(model.ValidProduct),
			model.ErrBuyOptionNotFound,
		},
		{
			"Failed to create a product, missing sellerID",
			model.StubProductWithCustomSellerID(model.ValidProduct, "", 1),
			model.ErrSellerIdNotFound,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			productService := NewProductService(&repository.ProductRepositoryStub{})
			product, err := productService.CreateProduct(d.product)
			if d.expectedError != "" {
				if customErr, ok := err.(*common_errors.CommonErrors); ok {
					if model.ErrorCode(customErr.Code) != d.expectedError {
						t.Errorf("Código de erro esperado: %s, obtido: %s", d.expectedError, customErr.Code)
					}
				} else {
					t.Error("Erro não é uma instância de CustomError")
				}
			} else {
				if diff := cmp.Diff(d.product, product); diff != "" {
					t.Error(diff)
				}
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	data := []struct {
		name                       string
		id                         string
		withBuyOptionsModification bool
	}{
		{
			"Get a product",
			"",
			false,
		},
		{
			"Get a product with discount in `book` category",
			"with-book-category",
			true,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			productService := NewProductService(&repository.ProductRepositoryStub{})
			product, _ := productService.GetProductByID(d.id)
			if d.withBuyOptionsModification {
				for i, bo := range product.BuyOptions {
					expectedDiscount := model.ValidProduct.BuyOptions[i].PromotionalDiscount + modelDiscount.BookDiscountValue
					if diff := cmp.Diff(expectedDiscount, bo.PromotionalDiscount); diff != "" {
						t.Error(diff)
					}
				}
			} else {
				if diff := cmp.Diff(model.ValidProduct, product); diff != "" {
					t.Error(diff)
				}
			}
		})
	}
}
