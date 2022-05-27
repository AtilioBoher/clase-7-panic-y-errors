package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func generateIdNumber() (int, error) {
	band := true
	var randNum int
	for band {
		randNum = rand.Intn(max-min) + min
		// randNum = 3 // uncomment this line to test the panic
		if randNum < min || randNum > max {
			return 0, fmt.Errorf("el número de legajo generado es %d, el cual se encuentra fuera del rango permitido (%d-%d)", randNum, min, max)
		}
		band = checkIfIdNumberExist(randNum) // repeats until the genereted number is different from the already existing ones
	}
	return randNum, nil
}

func checkIfIdNumberExist(id int) bool {
	aux, err := os.ReadFile("eje2/customers.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := string(aux)
	return strings.Contains(data, "Legajo: "+strconv.Itoa(id))
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

func checkIfClientExist(c cliente) bool {
	aux, err := os.ReadFile("eje2/customers.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := string(aux)
	// fmt.Println(data)
	condition1 := strings.Contains(data, "Nombre y Apellido: "+c.nombYApell)
	condition2 := strings.Contains(data, "DNI: "+strconv.Itoa(c.dni))
	return condition1 && condition2 // name and DNI has to be the same, the rest could be checked, but I consider is not necessary
}

func addClient(c cliente) {
	f := openFile("eje2/customers.txt") // the correct way to type the path is like this "eje2/customers.txt"
	defer f.Close()

	if f != nil {

		if _, err := f.Write([]byte("Legajo: " + strconv.Itoa(c.legajo) + "\n" +
			"Nombre y Apellido: " + c.nombYApell + "\n" +
			"DNI: " + strconv.Itoa(c.dni) + "\n" +
			"Telefono: " + strconv.Itoa(c.telef) + "\n" +
			"Domicilio: " + c.domic + "\n\n")); err != nil {
			fmt.Println(err)
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano()) // this is here to have a different auto generated ID each time we run the program

	var (
		cliente1 cliente
		err      error
	)

	cliente1.legajo, err = generateIdNumber()
	if err != nil {
		panic(err)
	}
	fmt.Printf("se generó el siguiente legajo: %d \n", cliente1.legajo)
	cliente1.nombYApell = "un nombre y apellido1"
	cliente1.dni = 40358487
	cliente1.telef = 15783828
	cliente1.domic = "calle tanto"
	// addClient(cliente1)
	fmt.Println("----------------------------")
	fmt.Println(checkIfClientExist(cliente1))
}
