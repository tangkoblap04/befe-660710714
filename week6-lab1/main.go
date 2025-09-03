package main

import (
	"fmt"
	"net/http"
	"slices" //for remove data in array
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

func getStudent(c *gin.Context){
	id := c.Param("id")
	
	for _, student := range students{
		if student.ID == id{
			c.JSON(http.StatusOK, student)
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error" : "Student not found!!"})
} 

func createStudent(c *gin.Context){
	var newStudent Student

	//Change JSON file into Struct 
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	if newStudent.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Can't find this name."})
		return
	}
	if newStudent.Year < 1 || newStudent.Year > 4 {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Year must be 1 - 4"})
		return
	}
	newStudent.ID = fmt.Sprintf("%d", len(students)+1)
	students = append(students, newStudent)
	c.JSON(http.StatusOK, students)

}

func updateStudent(c *gin.Context){
	id := c.Param("id")
	var updateStudent Student

	if err := c.ShouldBindJSON(&updateStudent); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	
	for i, student := range students{
		if student.ID == id {
			updateStudent.ID = id
			students[i] = updateStudent
			c.JSON(http.StatusOK, updateStudent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error" : "Student not found na ja"})
}

func deleteStudent(c *gin.Context){
	id := c.Param("id")

	for i,student := range students{
		if student.ID == id{
			students = slices.Delete(students, i, i+1)
			c.JSON(http.StatusOK, gin.H{"Success" : "Delete completed."})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error" : "Student not found na ja"})
}
func main(){
	r := gin.Default()
	r.GET("/health",getHealth) //This API path for check only health.

	api := r.Group("/api/v1")
	api.GET("/students", getStudents)
	api.GET("/students/:id", getStudent)
	api.POST("/students/", createStudent)
	api.PUT("/students/:id", updateStudent)
	api.DELETE("/students/:id", deleteStudent)
	 //Group for remine, This is api

	r.Run(":8080") 
}