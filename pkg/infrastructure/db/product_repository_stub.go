package repository

import (
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
	repository "github.com/diegobcaetano/product-service/pkg/domain/repository"
)

type ProductRepositoryStub struct {
	repository.Repository[model.Product]
}

func (r *ProductRepositoryStub) Create(product model.Product) (model.Product, error) {
	return model.ValidProduct, nil
}

func (r *ProductRepositoryStub) GetByID(id string) (model.Product, error) {

	switch id {
	case "with-book-category":
		return model.StubProductWithCustomCategory(model.ValidProduct, "book"), nil

	}
	return model.ValidProduct, nil
}
