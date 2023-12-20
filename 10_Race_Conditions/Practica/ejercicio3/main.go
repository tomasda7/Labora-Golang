package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Desarrolla un programa en Golang con un canal (de enteros) compartido entre dos hilos “productores” y un hilo “consumidor”. La regla es que los hilos productores deben colocar valores en el canal de manera intercalada, asegurándose de que un hilo productor no ponga dos valores consecutivos sin que el otro hilo productor haya puesto uno.
*/

func main() {
	var channel = make(chan int)

	var prod1, prod2 sync.WaitGroup
	var cons sync.WaitGroup

	var mutex sync.Mutex

	cons.Add(3)

	prod1.Add(1)
	prod2.Add(1)

	var mainRunning bool = true

	go func() {
		defer cons.Done()

		for mainRunning {
			prod1.Wait()

			mutex.Lock()
			if mainRunning {
				channel <- 0
			}
			mutex.Unlock()

			prod1.Add(1)
			//fmt.Println("thread1 added 1 to prod1")
			prod2.Done()
			//fmt.Println("thread1 subtracts 1 to prod2")
		}
	}()

   go func() {
		defer cons.Done()

		for mainRunning {
			prod2.Wait()

			mutex.Lock()
			if mainRunning {
				channel <- 1
			}
			mutex.Unlock()

			prod2.Add(1)
			//fmt.Println("thread2 added 1 to prod2")
			prod1.Done()
			//fmt.Println("thread2 subtracts 1 to prod1")
		}
   }()

   go func() {
		defer cons.Done()

		for v := range channel {
			fmt.Println(v)
			time.Sleep(100 * time.Millisecond)
		}
   }()

   prod2.Done()
   //fmt.Println("main subtracts 1 to prod2")
   time.Sleep(1 * time.Second)

   mutex.Lock()
   mainRunning = false
   close(channel)
   mutex.Unlock()

   cons.Wait()
}
