package main

import (
	"fmt"
)

func main()  {
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Colombia"] = "Bogota"
	paises["Peru"] = "Lima"

	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Bulls": 87,
		"Lakers": 60,
		"Heat": 70,
		"Bucks": 47}

		fmt.Println(campeonato)

		campeonato["Ditroid"] = 68

		delete(campeonato, "Lakers")

		fmt.Println(campeonato)

		for equipo, puntaje := range campeonato{
			fmt.Printf("El equipo: %s, tiene un puntaje de: %d\n", equipo, puntaje)
		}

		puntaje, existe := campeonato["Pistons"]
		fmt.Printf("El puntaje capturado es %d, y el equipo existe %t\n", puntaje, existe)
}