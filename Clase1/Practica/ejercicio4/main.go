package main

import (
	"fmt"
	"time"
)

func main() {
	var days int

	fmt.Println("Ingrese el numero de dias")
	fmt.Scan(&days)

	duration := time.Duration(days) * 24 * time.Hour
	seconds := duration.Seconds()

	fmt.Printf("El numero de dias %d equivale a %.0f segundos", days, seconds)
}
