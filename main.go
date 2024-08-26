package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float32
}

func main() {
	contaDoJean := ContaCorrente{"Jean Rodrigo", 589, 123456, 125.5}

	fmt.Println(contaDoJean.saldo)

	fmt.Println(contaDoJean.sacar(80))
	fmt.Println(contaDoJean.saldo)
}

func (c *ContaCorrente) sacar(valorDoSaque float32) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insulficiente"
	}
}
