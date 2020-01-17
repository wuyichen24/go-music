package rest

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunMockAPI(address string) error {
	h := NewMockHandler()
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	r.Use(static.ServeRoot("/", "../public/build"))
	return r.Run(address)
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("************************************")
		c.Next()
		fmt.Println("************************************")
	}
}
