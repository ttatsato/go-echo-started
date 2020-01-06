package Controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)
type OrderParam struct {
	Products []Product `json:products`
}
type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
}

/**
 * オーダーを受け取る
 */
func SaveProduct(c echo.Context) error {
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
