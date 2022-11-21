package main

import "fmt"

func escribirEnCanal(texto string, c chan string) {
	c <- texto
}

func main() {
	c := make(chan string)
	go escribirEnCanal("Dato de un canal", c)
	fmt.Println(<-c)

}
