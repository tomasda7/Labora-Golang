package main

import (
	"fmt"
	"sync"
)

/*
Se desea implementar un tipo de datos abstracto (puede ser una struct) que represente un mapa sincronizado, es decir, que tenga accesos concurrentes seguros o que permite operaciones concurrentes en forma segura (thread safe).
A) ¿Es importante que sea concurrentemente seguro a nivel solo lectura?
-No hay "race conditions" cuando el recurso al que se accede no se modifica.
B) ¿Qué diferencia observas si usas métodos mutadores que sean pointer receiver y value receiver? Los métodos mutadores son los que realizan cambios en el estado.
- Al usar una propiedad de tipo mutex en la struct "ThreadSafe" se debe usar punteros.
*/

func main() {
	var wg sync.WaitGroup
	myMap := ThreadSafe{safeMap: make(map[int]int)}

	wg.Add(4)
	go func() {
		defer wg.Done()
		myMap.getOrCreate(1, 1)
	}()
	go func() {
		defer wg.Done()
		myMap.getOrCreate(1, 2)
	}()
	go func() {
		defer wg.Done()
		myMap.getOrCreate(1, 3)
	}()
	go func() {
		defer wg.Done()
		myMap.getOrCreate(1, 4)
	}()


	wg.Wait()

	fmt.Printf("My map:%+v\n", myMap.safeMap)
}

type ThreadSafe struct {
	safeMap map[int]int
	mutex sync.Mutex
}

func (ts *ThreadSafe) getOrCreate(key int, value int) int {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	_, exists := ts.safeMap[key]
	if !exists {
		ts.safeMap[key] = value
	}

	return ts.safeMap[key]
}
