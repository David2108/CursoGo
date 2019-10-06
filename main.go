package main

import (
	"fmt"
)

const fullName string = "%s %s, David Vivanco"
const edad = 29
const expenses = 100

func main() {
	//variables()
	getArray()
	getSlice()
	//ifTest()
	forTest()

}

func variables()  {
	hijos := 1
	var estudio = "Superior"
	var (
		escuela     = "Escuela Juan Genaro Jaramillo"
		colegio     = "Colegio Miguel de Santiago"
		universidad = "Escuela Politécnica Nacional"
	)

	name := getName()
	sister, brother := getNamesSiblings()
	income := getIncome(540)
	f32, f64 := getPrecition()

	fmt.Printf("Hola %s, ingreso \n", name)
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

func getName() string {
	var name string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanf("%s", &name)
	return name
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
func getArray() {
	var arr1 [2]string
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1)
}

func getSlice() {
	var slice1 []string //No se pone el tamaño
	slice1 = append(slice1, "slice1", "slice2")
	fmt.Println(slice1)
}

func ifTest() {
	var old int

	fmt.Print("Ingrese un número entero: ")
	fmt.Scan(&old)

	if old%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if number := 2; number == 2 {
		fmt.Println("Second kind for declarate statement if")
	}

}

func forTest(){
	for i:=0; i<5; i++{
		fmt.Println("[FOR] ",i)
	}

	c := 100
	for ; c>0 ;{
		c -= 10
		fmt.Println("[FOR] Solo con una condicion ", c)
	}

	s := 1000
	for {
		s -= 1
		if s==0{
			fmt.Println("[FOR] infinito")
			break
		}
	}
}
