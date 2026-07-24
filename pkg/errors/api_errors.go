package errors

import "errors"

var (
	BadRequestError     = errors.New("Incorrect input data")
	InternalServerError = errors.New("Internal Server Erorr")
	NotFound            = errors.New("Record not found")
)
