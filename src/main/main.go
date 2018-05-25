package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func getCat(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\nand her type was: %s", catName, catType))
}

func main() {

	e := echo.New()

	e.GET("/", getIndex)
	e.GET("/cat", getCat)

	e.Start(":8000")
}
