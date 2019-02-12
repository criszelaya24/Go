package main

import (
  test "github.com/criszelaya24/Go/functions"
  "github.com/gin-gonic/gin"
)
func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  // Loading files
  router.LoadHTMLGlob("views/*")

  // index page default
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H {
      "data": "Sending information",
      })
  })

   // Setup route group for the API
  api := router.Group("/api")
  {
    api.GET("/", test.Index)
    api.GET("/jokes", test.JokeHandler)
    api.POST("/jokes/like/:jokeID", test.LikeJoke)
  }
  // Start and run the server
  router.Run(":3000")
}

