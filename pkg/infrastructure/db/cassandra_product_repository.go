package repository

import (
	"encoding/json"
	"time"

	"github.com/diegobcaetano/product-service/internal/api/config"
	"github.com/diegobcaetano/product-service/internal/logging"
	model "github.com/diegobcaetano/product-service/pkg/domain/model/product"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type CassandraProductRepository struct {
	session *gocql.Session
	logger  logging.Logger
}

func NewCassandraProductRepository(session *gocql.Session, logger logging.Logger) *CassandraProductRepository {
	return &CassandraProductRepository{
		session: session,
		logger:  logger,
	}
}

func NewCassandraSession(env *config.Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(env.DBHost)
	cluster.Keyspace = env.DBKeyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	return session, err
}

func (r *CassandraProductRepository) GetByID(id string) (model.Product, error) {

	var product model.Product
	var customAttributesString string

	err := r.session.Query("SELECT id, name, categories, custom_attributes, created_at, is_blocked, images "+
		"FROM products WHERE id = ?", id).Scan(
		&product.ID,
		&product.Name,
		&product.Categories,
		&customAttributesString,
		&product.CreatedAt,
		&product.IsBlocked,
		&product.Images)
	if err != nil {
		r.logger.Error(err.Error())
	}

	if customAttributesString != "" {
		var customAttributes map[string]interface{}
		if err := json.Unmarshal([]byte(customAttributesString), &customAttributes); err != nil {
		} else {
			product.CustomAttributes = customAttributes
		}
	}
	product.BuyOptions, err = r.GetProductBuyOptionsByProductID(product.ID)
	return product, err
}

func (r *CassandraProductRepository) GetProductBuyOptionsByProductID(product_id string) ([]model.BuyOption, error) {
	var buyOptions []model.BuyOption
	var err error

	scanner := r.session.Query("SELECT seller_id, product_id, price, promotional_discount, stock "+
		"FROM buy_options WHERE product_id = ?", product_id).Iter().Scanner()

	for scanner.Next() {
		buyOption := model.BuyOption{}
		err := scanner.Scan(
			&buyOption.SellerID,
			&buyOption.ProductID,
			&buyOption.Price,
			&buyOption.PromotionalDiscount,
			&buyOption.Stock)
		if err != nil {
			r.logger.Error(err.Error())
		}
		buyOptions = append(buyOptions, buyOption)
	}
	return buyOptions, err
}

func (r *CassandraProductRepository) Create(product model.Product) (model.Product, error) {
	product.ID = uuid.New().String()
	product.CreatedAt = time.Now()
	var customAttributesBytes []byte

	if product.CustomAttributes != nil {
		var err error
		customAttributesBytes, err = json.Marshal(product.CustomAttributes)
		if err != nil {
			r.logger.Error("Failed to Marshal the custom_attributes", "error", err)
		}
	}

	err := r.session.Query("INSERT INTO products (id, name, categories, custom_attributes, created_at, is_blocked, images)"+
		"VALUES (?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS",
		product.ID,
		product.Name,
		product.Categories,
		string(customAttributesBytes),
		product.CreatedAt,
		product.IsBlocked,
		product.Images).Exec()

	if err != nil {
		r.logger.Error("Failed to insert a new record on products table", "error", err.Error())
		product.ID = ""
	}

	product.BuyOptions, _ = r.CreateProductBuyOptions(product.ID, product.BuyOptions)

	if err != nil {
		panic("Failed to insert a product")
	}

	return product, err
}

func (r *CassandraProductRepository) CreateProductBuyOptions(
	productId string,
	buyOptions []model.BuyOption) ([]model.BuyOption, error) {

	//Make it a batch operation
	//Make it with goroutine
	var resultErr error
	for i := range buyOptions {
		buyOptions[i].ProductID = productId
		err := r.session.Query("INSERT INTO buy_options (seller_id, product_id, price, promotional_discount, stock)"+
			"VALUES (?, ?, ?, ?, ?) IF NOT EXISTS",
			buyOptions[i].SellerID,
			buyOptions[i].ProductID,
			buyOptions[i].Price,
			buyOptions[i].PromotionalDiscount,
			buyOptions[i].Stock).Exec()

		if err != nil {
			r.logger.Error("Failed to insert a new record on buy_options table", "error", err.Error())
			resultErr = err
			break
		}
	}
	return buyOptions, resultErr
}
