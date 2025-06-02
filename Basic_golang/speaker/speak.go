package speaker

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p Person) Speak() string {
	return fmt.Sprintf("%s says hello!", p.Name)
}

func (p Person) Hand() string {
	return fmt.Sprintf("%s waves hand!", p.Name)
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s barks!", d.Name)
}

func Speak(s Speaker) {
	fmt.Println(s.Speak())
}