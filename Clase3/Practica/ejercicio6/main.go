package main

import "fmt"

func unAlgoritmo(a int) int {
	//guarda en una variable el resultado de la multiplicacion de 'a' * 2
	quieroCunfundir := a * 2

	//itera mientras el valor de 'a' sea distinto a 0
	for a != 0 {
		//le resta 1 al valor de 'a' en cada iteracion
		a--
		//multiplica el valor actual de 'a' * 2 y suma el resultado a la variable en cada iteracion
		quieroCunfundir += a * 2
	}

	//una vez terminada la iteracion se retorna el resultado final guardado en la variable
	return quieroCunfundir

	/*
	Conclusion: la funcion esta calculando la suma de la serie aritmetica de numeros enteros desde 'a' hasta 1, donde cada termino se multiplica por 2,
	por lo que le cambiaria el nombre de 'unAlgoritmo(a int)' a 'sumaSerieAritmetica(num int)' y el nombre de la variable 'quieroConfundir' por 'restultado'.
	*/
}


func unaCosa(valor int) {
	//toma un entero como parametro, lo pasa dos veces a la funcion 'unAlgoritmo', multiplica sus resultados y muestra el resultado final.
	resultado := unAlgoritmo(valor) * unAlgoritmo(valor)
	fmt.Println("El resultado es:", resultado)

	/*
	Conclusion: esta funcion esta calculando el cuadrado de la suma aritmetica de un numero por lo que le cambiaria el nombre a 'cuadradoSumaAritmetica()'
	*/
}

func main() {
	unaCosa(2)
}
