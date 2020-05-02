package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuseiasumo/tanbun-api/types"
	"net/http"
	"fmt"
	"time"
	"math/rand"
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

	rand.Seed(time.Now().UnixNano())

	idMessage.Init()
	idMessage.Set("123456", "ジョンソン")

	// Routes
	e.GET("/", getid)
	e.POST("/", postmessage)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}


func generateUniqueId() string {
	var idtmp int
	var idstr string

	for {
		idtmp = rand.Intn(1000000)
		idstr = fmt.Sprintf("%06d",idtmp)
		if !(idMessage.ExistKey(idstr)) {
			return idstr
		}
	}
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


// MessageParam は /api/hello が受けとるJSONパラメータを定義します。
type MessageParam struct {
    Message string `json:"message"`
}

func postmessage(c echo.Context) error {
	// メッセージを取得した時
	message := new(MessageParam)

	if err := c.Bind(message); err != nil {
        return err
	}

	id := generateUniqueId()

	idMessage.Set(id, message.Message)

	return c.JSON(http.StatusOK, map[string]string{"id": id})

}
