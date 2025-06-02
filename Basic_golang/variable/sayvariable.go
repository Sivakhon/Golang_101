package variable

import (
	"fmt"
)

type Student struct {
	Name string	
	LastName string
	Age  int
	Grade string
}

func SayVariable() {
	fullName := "Sivakhon"
	age := 30
	fmt.Println("Hello, it defalut variable fucntion!")
	fmt.Printf("Enter your full name: %s\n", fullName)
	fmt.Printf("Enter your age: %d\n", age)

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	for i:= 0; i < len(myArray); i++ {
		fmt.Printf("myArray[%d] = %d\n", i, myArray[i])
	}

	mySlice := []int {}

	mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, "2")
	for i:= 0; i < len(mySlice); i++ {
		fmt.Printf("mySlice[%d] = %d\n", i, mySlice[i])
	}

	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	for key, value := range myMap {
		fmt.Printf("myMap[%s] = %d\n", key, value)
	}

	value, ok := myMap["one"]
	if ok {
		fmt.Printf("myMap[one] = %d\n", value)
	} else {
		fmt.Println("Key not found")
	}

	var student []Student

	student = append(student, Student{Name: "Sivakhon", Age: 30, Grade: "A"})
	student = append(student, Student{Name: "John", Age: 40, Grade: "B"})

	for i := 0; i < len(student); i++ {
		fmt.Printf("Student[%d] = %s, %d, %s\n", i, student[i].Name, student[i].Age, student[i].Grade)
	}

	fmt.Printf("capacity pf student: %d\n", cap(student))

}

func (s Student) Fullname() string{
		fullname := s.Name + " " + s.LastName
		return fullname
	}