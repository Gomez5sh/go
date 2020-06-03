package main

import(
	"fmt"
)

type humano interface {
	respirar()
	pensar()
	comer()
	sexo()
}

type animal interface {
	respirar()
	comer()
	carniboro() bool
}

type vegetal interface {
	Tipovegetal() string
}

/*Humano*/

type hombre struct {
	edad		int
	altura		float32
	peso		float32
	respirando	bool
	pensando	bool
	comiendo	bool
}

/*mujer*/

type mujer struct {
	edad		int
	altura		float32
	peso		float32
	respirando	bool
	pensando	bool
	comiendo	bool
}

func (h *hombre) respirar() { h.respirando = true } 
func (h *hombre) comer() { h.comiendo = true }
func (h *hombre) pensar() { h.pensando = true }
func (h *hombre) sexo() string { return "hombre" }

func (m *mujer) respirar() { m.respirando = true } 
func (m *mujer) comer() {  m.comiendo = true }
func (m *mujer) pensar() { m.pensando = true }
func (m *mujer) sexo() string { return "mujer" }


func humanosrespirando(hu humano)  {
	hu.respirar()
	fmt.Printf("Soy un/a %s y estoy respirando\n", hu.sexo())
}

func main()  {
	Pedro := new(hombre)
	humanosrespirando(Pedro)

	Maria := new(mujer)
	humanosrespirando(Maria)
}