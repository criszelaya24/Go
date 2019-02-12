package function

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

// index function

func Index(c *gin.Context) {
  c.HTML(200, "index.html", gin.H {
      "data": "Sending information",
      })
}
func IndexApi(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message": "Welcome",
    })
}

// JokeHandler retrieves a list of available jokes
func GetBookmars(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"Bookmark handler not implemented yet",
  })
}
