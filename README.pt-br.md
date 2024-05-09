[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-defer/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-defer/blob/main/README.md)

## Defer (Adiar)
Clausula ```defer```, basicamente adia, a execução de determinado trecho de código, função, funções de pacote etc.  

### Iniciando projeto
Crie um diretório ```go-defer``` contendo um arquivo ```main.go``` com o seguinte conteúdo inicial:  
```go
package main
import "fmt"

func f1() {
	fmt.Println("f1")
}
func f2() {
	fmt.Println("f2")
}
func main() {
	f1()
	f2()
}
//outputs
f1
f2
```
Agora vejamos o que acontece caso adicionemos nossa clausula ```defer```, que significa "adiar", como prefixo de ```f1()```:  
```go
....
func main() {
	defer f1()
	f2()
}
/// output
f2
f1
```
O que acontece aqui é que ela diz pro seu código adiar a função ```f1()``` até o ultimo momento possivel, como estamos lidando com a função ```main``` e ela não tem retorno, o último momento possivel é antes dela terminar.  
Vejamos mais exemplos:
```go
...

func studentAproved(n1, n2 float32) bool {
  fmt.Println("Verificando se o aluno está aprovado")
  average := ((n1 + n2) /2)

  if average >= 6 {
    return true
  }
  return false
}

func main() {
	defer f1()
	f2()

	fmt.Println(studentAproved(7, 8)) // output true

  //outputs
  f2
  Verificando se o aluno está aprovado
  true
  f1
}
```
Então sempre que o ```defer``` é usado como prefixo de algo, como uma função um print que seja, ele sempre será exucutado no último momento possível.  
