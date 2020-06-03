package main

import (
	"fmt"
)

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	Escarnivoro() bool
}

type vegetal interface {
	Tipovegetal() string
}

/*Humano*/

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

/*mujer*/

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	}
	return "Mujer"
}

func humanosrespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s y estoy respirando\n", hu.sexo())
}

/* Gnero anumal */

/*perro*/

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) Escarnivoro() bool { return p.carnivoro }

func animalesrespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func animalescarnivoro(an animal) int {
	if an.Escarnivoro() == true {
		return 1
	}
	return 0

}

func main() {
	tcarrnivoro := 0

	Pedro := new(hombre)
	Pedro.esHombre = true
	humanosrespirando(Pedro)

	Maria := new(mujer)
	humanosrespirando(Maria)

	Drogo := new(perro)
	Drogo.carnivoro = true
	animalesrespirar(Drogo)
	tcarrnivoro = +animalescarnivoro(Drogo)

	fmt.Printf("Carnivoros %d\n", tcarrnivoro)

}
