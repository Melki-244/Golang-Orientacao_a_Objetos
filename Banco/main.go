package main

import "fmt"

/* Criando Uma Estrutura em Go*/
type ContaCorrente struct {
  nome string
  numeroDaConta int
  numeroDaAgencia int
  saldo float64
}

func main()  {

  /*Atribuindo valores de forma Explicita*/
  contaGuilherme := ContaCorrente{nome: "Guilherme", numeroDaConta: 111333, numeroDaAgencia:999000, saldo: 30.58} 
  /* Atribuindo Valores de forma Impicita*/
  contaMelki := ContaCorrente{"Melki", 111222, 999111, 99.08}

  fmt.Println(contaMelki)
  fmt.Println(contaGuilherme)

  /* Neste Exemplo vemos um modo diferente de como trabalhar com dados de uma estrutura 
  ultilizando o metodo new
  */
  var contaDaChris *ContaCorrente  
  contaDaChris = new(ContaCorrente) 
  contaDaChris.nome = "Chris"
  contaDaChris.saldo = 75.88

  fmt.Println(contaDaChris)
}
