package estructura

import "fmt"

type PlatziCourse struct{
	Name string
	Slug string
	Skills []string
	Courses []string
}

//Metodo de un struct
func (p PlatziCourse) Subscribe(name string){
	fmt.Printf("La persona %s se ha registrado al curso %s",  name, p.Name)
}

type PlatziCareer struct{
	PlatziCourse
}

//Metodo de un struct
func (p PlatziCareer) Subscribe(name string){
	fmt.Printf("La persona %s se ha registrado a la carrera %s",  name, p.Name)
}