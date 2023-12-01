package main

import "fmt"

func main() {
	persona1 := Persona{
		nombre:   "Tomas",
		edad:     25,
		ciudad:   "Corrientes",
		telefono: "3777123456",
	}

	persona2 := Persona{
		nombre:   "Manuel",
		edad:     28,
		ciudad:   "Buenos Aires",
		telefono: "221452388",
	}

	fmt.Println(persona1)
	fmt.Println(persona2)

	persona2.cambiarCiudad("Cordoba")
	fmt.Println(persona1)
	fmt.Println(persona2)
}

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono string
}

func (p *Persona) cambiarCiudad(nuevaCiudad string) {
	if(p.ciudad == nuevaCiudad) {
		return
	}

	p.ciudad = nuevaCiudad
	fmt.Println("Ciudad actualizada con exito.")
}
