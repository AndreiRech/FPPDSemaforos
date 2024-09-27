package main

import (
	"fmt"
	"semaforo/FPPDSemaforo"
)

const PHILOSOPHERS = 5

var (
	state = []string{"thinking", "thinking", "thinking", "thinking", "thinking"}
	sem  [PHILOSOPHERS]*FPPDSemaforo.Semaphore
	mutex = FPPDSemaforo.NewSemaphore(1)
)

func createSemaphores() {
	for i := 0; i < PHILOSOPHERS; i++ {
		sem[i] = FPPDSemaforo.NewSemaphore(0)
	}
}

func left(i int) int {
	return (i + PHILOSOPHERS - 1) % PHILOSOPHERS
}

func right(i int) int {
	return (i + 1) % PHILOSOPHERS
}

func getFork(i int) {
	mutex.Wait()
	state[i] = "hungry"
	test(i)
	mutex.Signal()
	sem[i].Wait()
}

func putFork(i int) {
	mutex.Wait()
	state[i] = "thinking"
	test(right(i))
	test(left(i))
	mutex.Signal()
}

func test(i int) {
	if state[i] == "hungry" && state[left(i)] != "eating" && state[right(i)] != "eating" {
		state[i] = "eating"
		sem[i].Signal()
	}
}

func philosophers(i int) {
	for {
		fmt.Printf("Filósofo [%d] está pensando.\n", i)
		getFork(i)
		fmt.Printf("Filósofo [%d] está comendo.\n", i)
		putFork(i)
	}
}

func main() {
	createSemaphores()

	for i := 0; i < PHILOSOPHERS; i++ {
		go philosophers(i)
	}

	for {}
}
