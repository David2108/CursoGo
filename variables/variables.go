package variables

import (
	"fmt"
	"CursoGo/Name"
)

const fullName string = "%s %s, David Vivanco"
const edad = 29
const expenses = 100

func Variables()  {
	hijos := 1
	var estudio = "Superior"
	var (
		escuela     = "Escuela Juan Genaro Jaramillo"
		colegio     = "Colegio Miguel de Santiago"
		universidad = "Escuela Politécnica Nacional"
	)

	firstName := name.GetName()
	sister, brother := getNamesSiblings()
	income := getIncome(540)
	f32, f64 := getPrecition()

	fmt.Printf("Hola %s, ingreso \n", firstName)
	fmt.Printf("Siblings: %s and %s \n", sister, brother)
	fmt.Printf("Income: %f \n", income)
	fmt.Println("Hijos: ", hijos)
	fmt.Println("Estudios: ", estudio)
	fmt.Println("Escuela: ", escuela)
	fmt.Println("Colegio: ", colegio)
	fmt.Println("Universidad: ", universidad)
	fmt.Printf("Float32: %f, Float64: %f\n", f32, f64)
	fmt.Println("hola"[0])
	fmt.Printf(string("hola"[0]) + "\n")
	fmt.Println("Número de letras: ", len("Hola"))
}

func getNamesSiblings() (string, string) {
	return "Pamela", "Diego"
}

func getIncome(salary float64) float64 {
	return salary - expenses
}

func getPrecition() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

//Array tamaño fijo, Slide tamaño variable
func GetArray() {
	var arr1 [2]string
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1)
}

func GetSlice() {
	var slice1 []string //No se pone el tamaño
	slice1 = append(slice1, "slice1", "slice2")
	fmt.Println(slice1)
}