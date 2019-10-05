package main

import "fmt"

const fullName string = "%s %s, David Vivanco"
const edad = 29

func main() {
	var name string
	hijos := 1
	var estudio = "Superior"
	var (
		escuela     = "Escuela Juan Genaro Jaramillo"
		colegio     = "Colegio Miguel de Santiago"
		universidad = "Escuela Politécnica Nacional"
	)
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hola %s, ingreso.", name)
	fmt.Println("Hijos: ", hijos)
	fmt.Println("Estudios: ", estudio)
	fmt.Println("Escuela: ", escuela)
	fmt.Println("Colegio: ", colegio)
	fmt.Println("Universidad: ", universidad)

}
