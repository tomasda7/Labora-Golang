package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Ingrese un numero")
	fmt.Scan(&num)
    fmt.Println(ejercicio5(num))
}

func ejercicio5(x int) (int, int, int, int, int) {
    s1, s2, s3, s4, s5 := 0, 0, 0, 0, 0
    switch {
    case x >= 0 && x <= 50:
        s1 = x
        return s1, s2, s3, s4, s5
    case x > 50 && x <= 100:
        s1 = 50
        s2 = x - s1
        return s1, s2, s3, s4, s5
    case x > 100 && x <= 700:
        s1 = 50
        s2 = 50
        s3 = x - s1 - s2
        return s1, s2, s3, s4, s5
    case x > 700 && x <= 1500:
        s1 = 50
        s2 = 50
        s3 = 600
        s4 = x - s1 - s2 - s3
        return s1, s2, s3, s4, s5
    case x > 1500:
        s1 = 50
        s2 = 50
        s3 = 600
        s4 = 800
        s5 = x - s1 - s2 - s3 - s4
        return s1, s2, s3, s4, s5
    default:
        return s1, s2, s3, s4, s5
    }
}

/*
5. Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.
]
Sea X el número.
X = S1 + S2 + S3 + S4 + S5
Donde
0 ≤ S1 ≤ 50
0 ≤ S2 ≤ 50
0 ≤ S3 ≤ 600
0 ≤ S4 ≤ 800
0 ≤ S5 < (Infinito)
De esta forma, si X vale 120, entonces
S1 = 50, S2 = 50, S3 = 20, S4 = 0 y S5 = 0
Si X vale 860, entonces
S1 = 50, S2 = 50, S3 = 600, S4 = 160 y S5 = 0
*/
