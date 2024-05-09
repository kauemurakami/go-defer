[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-defer/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-defer/blob/main/README.md)

## Defer 
```defer``` clause, basically postpones the execution of a certain piece of code, function, package functions, etc.  

### Starting project
Create a ```go-defer``` directory containing a ```main.go``` file with the following initial content:  
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
Now let's see what happens if we add our ```defer``` clause, which means "postpone", as a prefix to ```f1()```:  
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
What happens here is that it tells your code to postpone the ```f1()``` function until the last possible moment, as we are dealing with the ```main``` function and it has no return, the last possible moment is before it ends.  
Let's look at more examples:  
```go
...

func studentAproved(n1, n2 float32) bool {
  fmt.Println("Checking if the student is approved")
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
  Checking if the student is approved
  true
  f1
}
```
So whenever ```defer``` is used as a prefix for something, such as a print function, it will always be executed at the last possible moment.
