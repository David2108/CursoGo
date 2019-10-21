package main


import (
	"fmt"
	"CursoGo/numbers"
)

func main() {
	/*variables.Variables()
	variables.GetArray()
	variables.GetSlice()*/
	
	/*estructurasControl.IfTest()
	estructurasControl.ForTest()
	estructurasControl.SwitchTest()*/

	/*fmt.Print(maps.GetMap())
	fmt.Print(maps.GetMapAsignacion())*/

	/*platziCourse := estructura.PlatziCourse{Name: "Go", Slug: "g", Skills: []string{"1", "2"}, Courses: []string{"Go"}} 

	platziCourse1 := new(estructura.PlatziCourse)
	platziCourse.Subscribe("David")

	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "go1"
	platziCourse1.Skills = []string{"backend1"}
	platziCourse1.Courses = []string{"Go"}

	fmt.Print(platziCourse)
	fmt.Print(platziCourse1)*/

	/*
	estructura.InterfaceTest()*/

	number, err := numbers.Sum(50, 50)

	if err != nil{
		panic(err)
	}

	fmt.Println(number)
}
