package main


import (
	"app/controllers"
	"app/db"
	"app/models"
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
	e.POST("/order", controllers.SaveProduct)
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