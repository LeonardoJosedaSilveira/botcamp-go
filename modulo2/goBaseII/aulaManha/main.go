package main

import (
	"errors"
	"fmt"
)

func caulcalaImposto(salario float64) float64 {
	if salario <= 50000.00 {
		return 0.0
	} else if salario <= 150000.00 {
		return salario * 0.17
	}
	return salario * 0.27
}

func exercise1() {
	fmt.Printf("imposto de até R$50.000: %.2f\n", caulcalaImposto(50000))
	fmt.Printf("imposto de até R$150.000: %.2f\n", caulcalaImposto(150000))
	fmt.Printf("imposto mais que R$150.000: %.2f\n", caulcalaImposto(150001))
}

func calculaMedia(notas ...float64) float64 {
	sum := 0.0
	for _, n := range notas {
		sum += n
	}
	return sum / float64(len(notas))
}

func exercise2() {
	fmt.Printf("Calcula Média: %.2f\n", calculaMedia(10, 8, 9))
}

func calculaSalario(categoria string, horas int) (float64, error) {
	if categoria == "C" {
		return float64(horas) * 1000, nil
	}
	if categoria == "B" {
		salario := float64(horas) * 1500
		if horas > 160 {
			return salario * 1.2, nil
		}
		return salario, nil
	}
	if categoria == "A" {
		salario := float64(horas) * 3000
		if horas > 160 {
			return salario * 1.5, nil
		}
		return salario, nil
	}
	return 0.0, errors.New("Categoria Inválida")
}

func exercise3() {
	salario, _ := calculaSalario("C", 162)
	fmt.Printf("Salario 01: %.2f\n", salario)
	salario, _ = calculaSalario("B", 176)
	fmt.Printf("Salario 02: %.2f\n", salario)
	salario, _ = calculaSalario("A", 172)
	fmt.Printf("Salario 03: %.2f\n", salario)
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMin(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min, nil
}

func calcMax(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, nil
}

func calcAvg(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}
	avg := 0.0
	for _, v := range values {
		avg += v
	}
	return avg / float64(len(values)), nil
}

func getFunc(t string) (func(values ...float64) (float64, error), error) {
	if t == minimum {
		return calcMin, nil
	}
	if t == average {
		return calcAvg, nil
	}
	if t == maximum {
		return calcMax, nil
	}
	return nil, errors.New("invalid function type")
}

func exercise4() {
	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)

	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}

const (
	dog       = "dog"
	cat       = "cat"
	hamter    = "hamster"
	tarantula = "tarantula"
)

func dogFunc(quantity int) int {
	return quantity * 10000
}

func catFunc(quantity int) int {
	return quantity * 5000
}

func hamsterFunc(quantity int) int {
	return quantity * 250
}

func tarantulaFunc(quantity int) int {
	return quantity * 150
}

func Animal(t string) (func(quantity int) int, error) {
	if t == dog {
		return dogFunc, nil
	}
	if t == cat {
		return catFunc, nil
	}
	if t == tarantula {
		return tarantulaFunc, nil
	}
	return nil, errors.New("invalid animal type")
}

func exercise5() {
	dfunc, _ := Animal(dog)
	fmt.Printf("quantidade de alimento do(s) cachorro(s) (em gramas): %d gramas \n", dfunc(5))
	cFunc, _ := Animal(cat)
	fmt.Printf("quantidade de alimento do(s) gato(s) (em grams): %d gramas \n", cFunc(8))
	_, err := Animal("invalid animal")
	if err != nil {
		fmt.Println("animal invalido")
	}
}

func main() {
	exercise5()
}
