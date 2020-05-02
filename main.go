package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuseiasumo/tanbun-api/types"
	"math/rand"
	"net/http"
	"time"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://quicker.netlify.app/"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	rand.Seed(time.Now().UnixNano())

	idMessage.Init()

	// Routes
	e.GET("/", getid)
	e.POST("/", postmessage)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

func generateUniqueId() string {
	var idtmp int
	var idstr string

	for {
		idtmp = rand.Intn(1000000)
		idstr = fmt.Sprintf("%06d", idtmp)
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

func deleteMessage(key string) {
	time.Sleep(time.Minute * 10)
	idMessage.RemoveByKey(key)
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

	if message.Message == "" {
		var apierr types.APIError
		apierr.Code = 400
		apierr.Message = "Bad Request"

		return c.JSON(http.StatusBadRequest, apierr)
	}

	id := generateUniqueId()

	idMessage.Set(id, message.Message)

	go deleteMessage(id)

	return c.JSON(http.StatusOK, map[string]string{"id": id})

}
