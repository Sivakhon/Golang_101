package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sivakhon/go-example/talk"
	"github.com/sivakhon/go-example/variable"
	"github.com/sivakhon/go-example/speaker"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello, World!")
	fmt.Println("Generated UUID:", id.String())

	talk.SayHello()
	talk.SayTest()
	variable.SayVariable()

	student := variable.Student{
		Name:     "Sivakhon",
		LastName: "Sivakhon",
		Age:      30,
		Grade:    "A",
	}

	fullname := student.Fullname()
	fmt.Printf("Student Fullname: %s\n", fullname)

	person := speaker.Person{Name: "John"}
	dog := speaker.Dog{Name: "Buddy"}
	speaker.Speak(person)
	speaker.Speak(dog)
	
}
