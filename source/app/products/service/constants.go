package service

import "errors"

var (
	ErrCouldNotAddProduct      = errors.New("can not perform create operation right now, try again")
	ErrCouldNotDeleteProduct   = errors.New("can not perform delete operation")
	ErrCouldNotUpdateProduct   = errors.New("can not perform update operation right now, try again")
	ErrInvalidProductId        = errors.New("invalid product id is passed")
	ErrProductValidationFailed = errors.New("failed to validate product before persisting")
)
