package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoJean := contas.ContaCorrente{}
	contaDoJean.Depositar(100)

	fmt.Println(contaDoJean.ObterSaldo())
}
