package handler

import(
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/criszelaya24/Go/model"
)

// index function
func Signup(c *gin.Context) {
  c.HTML(200, "signup.html", nil)
}

func SignupCreate(c *gin.Context) {
	err := model.NewUser(c.PostForm("name"), c.PostForm("username"), c.PostForm("email"), c.PostForm("password")))
	if err != " " {
		fmt.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Exito"})
	// name := c.PostForm("name")
	// u_name := c.PostForm("username")
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	
}