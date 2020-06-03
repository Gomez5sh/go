package main

import (
	"fmt"
	"time"
	us "./user"
)

type pipo struct{
	us.Usuario

}

func main()  {
	u := new(pipo)
	u.AltaUsuario(1, "Jhon wick", time.Now(), true)
	fmt.Println(u.Usuario)
}