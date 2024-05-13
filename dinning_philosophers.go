package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var (
	eatingPerDay = 3
	eatTime      = 1 * time.Second
	thinkTime    = 3 * time.Second
	sleepTime    = 1 * time.Second
)

var (
	orderMutex    sync.Mutex
	orderFinished []string
)

func dine() {

	numOfPhilosophers := len(philosophers)

	/*
	* wg is the WaitGroup that keeps track of how many philosophers are still at the table. When it reaches zero, everyone
	* is finished eating and has left.
	 */
	wg := &sync.WaitGroup{}

	/*
	* We want everyone to be seated before they start eating, so create a WaitGroup for that
	 */
	seated := &sync.WaitGroup{}

	wg.Add(numOfPhilosophers)
	seated.Add(numOfPhilosophers)

	/*
	* forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher type.
	* Each fork, then, can be found using the index (an integer), and each fork has a unique mutex.
	 */
	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < numOfPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < numOfPhilosophers; i++ {
		go diningProblem(philosophers[i], wg, seated, forks)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, seated *sync.WaitGroup, forks map[int]*sync.Mutex) {

	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	seated.Done()
	seated.Wait()

	for i := eatingPerDay; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, "is satisified.")
	fmt.Println(philosopher.name, "left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}

func run_dinning_philosophers() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)

	dine()

	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)
	fmt.Printf("Order finished: %s.\n", strings.Join(orderFinished, ", "))
}
