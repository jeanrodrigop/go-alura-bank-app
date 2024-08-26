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

	fmt.Println(contaDoJean)
}
