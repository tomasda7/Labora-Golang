package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds int

	fmt.Println("Ingrese una cantidad de segundos")
	fmt.Scan(&seconds)

	duration := time.Duration(seconds) * time.Second

	totalHours := int(duration.Hours())
	days :=  totalHours / 24
	hours := totalHours % 24

	totalMinuts := int(duration.Minutes())
	minuts := totalMinuts % 60

	totalSeconds := int(duration.Seconds())
	finalSeconds := totalSeconds % 60

	if(days == 0 && hours > 0) {
		fmt.Printf("%d segundos son: %d horas, %d minutos con %d segundos", seconds, hours, minuts, finalSeconds)

	} else if(days == 0 && hours == 0) {
		fmt.Printf("%d segundos son: %d minutos con %d segundos", seconds, minuts, finalSeconds)
	} else {
		fmt.Printf("%d segundos son: %d dias, %d horas, %d minutos con %d segundos", seconds, days, hours, minuts, finalSeconds)
	}
}
