package main

import (
	"fmt"
)

func exercise1(wordIn string) {
	var word string = wordIn
	var maxCharacters int = len(word)

	println(maxCharacters)

	for _, character := range word {
		fmt.Println(string(character))
	}

}

func exercise2() {
	clients := []map[string]interface{}{
		{
			"Nome": "leo",
			"Idade": 12,
			"Atividade": 1,
			"Salario": 500,
		},
		{
			"Nome": "jao",
			"Idade": 18,
			"Atividade": 2,
			"Salario": 2000,
		},		{
			"Nome": "jao",
			"Idade": 28,
			"Atividade": 1,
			"Salario": 3000,
		},		{
			"Nome": "jao",
			"Idade": 38,
			"Atividade": 4,
			"Salario": 4000,
		},		{
			"Nome": "jao",
			"Idade": 48,
			"Atividade": 4,
			"Salario": 10001,
		},
	}

	for _, client := range clients {
		name := client["Nome"]
		fmt.Println("Cliente: ", name)

		age := client["Idade"].(int)
		if age <= 22 {
			fmt.Println("Não tem idade pra isso")
			continue
		}

		activity := client["Atividade"].(int)
		if activity <= 1 {
			fmt.Println("Não atividade para isso")
			continue
		}

		wage := client["Salario"].(int)
		if wage > 10000 {
			fmt.Println("Emprestimo disponivel sem juros")
		} else {
			fmt.Println("Emprestimo disponivel com juros")
		}
	}
}

func main() {
	exercise2()
}