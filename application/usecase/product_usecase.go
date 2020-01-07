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

type Product struct {
	Id uint `json:"id"`
	ShopId int `json:"shopId"`
	Name string `json:"name"`
	Price int `json:"price"`
}

func CreateNewProduct (c echo.Context) error {
	param := new(Product)
	if err := c.Bind(param); err != nil {
		log.Println("bad request")
		bytes, err := json.Marshal(param)
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
			return nil
		}
		return c.String(http.StatusBadRequest, string(bytes))
	}
	insertRecord := &domain.Product{ShopId: param.ShopId, Name:param.Name, Price: param.Price}
	log.Println(insertRecord)
	conn, err := config.ConnectMySql()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.ProductRepositoryWithRDB(conn)
	bytes, err := json.Marshal(repo.Create(insertRecord))
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return nil
	}
	return c.String(http.StatusCreated, string(bytes))
}

/**
 * 顧客の注文履歴を確認する
 */
func ListShopProduct(shopId int) []domain.Product {
	conn, err := config.ConnectMySql()
	if err != nil {
		return nil
	}
	defer conn.Close()
	repo := persistence.ProductRepositoryWithRDB(conn)
	res, err := repo.GetByShopId(shopId)
	if err != nil {
		return nil
	}
	return res
}

