package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/order", saveProduct)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id is " + id)
}


type OrderParam struct {
	Products []Product `json:products`
}
type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
}
/**
curl -F "name=ハンバーグ" -F "price=1280" http://localhost:1323/order
 */
func saveProduct(c echo.Context) error {
	param := new(OrderParam)
	if err := c.Bind(param); err != nil {
		return err
	}
	fmt.Println(param)
	for key, product := range param.Products {
		fmt.Print(key)
		fmt.Print(": ")
		fmt.Print(product.Name)
		fmt.Print(", ")
		fmt.Print(product.Price)
		fmt.Println(" ")
	}
	return c.String(http.StatusOK, "name:")
}