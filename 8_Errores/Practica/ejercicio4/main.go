package main

import (
	"fmt"
	"os"
)

/*
# Ejercicio 4: Manejo de errores en operaciones de archivos

Escribe un programa que intente abrir un archivo cuyo nombre y directorio se pasa como argumento en la línea de comandos. Si el archivo no existe, el programa debe crearlo. Si hay algún otro error al abrir el archivo, el programa debe informar del error. No olvides cerrar el archivo correctamente en todos los casos.
*/

func main() {
	var dirName string
	var fileName string

	fmt.Println("Enter a directory and file name to open a file or create one.")
	fmt.Scan(&dirName)
	fmt.Scan(&fileName)

	if file, err := openFile(dirName, fileName); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("File '%v' open.\n", file.Name())
	}

}

func openFile(dir, name string) (*os.File,error) {

	path := dir + "/" + name

	file, err := os.OpenFile(path, os.O_CREATE, 0755)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return file, nil
}

