package model

type Repository[T any] interface {
	// GetAll() ([]*Product, error)
	GetByID(id string) (T, error)
	Create(entity T) (T, error)
}
