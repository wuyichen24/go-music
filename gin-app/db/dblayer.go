package db

import (
	"errors"
	"go-music/gin-app/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)                      // Get a list of products.
	GetPromos() ([]models.Product, error)                           // Get a list of promotions.
	GetCustomerByName(string, string) (models.Customer, error)      // Get the customer by name.
	GetCustomerByID(int) (models.Customer, error)                   // Get the customer by id.
	GetProduct(int) (models.Product, error)                         // Get the product by id.
	AddUser(models.Customer) (models.Customer, error)               // Add a new customer.
	SignInUser(username, password string) (models.Customer, error)  // Sign in a customer.
	SignOutUserById(int) error                                      // Sign out a customer.
	GetCustomerOrdersByID(int) ([]models.Order, error)              // Get all the orders for a specific customer.
	AddOrder(models.Order) error                                    // Add a new order.
	GetCreditCardCID(int) (string, error)                           // Get a customer saved credit card number.
	SaveCreditCardForCustomer(int, string) error                    // Save the stripe customer ID to our database
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")
