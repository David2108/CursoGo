package estructurasControl

import "fmt"

func IfTest() {
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

func ForTest(){
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

func SwitchTest(){
	var n int
	fmt.Println("Ingrese un número: ")
	fmt.Scan(&n)

	switch n{
	case 1: fmt.Println("El número es 1")
	default: fmt.Println("El número no es 1")
	}

	switch {
	case n%2==0: fmt.Println("El número es par")
	default: fmt.Println("El número no es par")
	}
}