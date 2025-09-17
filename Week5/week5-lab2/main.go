package main

import (
	"github.com/gin-gonic/gin"
)
	


type User struct {
	ID string `json : "id"`
	Name string `json : "name"`
}

func getUser(c *gin.Context){
	user := []User{{ID : "1", Name : "Thanin"}}
	c.JSON(200, user)
}

func main(){
	r := gin.Default()
	r.GET("/users", getUser)

	r.Run(":8080")
}