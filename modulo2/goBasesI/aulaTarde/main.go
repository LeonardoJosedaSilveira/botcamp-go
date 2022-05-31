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

func exercise3() {
	ms := []string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	mes1 := 1
	mes10 := 10
	mes12 := 12

	fmt.Println("Mês 1: ", ms[mes1-1])
	fmt.Println("Mês 10: ", ms[mes10-1])
	fmt.Println("Mês 12: ", ms[mes12-1])
}

func exercise4() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel": 26,
		"Brenda": 19,
		"Dario": 44,
		"Pdro": 30,
	}

	fmt.Println("idade de Benjamin:", employees["Benjamin"])

	c := 0
	for _, v := range employees {
		if v > 21 {
			c++
		}
	}
	fmt.Println("funcionários com mais de 21 anos: %d\n", c)

	fmt.Println("##### adicionar Frederico")
	fmt.Println("antes:", employees)
	employees["Frederico"] = 25
	fmt.Println("depois:", employees)

	fmt.Println("##### excluir Pedro")
	fmt.Println("antes:", employees)
	delete(employees, "Pedro")
	fmt.Println("depois:", employees)
}

func main() {
	exercise4()
}