package main

import "fmt"

/* Criando Uma Estrutura em Go*/
type ContaCorrente struct {
  titular string
  numeroDaConta int
  numeroDaAgencia int
  saldo float64
}
func (conta *ContaCorrente) Saca (valorDoSaque float64) (string, float64)  {
  podeSacar := valorDoSaque >= 0 && valorDoSaque <= conta.saldo 

  if podeSacar {
    conta.saldo -= valorDoSaque 
    return "Saque Realizado Com Sucesso", conta.saldo
  }else {
    return "Saldo Insulficiente", conta.saldo
  }
}
func (conta *ContaCorrente) Deposita (valorDoDeposito float64) (string, float64) {
  podeDepositar := valorDoDeposito >= 0
  if podeDepositar {
    conta.saldo += valorDoDeposito
    return "Deposito Realizado Com Sucesso!" , conta.saldo
  }else {
    return "Deposito Não Realizado" , conta.saldo
  }
}
func (conta *ContaCorrente) Transfere (valorDaTranferencia float64, contaDestino *ContaCorrente) (string,bool) {
  podeTranferir := valorDaTranferencia <= conta.saldo && valorDaTranferencia >= 0   
  if podeTranferir {
    conta.saldo -= valorDaTranferencia
    contaDestino.saldo += valorDaTranferencia 
    return "Transação Realizada com Sucesso!", true
  }else {
    return "A Transação Não Pode Ser Efetuada", false
  }
}

func main()  {
  contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300.00}
  contaDoMarcos := ContaCorrente{titular: "Marcos", saldo: 1200.00}

  mensagem, status := contaDaSilvia.Transfere(-300,&contaDoMarcos)

  fmt.Println(mensagem,status)
  fmt.Println(contaDaSilvia.saldo)
  fmt.Println(contaDoMarcos.saldo)
}

