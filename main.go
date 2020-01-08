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
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/shop/:shopId", func(context echo.Context) error {
		shopId, _ := strconv.Atoi(context.Param("shopId"))
		return context.String(http.StatusOK, convertMapToJsonString(usecase.FetchShopInfo(shopId)))
	})
	e.GET("/shop/order/:shopId", func(context echo.Context) error {
		shopId, _ := strconv.Atoi(context.Param("shopId"))
		return context.String(http.StatusOK, convertMapToJsonString(usecase.ListShopOrder(shopId)))
	})
	e.GET("/order/:customerId", func(context echo.Context) error {
		customerId, _ := strconv.Atoi(context.Param("customerId"))
		return context.String(http.StatusOK, convertMapToJsonString(usecase.ListCustomerOrderHistory(customerId)))
	})
	e.POST("/order", usecase.MakeOrder)
	e.GET("/product/:shopId", func(context echo.Context) error {
		shopId, _ := strconv.Atoi(context.Param("shopId"))
		return context.String(http.StatusOK, convertMapToJsonString(usecase.ListShopProduct(shopId)))
	})
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