package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

type safeMap struct {
  v map[string]string
}

func (m *safeMap) existKey(key string) bool {
  _, ok := m.v[key]
  return ok
 }

func (m *safeMap) get(key string) (string, error) {
  return m.v[key]
}

func (m *safeMap) set(key string, value string) {
  m.v[key] = value
}


// map作成
var id_message safeMap

func init() {
	id_message.set("123456" , "ジョンソン")
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
  if id_message.existKey(id) {
    result = id_message.get(id)
    return c.JSON(http.StatusOK, map[string]string{"message": result})
  }

  else {
    var apierr APIError
    apierr.Code    = 404
    apierr.Message = "Not Found"

    c.JSON(htt.StatusBadRequest, apierr)
    return err
  }
}
