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

  /*Atribuindo valores de uma estrutura forma Explicita*/
  contaGuilherme := ContaCorrente{nome: "Guilherme", numeroDaConta: 111333, numeroDaAgencia:999000, saldo: 30.58} 
  contaGuilherme2 := ContaCorrente{nome: "Guilherme", numeroDaConta: 111333, numeroDaAgencia:999000, saldo: 30.58} 
  /* Atribuindo Valores de uma estrutura de forma Impicita*/
  contaMelki := ContaCorrente{"Melki", 111222, 999111, 99.08}
  contaMelki2 := ContaCorrente{"Melki", 111222, 999111, 99.08}

  fmt.Println(contaMelki == contaMelki2)
  fmt.Println(contaGuilherme == contaGuilherme2)

  /* Neste Exemplo vemos um modo diferente de como trabalhar com dados de uma estrutura 
  ultilizando o metodo new
  */
  var contaDaChris *ContaCorrente  
  contaDaChris = new(ContaCorrente) 
  contaDaChris.nome = "Chris"
  contaDaChris.saldo = 75.88
  
  var contaDaChris2 *ContaCorrente  
  contaDaChris = new(ContaCorrente) 
  contaDaChris.nome = "Chris"
  contaDaChris.saldo = 75.88

  fmt.Println(&contaDaChris)
  fmt.Println(&contaDaChris2)

  fmt.Println(*contaDaChris == *contaDaChris2)
}
