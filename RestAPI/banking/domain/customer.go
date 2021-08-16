package domain

import "github.com/Rafael-Lopez/GO/RESTAPI/banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	// We use *Customer because we want to be able to return 'nil' if there's no customer for a given ID.
	// And you can only do that with a pointer
	ById(string) (*Customer, *errs.AppError)
}
