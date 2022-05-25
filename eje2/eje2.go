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

func generateIdNumber() (int, error) {
	randNum := rand.Intn(max-min) + min
	// randNum = 3 // uncomment this line to test the panic
	if randNum < min || randNum > max {
		return 0, fmt.Errorf("el número de legajo generado es %d, el cual se encuentra fuera del rango permitido (%d-%d)", randNum, min, max)
	}
	return randNum, nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // this is here to have a different auto generated ID each time we run the program

	id, err := generateIdNumber()
	if err != nil {
		panic(err)
	}
	fmt.Printf("se generó el siguiente legajo: %d", id)

	myFile := openFile("eje2/customers.txt") // the correct way to type the path is like this "eje2/customers.txt"
	defer myFile.Close()
	fmt.Println("la ejecución del programa sigue")

}
