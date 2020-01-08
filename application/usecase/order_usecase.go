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
type OrderSet struct {
	Order []domain.Order `json:order`
}

/**
 * オーダーを注文する
 */
func MakeOrder(c echo.Context) error {
	param := new(OrderSet)
	if err := c.Bind(param); err != nil {
		return nil
	}
	for _, order := range param.Order {
		insertRecord := &domain.Order{
			ProductRefer: order.Product.ID,
			UserId: order.UserId,
			ShopId: order.ShopId,
			Code: order.Code,
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
func ListShopOrder(shopId int) []domain.Order {
	conn, err := config.ConnectMySql()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.OrderRepositoryWithRDB(conn)
	res, err := repo.GetByShopId(shopId)
	if err != nil {
		return nil
	}
	return res
}

/**
 * 顧客の注文履歴を確認する
 */
func ListCustomerOrderHistory(customerId int) []domain.Order {
	conn, err := config.ConnectMySql()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.OrderRepositoryWithRDB(conn)
	res, err := repo.GetByCustomerIdSortByCreatedAt(customerId)
	if err != nil {
		return nil
	}
	return res
}

