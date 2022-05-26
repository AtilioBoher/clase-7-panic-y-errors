package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// range for the ID generation
const min = 10000
const max = 99999

//----------------------------------

type cliente struct {
	legajo     int
	nombYApell string
	dni        int
	telef      int
	domic      string // domicilio
}

func openFileRecoverFunc() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func openFile(fileName string) *os.File {
	myFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0660)
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

func addClient(c cliente) {
	f := openFile("eje2/customers.txt") // the correct way to type the path is like this "eje2/customers.txt"
	defer f.Close()

	if _, err := f.Write([]byte(strconv.Itoa(c.legajo) + "\n" +
		c.nombYApell + "\n" +
		strconv.Itoa(c.dni) + "\n" +
		strconv.Itoa(c.telef) + "\n" +
		c.domic + "\n\n")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // this is here to have a different auto generated ID each time we run the program

	id, err := generateIdNumber()
	if err != nil {
		panic(err)
	}
	fmt.Printf("se generó el siguiente legajo: %d \n", id)

	var cliente1 cliente
	cliente1.legajo, _ = generateIdNumber()
	cliente1.nombYApell = "un nombre y apellido"
	cliente1.dni = 40358487
	cliente1.telef = 15783828
	cliente1.domic = "calle tanto"
	addClient(cliente1)
}
