package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuseiasumo/tanbun-api/types"
	"net/http"
)

// map作成
var (
	idMessage types.SafeMap
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	idMessage.Init()
	idMessage.Set("123456", "ジョンソン")

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
	if idMessage.ExistKey(id) {
		result := idMessage.Get(id)
		return c.JSON(http.StatusOK, map[string]string{"message": result})
	} else {
		var apierr types.APIError
		apierr.Code = 404
		apierr.Message = "Not Found"

		return c.JSON(http.StatusNotFound, apierr)
	}
}
