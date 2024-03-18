package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/diegobcaetano/product-service/internal/common_errors"
	"github.com/diegobcaetano/product-service/internal/logging"
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
	"github.com/diegobcaetano/product-service/pkg/usecase"
)

type ProductHandler struct {
	useCase usecase.ProductUseCase
	logger  logging.Logger
}

func NewProductHandler(
	usecase usecase.ProductUseCase,
	logger logging.Logger) *ProductHandler {

	return &ProductHandler{
		useCase: usecase,
		logger:  logger,
	}
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/product/")
	h.logger.Info("Fetching up the product", "id", id)
	p, err := h.useCase.GetProductByID(id)
	if err != nil {
		h.logger.Error("Something went wrong while fetching the product", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if p.ID == "" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)
	product, err := h.useCase.CreateProduct(product)
	if err != nil {
		var commonError *common_errors.CommonErrors
		if errors.As(err, &commonError) {
			switch commonError.Code {
			case common_errors.ErrorCode(model.ErrBuyOptionNotFound):
				http.Error(w, commonError.Message, http.StatusBadRequest)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(product)
}
