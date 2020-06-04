package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoarchivo()
	leoarchivo2()
	graboarchivo()
	graboarchivo2()
}

func leoarchivo() {
	archivo, err := ioutil.ReadFile("./lorem.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoarchivo2() {
	archivo, err := os.Open("./lorem.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()

			fmt.Printf("Linea > " + registro + "\n")
		}
	}

	archivo.Close()
}

func graboarchivo() {
	archivo, err := os.Create("./nuevoarchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}

	fmt.Fprintln(archivo, "Esta es una nueva linea")
	archivo.Close()
}

func graboarchivo2() {
	fileName := "./nuevoarchivo.txt"

	if Appen(fileName, "\nSEGUNDA LINEA") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Appen(archivo string, texto string) bool {
	archivo, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = archivo.WriteString(texto)

}
