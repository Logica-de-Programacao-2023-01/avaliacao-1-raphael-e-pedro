package q4

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	package main

import (
	"errors"
	"fmt"
)

func classifyPrices(prices []int) (int, error) {
	n := len(prices)
	if n == 0 {
		return 0, errors.New("a lista está vazia")
	}
	if n == 1 {
		return 3, nil
	}
	Crescente := true
	Decrescente := true
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			Decrescente = false
		} else if prices[i] < prices[i-1] {
			Crescente = false
		}
	}
	if Crescente {
		return 1, nil
	} else if Decrescente {
		return 2, nil
	} else {
		return 3, nil
	}
}

func main() {
	prices1 := []int{5, 4, 3, 2, 1}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{3, 2, 4, 1, 5}
	prices4 := []int{4}

	result1, err1 := classifyPrices(prices1)
	result2, err2 := classifyPrices(prices2)
	result3, err3 := classifyPrices(prices3)
	result4, err4 := classifyPrices(prices4)

	if err1 != nil {
		fmt.Println("Erro:", err1)
	} else {
		fmt.Println("Resultado 1:", result1)
	}

	if err2 != nil {
		fmt.Println("Erro:", err2)
	} else {
		fmt.Println("Resultado 2:", result2)
	}

	if err3 != nil {
		fmt.Println("Erro:", err3)
	} else {
		fmt.Println("Resultado 3:", result3)
	}

	if err4 != nil {
		fmt.Println("Erro:", err4)
	} else {
		fmt.Println("Resultado 4:", result4)
	}
}

	return 0, nil
}
