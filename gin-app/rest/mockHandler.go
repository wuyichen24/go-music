package rest

import (
	"encoding/json"
	"go-music/gin-app/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MockHandler struct {
}

func NewMockHandler() *MockHandler {
	return &MockHandler{}
}

func (mh *MockHandler) GetMainPage(c *gin.Context) {

}
func (mh *MockHandler) GetProducts(c *gin.Context) {

}
func (mh *MockHandler) GetPromos(c *gin.Context) {

}
func (mh *MockHandler) AddUser(c *gin.Context) {
	f, err := os.Open("mockdata.json")
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "mock data file not found"})
		return
	}
	var user models.Customer
	err = json.NewDecoder(f).Decode(&user)
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not decode mock data json file"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (mh *MockHandler) SignIn(c *gin.Context) {
	f, err := os.Open("mockdata.json")
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "mock data file not found"})
		return
	}
	var user models.Customer
	err = json.NewDecoder(f).Decode(&user)
	if err != nil {
		log.Println("Could not add user", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "could not decode mock data json file"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (mh *MockHandler) SignOut(c *gin.Context) {

}
func (mh *MockHandler) GetOrders(c *gin.Context) {

}
func (mh *MockHandler) Charge(c *gin.Context) {

}
