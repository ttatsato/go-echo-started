package main


import (
	"app/controllers"
	"app/db"
	"app/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
		return context.String(http.StatusOK, convertMapToJsonString(controllers.FetchOrder()))
	})
	e.GET("/order/:user_id", func(context echo.Context) error {
		return context.String(http.StatusOK, convertMapToJsonString(controllers.FetchOrderHistory()))
	})
	e.POST("/order", controllers.MakeOrder)
	migration()
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id is " + id)
}

func migration () {
	mysqlConnection := db.ConnectMySql()
	mysqlConnection.AutoMigrate(&models.Product{})
	mysqlConnection.AutoMigrate(&models.Order{})
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