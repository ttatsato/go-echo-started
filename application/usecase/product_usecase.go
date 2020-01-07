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
	Name string `json:"name"`
	Price int `json:"price"`
}

func CreateNewProduct (c echo.Context) error {
	param := new(Product)
	log.Println(param)
	if err := c.Bind(param); err != nil {
		return nil
	}
	insertRecord := &domain.Product{Name:param.Name, Price: param.Price}
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