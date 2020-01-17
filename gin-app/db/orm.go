package db

import (
	"errors"
	"go-music/gin-app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con+"?parseTime=true")
	return &DBORM{
		DB: db,
	}, err
}

// Get a list of products.
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

// Get a list of promotions.
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

// Get the customer by name.
func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

// Get the customer by id.
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

// Get the product by id.
func (db *DBORM) GetProduct(id int) (product models.Product, error error) {
	return product, db.First(&product, id).Error
}

// Add a new customer.
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

// Get the hash value of a password.
func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	sBytes := []byte(*s)                                                        // Convert password string to byte slice.
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost) // Obtain hashed password.
	if err != nil {
		return err
	}

	*s = string(hashedBytes[:])                                                 // Update password string with the hashed version.
	return nil
}

// Sign in a customer.
func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	if !checkPassword(customer.Pass, pass) {
		return customer, ErrINVALIDPASSWORD
	}

	customer.Pass = ""

	err = result.Update("loggedin", 1).Error                             // Update the loggedin field.
	if err != nil {
		return customer, err
	}

	return customer, result.Find(&customer).Error
}

// Check the password is correct or not.
// This method will return an error if the hash does not match the provided password string.
func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}

// Sign out a customer.
func (db *DBORM) SignOutUserById(id int) error {
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

// Get all the orders for a specific customer.
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error
}

// Add a new order.
func (db *DBORM) AddOrder(order models.Order) error {
	return db.Create(&order).Error
}

// Get a customer saved credit card number.
func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	cusomterWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}

	return cusomterWithCCID.CCID, db.First(&cusomterWithCCID, id).Error
}

// Save a customer credit card number.
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}