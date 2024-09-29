package main

import (
	"fmt"
	"semaforo/FPPDSemaforo"
	"time"
)

var (
	servings = 0
	mutex = FPPDSemaforo.NewSemaphore(1)
	emptyPot = FPPDSemaforo.NewSemaphore(0)
	fullPot = FPPDSemaforo.NewSemaphore(0)
)

const (
	M = 20
	SAVAGES = 200
)

func putServingsInPot() {
	fmt.Printf("Cozinheiro está colocando comida no pote!\n")
	time.Sleep(100 * time.Millisecond)
	servings = M
}

func getServingFromPot(i int) {
	fmt.Printf("Selvagem %d está servindo seu prato!\n", i)
	time.Sleep(100 * time.Millisecond)
	servings -= 1
}

func eat(i int) {
	fmt.Printf("Selvagem %d está comendo!\n", i)
	time.Sleep(100 * time.Millisecond)
}

func cook() {
	for {
		emptyPot.Wait()
		putServingsInPot()
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
		getServingFromPot(i)
		mutex.Signal()
	
		eat(i)
	}
}

func main() {
	for i := 0; i<SAVAGES; i++ {
		go savages(i)
	}

	go cook()

	for {}
}