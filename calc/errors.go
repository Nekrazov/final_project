package calc

import "errors"

var (
	ErrNecorV = errors.New("Некорректный ввод")
	ErrDelZero   = errors.New("Нельзя делить на 0")
	ErrPystV   = errors.New("Пустой ввод")
)