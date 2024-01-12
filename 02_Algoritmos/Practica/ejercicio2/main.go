package main

func main() {
	//Detecte el “pequeño inconveniente” con el siguiente algoritmo (no vale correrlo, sino que tienen que pensarlo!

	var n int = 0
	for n != 1 || n != 0 {
		n++
	}

	//El inconvenientes es que el bucle for no cuenta con una caso base de corte o limite, de pasar la condicion la cual siempre es verdadera se crearia un loop infinito.
}
