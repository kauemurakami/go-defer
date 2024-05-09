package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func studentAproved(n1, n2 float32) bool {
	fmt.Println("Verificando se o aluno está aprovado")
	average := ((n1 + n2) / 2)

	if average >= 6 {
		return true
	}
	return false
}

func main() {
	defer f1()
	f2()

	fmt.Println(studentAproved(7, 8)) // output true
}
