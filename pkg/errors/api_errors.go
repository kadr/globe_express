package api_errors

import "errors"

var (
	ErrorBadRequest        = errors.New("incorrect input data")
	ErrorInternalServer    = errors.New("internal Server Erorr")
	ErrorNotFound          = errors.New("record not found")
	ErrorIncorrectStatus   = errors.New("incorrect status")
	ErrorCanNotCancel      = errors.New("can't cancel ongoing or completed trip")
	ErrorFieldValidation   = errors.New("validation error")
	ErrorCantDeleteProduct = errors.New("can't delete product which reference to order")
)
