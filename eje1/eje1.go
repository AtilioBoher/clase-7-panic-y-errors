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

func openFile(fileName string) {
	_, err := os.Open(fileName)
	defer fmt.Println("Ejecución finalizada")
	defer openFileRecoverFunc()
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}

func main() {
	openFile("customers.txt")
	fmt.Println("la ejecución del programa sigue")
}
