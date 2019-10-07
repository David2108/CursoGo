package main

import (
	"fmt"
	"math/rand"
)

func main() {

	end := 0
	for end == 0 {
		var numberMachice int
		var numberPerson int
		numberMachice = getNumberMachine()

		fmt.Printf("Ingrese un número: ")
		fmt.Scanf("%d",&numberPerson)

		if numberMachice == numberPerson {
			fmt.Println("You WIN!!.")
		}

		end = askQuestion()

	}

}

func getNumberMachine() int {
	return rand.Intn(10)
}

func askQuestion() int{
	end := 0
	var resp string
	fmt.Printf("Do you want play again? y/n ")
	fmt.Scanf("%s", &resp)
	if resp == "N" || resp == "n" {
		end = 1
		fmt.Println("GAME OVER ('-') (´-´) ");
	}
	return end
}


