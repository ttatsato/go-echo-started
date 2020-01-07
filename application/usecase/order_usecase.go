package usecase

import (
	"app/config"
	"app/domain"
	"app/infrastructure/persistence"
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
 * オーダーを注文する
 */
func MakeOrder(c echo.Context) error {
	param := new(OrderParam)
	if err := c.Bind(param); err != nil {
		return nil
	}
	for _, product := range param.Products {
		insertRecord := &domain.Order{
			Product: domain.Product{
				Name:  product.Name,
				Code:  "code",
				Price: product.Price},
			UserId: 12,
			ShopId: 133,
			Memo:   "これはメモです"}
		conn, err := config.ConnectMySql()
		if err != nil {
			return nil
		}
		defer conn.Close()
		repo := persistence.OrderRepositoryWithRDB(conn)
		repo.Save(insertRecord)
	}
	bytes, err := json.Marshal(param)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return nil
	}
	return c.String(http.StatusCreated, string(bytes))
}

/**
 * ショップのオーダーを一覧確認する
 * TODO 引数にショップIDを指定
 */
func ListShopOrder() []domain.Order {
	conn, err := config.ConnectMySql()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.OrderRepositoryWithRDB(conn)
	res, err := repo.GetByShopId(133)
	if err != nil {
		return nil
	}
	return res
}

/**
 * 顧客の注文履歴を確認する
 * @TODO 引数に顧客IDを指定
 */
func ListCustomerOrderHistory() []domain.Order {
	conn, err := config.ConnectMySql()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.OrderRepositoryWithRDB(conn)
	res, err := repo.GetByCustomerId(12)
	if err != nil {
		return nil
	}
	return res
}

