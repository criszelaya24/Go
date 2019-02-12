package main

import (
  function "github.com/criszelaya24/Go/functions"
  "github.com/gin-gonic/gin"
)
func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  // Loading files
  router.LoadHTMLGlob("views/*")

  // index page default
  router.GET("/", function.Index)

   // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", function.IndexApi)
    api.GET("/jokes", function.JokeHandler)
    api.POST("/jokes/like/:jokeID", function.LikeJoke)
  }
  // Start and run the server
  router.Run(":3000")
}

