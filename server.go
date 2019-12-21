package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/order", saveProduct)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id is " + id)
}

/**
curl -F "name=ハンバーグ" -F "price=1280" http://localhost:1323/order
 */
func saveProduct(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	if price == "" {
		return c.String(http.StatusBadRequest, "401")
	}
	return c.String(http.StatusOK, "name:" + name + " price:" + price)
}