package usecase

import (
	"app/config"
	"app/domain"
	"app/infrastructure/persistence"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)
type OrderSet struct {
	Order []OrderParam `json:order`
}
type OrderParam struct {
	UserId int `json:userId`
	ShopId int `json:shopId`
	Product Product `json:product`
	Memo string `json:memo`
}

/**
 * オーダーを注文する
 */
func MakeOrder(c echo.Context) error {
	param := new(OrderSet)
	log.Println(param)
	if err := c.Bind(param); err != nil {
		return nil
	}
	for _, order := range param.Order {
		log.Println(order)
		insertRecord := &domain.Order{
			ProductRefer: order.Product.Id,
			UserId: order.UserId,
			ShopId: order.ShopId,
			Memo:  order.Memo}
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

