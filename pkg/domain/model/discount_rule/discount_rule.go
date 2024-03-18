package model

type DiscountRule[T any] interface {
    Apply(product T) T
}

func CalculateDiscount[T any](product T, rule DiscountRule[T]) T {
    return rule.Apply(product)
}
