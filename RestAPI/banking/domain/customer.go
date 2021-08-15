package domain

import "github.com/Rafael-Lopez/GO/RESTAPI/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	// We use *Customer because we want to be able to return 'nil' if there's no customer for a given ID.
	// And you can only do that with a pointer
	ById(string) (*Customer, *errs.AppError)
}
