package maps

func GetMap() map[string]int {

	//Se asigna con la funcion make	
	mapTest := make(map[string]int) 
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2
	return mapTest
}

func GetMapAsignacion() map[string]int {

	mapTest := map[string]int{
		"Sebastian": 6,
		"David": 29,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 2

	delete(mapTest, "llave1")
	delete(mapTest, "llave2")
	return mapTest
}