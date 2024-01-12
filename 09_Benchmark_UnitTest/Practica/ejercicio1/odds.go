package odds

/*
Escribir un algoritmo para sumar los primeros “n” números impares, siendo un “n” un parámetro en un archivo llamado “odds.go”. Luego se desea poner a prueba la correctitud escribiendo en un archivo “odds_test.go”,se pide usar el enfoque de table driven test.
*/

func sumFirstOdds(n int) int {
	var acc int
	odd := 1

	for i := 0; i < n; i++ {
		acc += odd
		odd += 2
	}

	return acc
}
