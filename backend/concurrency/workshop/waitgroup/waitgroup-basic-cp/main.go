package main

import "sync"

func testWG(output chan<- []bool) {
	// TODO: answer here

	ch := make(chan int)
	var wg sync.WaitGroup
	done := make([]bool, 5)

	for i := 0; i < 5; i++ {
		count := 0
		// TODO: answer here
		wg.Add(1)
		go func(i int) {
			// TODO: answer here
			defer wg.Done()
			done[i] = true
			count <- i

		}(i)
	}

	// TODO: answer here
	wg.Wait()
	output <- done
}
