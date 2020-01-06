package controllers

import (
	"app/models"
	"github.com/jinzhu/gorm"
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
func SaveProduct(c echo.Context) error {
	param := new(OrderParam)
	if err := c.Bind(param); err != nil {
		return err
	}
	for _, product := range param.Products {
		insertRecord := &models.Order{
			Product: models.Product{
				Model: gorm.Model{},
				Name: product.Name,
				Code:  "code",
				Price: product.Price},
			UserId: 12,
			ShopId: 133}
		models.CreateNewOrder(insertRecord)
	}
	return c.String(http.StatusOK, "name:")
}
