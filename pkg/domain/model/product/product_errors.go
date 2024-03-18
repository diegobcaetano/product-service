package model

type ErrorCode string

const (
	ErrBuyOptionNotFound ErrorCode = "BUY_OPTION_NOT_FOUND"
	ErrSellerIdNotFound ErrorCode = "SELLER_ID_NOT_FOUND"
	ErrInvalidDiscountValue ErrorCode = "INVALID_DISCOUNT_VALUE"
	ErrNotEnoughImages ErrorCode = "NOT_ENOUGH_IMAGES"
)