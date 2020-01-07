package main


import (
	"app/application/usecase"
	"app/config"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/order", func(context echo.Context) error {
		return context.String(http.StatusOK, convertMapToJsonString(usecase.ListShopOrder()))
	})
	e.GET("/order/:customer_id", func(context echo.Context) error {
		return context.String(http.StatusOK, convertMapToJsonString(usecase.ListCustomerOrderHistory()))
	})
	e.POST("/order", usecase.MakeOrder)

	e.POST("/product/new", usecase.CreateNewProduct)
	migrate()
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id is " + id)
}


/**
 * JSON文字列にして返す
 */
func convertMapToJsonString(src interface{}) string {
	bytes, err := json.Marshal(src)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return ""
	}
	return string(bytes)
}

func migrate() {
	_, err := config.Migrate()
	if err != nil {
		panic("migrate error")
		return
	} else {
		log.Println("migrate")
	}
}