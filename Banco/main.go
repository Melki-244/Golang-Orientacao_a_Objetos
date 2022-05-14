package main

import "fmt"

/* Criando Uma Estrutura em Go*/
type ContaCorrente struct {
  titular string
  numeroDaConta int
  numeroDaAgencia int
  saldo float64
}
func (conta *ContaCorrente) Sacar(valorDoSaque float64) string  {
  podeSacar := valorDoSaque >= 0 && valorDoSaque <= conta.saldo 

  if podeSacar {
    conta.saldo -= valorDoSaque 
    return "Saque Realizado Com Sucesso"
  }else {
    return "Saldo Insulficiente"
  }
}

func main()  {
  contaDaSilvia := ContaCorrente{}
  contaDaSilvia.titular = "Silvia" 
  contaDaSilvia.saldo = 77.94

  fmt.Println(contaDaSilvia.saldo)
  
  fmt.Println(contaDaSilvia.Sacar(-50))

  fmt.Println(contaDaSilvia.saldo)
}
