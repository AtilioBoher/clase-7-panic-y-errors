package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// range for the ID generation
const min = 10000
const max = 99999

//----------------------------------

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

func generateIdNumber() int {
	randNum := rand.Intn(max-min) + min
	// randNum = 3  // uncomment this line to test the panic
	if randNum < min || randNum > max {
		panic("Error: el de legajo generado se encuentra fuera del rango definido")
	}
	return randNum
}

func main() {
	rand.Seed(time.Now().UnixNano())    // this is here to have a diferent auto generated ID each time we run the program
	myFile := openFile("customers.txt") // the correct way to type the path is like this "eje1/customers.txt"
	defer myFile.Close()
	fmt.Println("la ejecución del programa sigue")
	fmt.Println(generateIdNumber())
}
