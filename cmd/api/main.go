package main

import (
	"fmt"
	"net/http"

	"github.com/diegobcaetano/product-service/internal/api/config"
	"github.com/diegobcaetano/product-service/internal/logging"
	"github.com/diegobcaetano/product-service/pkg/handler"
	repository "github.com/diegobcaetano/product-service/pkg/infrastructure/db"
	"github.com/diegobcaetano/product-service/pkg/usecase"
)

func main() {
	//test
	logger := logging.NewSlogAdapter()
	env := config.NewEnvLoad(logger).Load("../../")
	fmt.Printf(env.GetConfig().DBHost)
	dbSession, err := repository.NewCassandraSession(env.GetConfig())
	defer dbSession.Close()
	if err != nil {
		panic(err)
	}
	productRepository := repository.NewCassandraProductRepository(dbSession, logger)
	var useCase usecase.ProductUseCase = usecase.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(useCase, logger)
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetProduct(w, r)
		case "POST":
			productHandler.CreateProduct(w, r)
		default:
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
