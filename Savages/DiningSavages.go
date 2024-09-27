package main

import (
	"fmt"
	"semaforo/FPPDSemaforo"
)

var (
	servings = 0
	emptyPot = FPPDSemaforo.NewSemaphore(0)
	fullPot = FPPDSemaforo.NewSemaphore(0)
	mutex = FPPDSemaforo.NewSemaphore(1)
)

const (
	M = 20
	COOK = 10
	SAVAGES = 200
)

func putServingsInPot(i int) {
	fmt.Printf("Cozinheiro %d = está colocando comida no pote!\n", i)
	servings = M
}

func getServingFromPot(i int) {
	fmt.Printf("Selvagem %d = está servindo seu prato!\n", i)
}

func eat(i int) {
	fmt.Printf("Selvagem %d = está comendo!\n", i)
}

func cook(i int) {
	for {
		emptyPot.Wait()
		putServingsInPot(i)
		fullPot.Signal()
	}
}

func savages(i int) {
	for {
		mutex.Wait()
		if servings == 0 {
			emptyPot.Signal()
			fullPot.Wait()
		}
		servings -= 1
		getServingFromPot(i)
		mutex.Signal()

		eat(i)
	}
}

func main() {
	for i := 0; i<SAVAGES; i++ {
		go savages(i)
	}

	for i := 0; i<COOK; i++ {
		go cook(i)
	}

	for {}
}