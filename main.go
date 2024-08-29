package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float32) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float32) string
}

func main() {
	contaDoJean := contas.ContaPoupanca{}
	contaDoJean.Depositar(100)
	PagarBoleto(&contaDoJean, 60)

	fmt.Println(contaDoJean.ObterSaldo())

	contaDoRodrigo := contas.ContaCorrente{}
	contaDoRodrigo.Depositar(500)
	PagarBoleto(&contaDoRodrigo, 400)

	fmt.Println(contaDoRodrigo.ObterSaldo())
}
