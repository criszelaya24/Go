package main
import (
  "net/http"
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)
func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  // Serve frontend static files
  router.Use(static.Serve("/", static.LocalFile("./views", true)))
  // Setup route group for the API
  // router.GET("/", func(c *gin.Context) {
  //   c.String(200, "GOORDA")
  // })
  api := router.Group("/api")
  {
    api.GET("/", Index)
    api.GET("/jokes", JokeHandler)
    api.POST("/jokes/like/:jokeID", LikeJoke)
  }
  // Start and run the server
  router.Run(":3000")
}

// index function
func Index(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message": "Welcome",
    })
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"Jokes handler not implemented yet",
  })
}
// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"LikeJoke handler not implemented yet",
  })
}