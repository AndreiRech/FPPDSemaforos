package main

import (
	"fmt"
	"semaforo/FPPDSemaforo"
	// "time"
)

const STUDENTS = 200

var (
	eating = 0
	readyToLeave = 0
	okToLeave = FPPDSemaforo.NewSemaphore(0)
	mutex = FPPDSemaforo.NewSemaphore(1)
)

func getFood(i int) {
	fmt.Printf("Estudante %d entrou e está indo buscar a comida!\n", i)
}

func dine(i int) {
	fmt.Printf("Estudante %d está jantando!\n", i)
	// time.Sleep(200 * time.Millisecond)
}

func leave(i int) {
	fmt.Printf("Estudante %d saindo da sala...\n", i)
}

func students(i int) {
	getFood(i)

	mutex.Wait()
	eating++
	if eating == 2 && readyToLeave == 1 {
		okToLeave.Signal()
		readyToLeave--
	}
	mutex.Signal()

	dine(i)

	mutex.Wait()
	eating--
	readyToLeave++

	if eating == 1 && readyToLeave == 1 {
		mutex.Signal()
		okToLeave.Wait()
	} else if eating == 0 && readyToLeave == 2 {
		okToLeave.Signal()
		readyToLeave -= 2
		mutex.Signal()
	} else {
		readyToLeave--
		mutex.Signal()
	}

	leave(i)
}

func main() {
	for i := 0; i<STUDENTS; i++ {
		go students(i)
	}

	for {}
}