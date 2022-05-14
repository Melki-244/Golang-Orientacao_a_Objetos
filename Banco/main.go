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
func (c *ContaCorrente) Deposita (valorDoDeposito float64) (string, float64) {
  podeDepositar := valorDoDeposito >= 0
  if podeDepositar {
    c.saldo += valorDoDeposito
    return "Deposito Realizado Com Sucesso!" , c.saldo
  }else {
    return "Deposito Não Realizado" , c.saldo
  }
}

func main()  {
  contaDaSilvia := ContaCorrente{}
  contaDaSilvia.titular = "Silvia" 
  contaDaSilvia.saldo = 77.94
  
  status, saldo := contaDaSilvia.Deposita(300) 

  fmt.Println(status, saldo)

  status, saldo = contaDaSilvia.Saca(400)

  fmt.Println(status, saldo)
}
