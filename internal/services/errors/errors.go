package errorsapi

import "errors"

var (
	ErrDuplicatedEmailOrUsername    = errors.New("email or username already exists")
	ErrInvalidCredentials           = errors.New("invalid credentials")
	ErrSellerNotExist               = errors.New("seller id not exist")
	ErrProductNotExist              = errors.New("product id not exist")
	ErrBidderNotExist               = errors.New("bidder id not exist")
	ErrBidLowerThanTheLast          = errors.New("bid must be higher than the one already launched")
	ErrBidLowerThanProductBaseprice = errors.New("bid must be higher than the product baseprice")
	ErrSomethingWentWrong           = errors.New("something went wrong")
)
