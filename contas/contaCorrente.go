package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float32
}

func (c *ContaCorrente) sacar(valorDoSaque float64) {

	podeSacar := valorDoSaque <= float64(c.Saldo) && valorDoSaque > 0
	if podeSacar == true {
		c.Saldo = c.Saldo - float32(valorDoSaque)
		fmt.Println("Saque realizado com sucesso. Seu Saldo é de: R$", c.Saldo)
	} else {
		fmt.Println("Saque inválido. Seu Saldo é de: R$", c.Saldo)
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {

	if valorDoDeposito > 0 {
		c.Saldo = c.Saldo + float32(valorDoDeposito)
		fmt.Println("Depósito efetuado com sucesso. Seu Saldo é de: R$", c.Saldo)
	} else {
		fmt.Println("Depósito inválido. Seu Saldo é de: R$", c.Saldo)
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorDaTransferencia < float64(c.Saldo) && valorDaTransferencia > 0 {
		c.Saldo = c.Saldo - float32(valorDaTransferencia)
		contaDestino.depositar(valorDaTransferencia)
		fmt.Println("Tranferência realizada com sucesso. Seu Saldo é de: R$", c.Saldo)
		return true
	} else {
		fmt.Println("Transferencia inválida. Seu Saldo é de: R$", c.Saldo)
		return false
	}
}
