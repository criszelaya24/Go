package main

import (
"github.com/criszelaya24/Go/db"
  "github.com/criszelaya24/Go/handlers"
  "github.com/gin-gonic/gin"
  "fmt"
  "github.com/joho/godotenv"
  _ "github.com/joho/godotenv"
  "github.com/kelseyhightower/envconfig"
)

// seeting an struct to load the .env value
type DBconfig struct {
  DbHost    string `envconfig:"DB_HOST"`
  DbUser    string `envconfig:"DB_USER"`
  DbPass    string `envconfig:"DB_PASS"`
  DbPort    string `envconfig:"DB_PORT"`
  DbName    string `envconfig:"DB_NAME"`
  GinMode   string `envconfig:"GIN_MODE"`
  // GinPort   string `envconfig:"PORT"`
  // AppSecret string `envconfig:"APP_SECRET"`
}

var ac DBconfig // declaring to access to the variable inside of the struct

// Getting configuration from .env file and environment variables
func LoadEnv() {

  //Loading configuration based on .env file
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file: %s", err)
  }

  //Setting up struct
  err = envconfig.Process("", &ac)
  if err != nil {
    fmt.Println("Error setting up AppConfig struct: %s", err)
  }
}

func main() {
  // connect with db
  LoadEnv()
  if err := postgres.ConnectDB(ac.DbUser, ac.DbPass, ac.DbHost, ac.DbPort, ac.DbName); err != nil {
    fmt.Println("Error", err)
    fmt.Println(ac.DbHost)
  }
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  // Loading files
  router.LoadHTMLGlob("views/*")

  // index page default
  router.GET("/", handler.Index)

   // Setup route group for the API
  user := router.Group("/user")
  {
    user.GET("/signup", handler.Signup)
    user.POST("/signup", handler.SignupCreate)
 
  }
  // Start and run the server
  router.Run(":3000")
}


