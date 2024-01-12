package main

import "fmt"

func main() {
    var arrayTest = []int{10, 20, 30, 40, 50}
    fmt.Println(deleteValue(arrayTest, 2))
}

/*
2. Haga que reciba un arreglo y una posición y “borre” el valor que hay en dicha posición. Por “borrar” entiendase que no se quita el elemento sino se mueven todos los elementos que están a la derecha un pasito a la izquierda, rellenado con el valor por defecto el extremo derecho. O sea…
    1. Si tengo el arreglo int[5]{10,20,30,40,50} y “borro” el elemento de la posición 2 (recordar que se indexan desde 0) quedaría ⇒ int[5]={10,20,40,50,0}
*/

func deleteValue(array []int, index int) []int {

    for i := range array {
        if i >= index && i < len(array)-1 {
            array[i] = array[i+1]
        }

    }

    array[len(array)-1] = 0

    return array
}
