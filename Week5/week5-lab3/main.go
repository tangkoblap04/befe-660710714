package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student struct
type Student struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Email    string  `json:"email"`
    Year     int     `json:"year"`
    GPA      float64 `json:"gpa"`
}

// In-memory database (In real project, we use database)
var students = []Student{
    {ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
    {ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
}

func getHealth(c * gin.Context){
	c.JSON(200, gin.H{
		"message" : "healthy",
	})
}

func getStudents(c * gin.Context){
	//This function is call "Hander function"
	yearQuery := c.Query("year")

	if yearQuery != "" {
		filter := []Student{}
		for _,student := range students{
			if fmt.Sprint(student.Year) == yearQuery{
				filter = append(filter, student)
			}

		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, students)

}
func main(){
	r := gin.Default()
	r.GET("/health",getHealth) //This API path for check only health.

	api := r.Group("/api/v1")
	api.GET("/students", getStudents)
	 //Group for remine, This is api

	r.Run(":8080") 
}