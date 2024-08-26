package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoJean := contas.ContaCorrente{"Jean", 589, 123456, 300}
	contaDoRodrigo := contas.ContaCorrente{"Rodrigo", 589, 111333, 150.5}

	status := contaDoJean.Transferir(200, &contaDoRodrigo)

	fmt.Println(status)
	fmt.Println(contaDoJean)
	fmt.Println(contaDoRodrigo)
}
