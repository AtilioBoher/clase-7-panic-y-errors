package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// constants and global variables
const min = 10000
const max = 99999

var didErrorOccurred bool = false

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

func checkIfClientExist(c cliente) bool {
	aux, err := os.ReadFile("eje2/customers.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := string(aux)
	// fmt.Println(data)
	condition1 := strings.Contains(data, "Nombre y Apellido: "+c.nombYApell)
	condition2 := strings.Contains(data, "DNI: "+strconv.Itoa(c.dni))
	return condition1 && condition2 // name and DNI has to be the same. The rest could be checked, for example the ID, but I consider is not necessary
}

func isClientNull(c cliente) error {
	switch {
	case c.legajo == 0:
		return errors.New("Error: no se ha ingresado el legajo, volver a intentar ingresando un valor")
	case c.nombYApell == "":
		return errors.New("Error: no se ha ingresado el nombre y apellido, volver a intentar ingresando un valor")
	case c.dni == 0:
		return errors.New("Error: no se ha ingresado el DNI, volver a intentar ingresando un valor")
	case c.telef == 0:
		return errors.New("Error: no se ha ingresado el telefono, volver a intentar ingresando un valor")
	case c.domic == "":
		return errors.New("Error: no se ha ingresado el domicilio, volver a intentar ingresando un valor")
	default:
		return nil
	}
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

func showAllClients() {
	aux, err := os.ReadFile("eje2/customers.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := string(aux)
	fmt.Println(data)
}

func lastMessage() {
	if didErrorOccurred == true {
		fmt.Println("Se detectaron varios errores en tiempos de ejecución")
	} else {
		fmt.Println("fin de la ejecución de manera correcta")
	}
}

func readInt() (num int) {
	text := readString()
	num, _ = strconv.Atoi(text)
	return num
}

func readString() (text string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	return text[:len(text)-2]
}

func main() {
	rand.Seed(time.Now().UnixNano()) // this is here to have a different auto generated ID each time we run the program

	defer lastMessage()

	var (
		cliente1 cliente
		err      error
	)

	// ------------este sería el procedimiento para cargar un cliente nuevo--------------
	fmt.Println("----------------------------")
	// el legajo se genera automáticamente
	cliente1.legajo, err = generateIdNumber()
	if err != nil {
		panic(err)
	}
	err = errors.New("no nil error")
	exist := true
	for err != nil || exist == true {
		fmt.Println("Ingrese nombre y apellido: ")
		cliente1.nombYApell = readString()
		fmt.Println("Ingrese DNI: ")
		cliente1.dni = readInt()
		fmt.Println("Ingrese número de telefono: ")
		cliente1.telef = readInt()
		fmt.Println("Ingrese el domicilio: ")
		cliente1.domic = readString()
		err = isClientNull(cliente1)
		if err != nil {
			fmt.Println(err)
			didErrorOccurred = true
		} else {
			exist = checkIfClientExist(cliente1)
			if exist == false {
				addClient(cliente1)
			} else {
				fmt.Println("Error: El cliente ya existe")
				didErrorOccurred = true
			}
		}
	}

	fmt.Println("Ingrese \"s\" si desea ver la lista de clientes: ")
	aux := readString()
	if aux == "s" {
		showAllClients()
	}

}
