package main

import (
	"fmt"
)

type servivo interface {
	estavivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estavivo() bool
}

type animal interface {
	respirar()
	comer()
	Escarnivoro() bool
	estavivo() bool
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
	vivo       bool
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
func (h *hombre) estavivo() bool { return h.vivo }

func humanosrespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s estoy respirando y estoy vivo %t\n", hu.sexo(), hu.estavivo())
}

/* Gnero anumal */

/*perro*/

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) Escarnivoro() bool { return p.carnivoro }
func (p *perro) estavivo() bool    { return p.vivo }

func animalesrespirar(an animal) {
	an.respirar()
}

func animalescarnivoro(an animal) int {
	if an.Escarnivoro() == true {
		return 1
	}
	return 0

}

/* ser vivio */

func estoyvivo(v servivo) bool {
	return v.estavivo()
}

func main() {
	tcarrnivoro := 0

	Pedro := new(hombre)
	Pedro.vivo = true
	Pedro.esHombre = true
	humanosrespirando(Pedro)

	Maria := new(mujer)
	Maria.vivo = false
	humanosrespirando(Maria)

	Drogo := new(perro)
	Drogo.carnivoro = true
	Drogo.vivo = true
	animalesrespirar(Drogo)
	tcarrnivoro = +animalescarnivoro(Drogo)

	fmt.Printf("Soy un animal estoy respirando y estoy vivo %t Carnivoros %d\n", estoyvivo(Drogo), tcarrnivoro)

}
