package main

import (
	"fmt"
	"os"
)

func openFileRecoverFunc() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func openFile(fileName string) *os.File {
	myFile, err := os.Open(fileName)
	defer fmt.Println("Ejecución finalizada")
	defer openFileRecoverFunc()
	if err != nil {
		panic("Error: el archivo indicado no fue encontrado o está dañado")
	}
	return myFile
}

func main() {
	myFile := openFile("eje1/customers.txt") // la forma correcta de escribir la dirección es "eje1/customers.txt"
	defer myFile.Close()
	fmt.Println("la ejecución del programa sigue")
}
