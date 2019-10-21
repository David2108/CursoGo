package name

import "fmt"

//Metodos deben empezar con mayuscula
func GetName() string {
	var name string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanf("%s", &name)
	return name
}