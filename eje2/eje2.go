package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// el rango para el generador de IDs
const min = 10000
const max = 90000

//----------------------------------

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
		panic("Error: el archivo indicado no fue encontrado o está dañado")
	}
}

func generateIdNumber() int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano()) // esto para tener IDs distintas en cada ejecución
	openFile("customers.txt")
	fmt.Println("la ejecución del programa sigue")
	fmt.Println(generateIdNumber())
}
