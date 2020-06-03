package main

import "fmt"
 
var matriz [5][7] int
func main()  {
	/*tabla := [10] int {0, 1, 2, 3, 4, 5, 6, 7, 8, 1000}
	for i := 0; i < len(tabla); i++{
		if (i < 8) {
		fmt.Println(tabla[i])
		}else{
			fmt.Println("Te jodiste")
		}
	}

	matriz [2][6] = 5

	fmt.Println(matriz)*/

	/*matriz := []int {2, 3, 88, 99}*/

	variante()
	variante2()
	variante3()
}

func variante() {
	elemntos := [5] int {1, 2, 3, 4, 5}
	porcion := elemntos [3:5]

	fmt.Println(porcion)
}

func variante2() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d capacodad %d", len(elementos), cap(elementos))
}

func variante3() {
	numbs := make([]int, 0, 0)
	for i := 0; i < 100; i++{
		numbs=append(numbs, i)
	}
	fmt.Printf("\n largo %d Capacidad %d", len(numbs), cap(numbs))
}