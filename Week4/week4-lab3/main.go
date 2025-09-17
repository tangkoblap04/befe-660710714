package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID string ` json:"id" `
	Name string ` json:"name" `
	Email string ` json:"email" `
	Year int ` json:"year" `
	GPA float64 ` json:"gpa" `
}

func (s * Student) IsHonor() bool{
	return s.GPA >= 3.50

}

func (s * Student) Validate() error {
	
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1 and 4")
	}
	if s.GPA < 0.0 || s.GPA > 4.0 {
		return errors.New("GPA must be between 0.0 and 4.0")
	}
	return nil
}

func main(){
	// var st Student = Student{
	// 	ID:"1", 
	// 	Name:"", 
	// 	Email:"tangkoblap_t@silpakorn.edu", 
	// 	Year:3, 
	// 	GPA:3.63,
	// }

	// st := Student{
	// 	ID: "1",
	// 	Name: "Thanin",
	// 	Email: "tangkoblap_t@silpakorn.edu",
	// 	Year: 3,
	// 	GPA: 3.63
	// }

	students := []Student{
		{ID: "1", Name: "Thanin", Email: "tangkoblap_t@silpakorn.edu", Year: 3, GPA: 3.63},
		{ID: "2", Name: "John Doe", Email: "johndoe@example.com", Year: 2, GPA: 3.40},
		
	}

	newStudent := Student{
		ID: "3",
		Name: "Jane Smith",
		Email: "janesmith@example.com",
		Year: 1,
		GPA: 3.80,
	}

	students = append(students, newStudent)

	for i, st := range students {
		fmt.Printf("Index : %d \nHonor : %v\n",i+1, st.IsHonor())
		fmt.Printf("Validation : %v\n", st.Validate())
		fmt.Println("___________________________")
	}
}