package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Hacer una suma de los primeros 100 numeros, de forma secuencial y luego de forma concurrente y apreciar la diferencia en el tiempo de ejecución.
*/

func main() {
	// Suma secuencial
	startTime := time.Now()
	sumaSecuencial := sumaSecuencial(1, 1000000)
	elapsedTimeSecuencial := time.Since(startTime)

	fmt.Printf("Suma secuencial: %d\n", sumaSecuencial)
	fmt.Printf("Tiempo de ejecución secuencial: %s\n", elapsedTimeSecuencial)

	// Suma concurrente
	startTime = time.Now()
	sumaConcurrente := sumaConcurrente(1, 1000000)
	elapsedTimeConcurrente := time.Since(startTime)

	fmt.Printf("Suma concurrente: %d\n", sumaConcurrente)
	fmt.Printf("Tiempo de ejecución concurrente: %s\n", elapsedTimeConcurrente)
}

func sumaSecuencial(inicio, fin int) int {
	resultado := 0
	for i := inicio; i <= fin; i++ {
		resultado += i
	}
	return resultado
}

func sumaConcurrente(inicio, fin int) int {
	var wg sync.WaitGroup
	resultado := 0
	canal := make(chan int)

	// Dividir el rango en dos partes y sumar cada parte concurrentemente
	medio := (inicio + fin) / 2

	wg.Add(2)
	go func() {
		defer wg.Done()
		canal <- sumaSecuencial(inicio, medio)
	}()

	go func() {
		defer wg.Done()
		canal <- sumaSecuencial(medio+1, fin)
	}()

	// Cerrar el canal después de que ambas goroutines hayan terminado
	go func() {
		wg.Wait()
		close(canal)
	}()

	// Recoger los resultados de las goroutines y sumarlos
	for valor := range canal {
		resultado += valor
	}

	return resultado
}
