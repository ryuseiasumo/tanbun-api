package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)


// map作成
var id_message map[string]string

func init() {
	id_message = map[string]string{
		"123456":
	}
}

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", getid)
  //e.POST()

  // Start server
  e.Logger.Fatal(e.Start(":8081"))
}

// Handler
func getid(c echo.Context) error {
	// 6桁のidを取得した時
	id := c.QueryParam("id")

	return c.JSON(http.StatusOK, map[string]string{
        "id": id,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"hello": greetingto})
  }
