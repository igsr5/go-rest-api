package config

import (
  "github.com/julienschmidt/httprouter"
  _"fmt"

  "post-api/app/handlers/posts"
)

var Route *httprouter.Router

func init() {
  Route = httprouter.New()
  Route.GET("/posts", posts.Index)
  Route.POST("/posts", posts.Create)
}
