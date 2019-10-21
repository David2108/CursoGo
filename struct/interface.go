package estructura

import "fmt"

type Platzi interface{
	Subscribe(name string)
}

func deferTest(){
	fmt.Println("La funcion Interface ha terminado")
}

func InterfaceTest(){
	defer deferTest() //Ejecuta al final
	platziCourse := PlatziCourse{Name: "Go", Slug: "g", Skills: []string{"1", "2"}, Courses: []string{"Go"}} 

	platziCareer := new(PlatziCareer)
	platziCareer.Name = "GoCareer"
	platziCareer.Slug = "go"

	callSubscribe(platziCareer)
	callSubscribe(platziCourse)
}

func callSubscribe(p Platzi){
	p.Subscribe("Sebastian")
}
