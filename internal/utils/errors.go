package utils

import "errors"

var (
	UserNotFound   = errors.New("user not found")
	NotEnoughFunds = errors.New("insufficient funds")
)
