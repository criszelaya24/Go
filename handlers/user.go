package handler

import(
	// "net/http"
	"github.com/gin-gonic/gin"
)

// index function
func Signup(c *gin.Context) {
  c.HTML(200, "signup.html", nil)
}

func SignupCreate(c *gin.Context) {
	
}