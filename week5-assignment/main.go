package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Employee struct
type Employee struct {
    employee_ID       string  `json:"id"`
    first_name     string  `json:"first_name"`
	last_name	  string  `json:"last_name"`
    Email    string  `json:"email"`
    salary     int     `json:"salary"`
    department_id     float64 `json:"department_id"`
}

// In-memory database (In real project, we use database)
var employees = []Employee{
    {employee_ID: "1", first_name: "John", last_name: "Doe", Email: "john@example.com", salary: 50000, department_id: 1},
    {employee_ID: "2", first_name: "Jane", last_name: "Smith", Email: "jane@example.com", salary: 60000, department_id: 2},
	{employee_ID: "3", first_name: "Alice", last_name: "Johnson", Email: "alice@example.com", salary: 70000, department_id: 1},
}

func getHealth(c * gin.Context){
	c.JSON(200, gin.H{
		"message" : "healthy",
	})
}

func getEmployees(c * gin.Context){
	//This function is call "Hander function"
	salaryQuery := c.Query("salary")

	if salaryQuery != "" {
		filter := []Employee{}
		for _,employee := range employees{
			if fmt.Sprint(employee.salary) == salaryQuery{
				filter = append(filter, employee)
			}

		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, employees)

}
func main(){
	r := gin.Default()
	r.GET("/health",getHealth) //This API path for check only health.

	api := r.Group("/api/v1")
	api.GET("/employees", getEmployees)
	 //Group for remine, This is api

	r.Run(":8080") 
}