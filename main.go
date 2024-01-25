package main

import (
  "net/http"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load("./.env")
  if err != nil {
    panic("Error loading .env file")
  }

  r := gin.Default()
  r.GET("/test", func(c *gin.Context) {
    c.JSON(http.StatusOK, "Ready to dev PetPal App !!!")
  })

  port := os.Getenv("PORT")
  r.Run(":" + port)
}
