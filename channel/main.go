package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	maxGoroutines := 3
	listIds := []string{
		"ARBA01", "BRSP01", "XPR1", "BRNPR879", "BRNSP168",
		"BRNSP324", "BRNSP8", "MXNAGU934", "MXNBCS259", "MXNBCS285",
		"MXNCAM716", "MXNCAM773", "MXNCHH539", "MXNCHH813", "MXNCHH988",
		"MXNCHP552", "MXNCHP580", "MXNCOA565", "MXNCOA703", "MXNGUA166",
		"MXNGUA393", "MXNGUA738", "MXNJAL215", "MXNJAL359", "MXNMEX179",
		"MXNMEX688", "MXNMEX701", "MXNMIC860", "MXNNLE578", "MXNPUE6",
		"MXNQUE121", "MXNQUE812", "MXNROO373", "MXNROO839", "MXNSIN376",
		"MXNSIN458", "MXNTAB312", "MXNVER596", "MXNYUC161"}

	tasks := make(chan string, len(listIds))
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go makeRequest(&wg, tasks, i)
	}

	// Send tasks to the channel
	for _, id := range listIds {
		tasks <- id
	}

	// Close the channel after sending all tasks
	close(tasks)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All tasks finished")

}

func makeRequest(wg *sync.WaitGroup, tasks <-chan string, id int) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d started task %s\n", id, task)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished task %s\n", id, task)
	}
}
