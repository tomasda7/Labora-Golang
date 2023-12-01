package main

import (
	"fmt"
	"sort"
)

func main() {

	var personas = []Persona{{"tomas", 24, 1.80, 90}, {"pedro", 22, 1.50, 70}, {"raul", 21, 1.60, 75.5} }
	perona1 := personas[0]
	//fmt.Println(ordenarPersonas(personas, 0))
	//fmt.Println(buscarPersona(personas, "22"))
	fmt.Println(perona1.calcuclarIMC())
}

type Persona struct {
	nombre string
	edad   int
	altura float64
	peso   float64
}

func ordenarPersonas(personas []Persona, criterio int) []Persona {

	switch criterio {
	case 1:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].nombre < personas[j].nombre
		})
	case 2:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].edad < personas[j].edad
		})
	case 3:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].altura < personas[j].altura
		})
	case 4:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].peso < personas[j].peso
		})
	default:
		return personas
	}

	return personas
}

func buscarPersona(personas []Persona, valor string) *Persona {

	for _, persona := range personas {
		if valor == persona.nombre || valor == fmt.Sprint(persona.edad) ||
			valor == fmt.Sprint(persona.altura) || valor == fmt.Sprint(persona.peso) {
			return &persona
		}
	}

	return nil
}

func (persona Persona) calcuclarIMC() float64 {
	imc := persona.peso / (persona.altura * persona.altura)

	return imc
}

func (persona Persona) clasificarIMC() string {
	imc := persona.calcuclarIMC()

	switch {
	case imc < 18.5:
		return "Bajo peso: IMC menor a 18.5"
	case imc >= 18.5 && imc < 25:
		return "Peso normal: IMC entre 18.5 y 24.9"
	case imc >= 25 && imc < 30:
		return "Sobrepeso: IMC entre 25 y 29.9"
	case imc > 30:
		return "Obesidad: IMC mayor a 30"
	default:
		return "IMC fuera de rango"
	}
}




