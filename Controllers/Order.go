package controllers

import (
	"app/models"
	"encoding/json"
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
 * オーダーを受け取りmySqlに格納する。
 */
func MakeOrder(c echo.Context) error {
	param := new(OrderParam)
	if err := c.Bind(param); err != nil {
		return err
	}
	for _, product := range param.Products {
		insertRecord := &models.Order{
			Product: models.Product{
				Name: product.Name,
				Code:  "code",
				Price: product.Price},
			UserId: 12,
			ShopId: 133,
			Memo: "これはメモです"}
		models.CreateNewOrder(insertRecord)
	}
	bytes, err := json.Marshal(param)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return nil
	}
	return c.String(http.StatusCreated, string(bytes))
}

func FetchOrder() []models.Order {
	return models.ReadOrder()
}

func FetchOrderHistory() []models.Order {
	return models.ReadOrderByUser()
}

