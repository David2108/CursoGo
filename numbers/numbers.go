package numbers

import "errors"

//Interfaz vacia
//interface{}
func Sum(a interface{}, b interface{}) (int, error){

	//devuelva el tipo interfaz.(type)
	switch a.(type){
	case string:
		return 0, errors.New("El valor A es un string")
	}

	switch b.(type){
	case string:
		return 0, errors.New("El valor B es un string")
	}

	//convertir a.(tipo)
	//deuvelve nulo nil
	return a.(int) + b.(int), nil
}