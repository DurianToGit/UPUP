package main

import "sync"

func main() {
	data := make(map[int]int)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				data[workerID*1000+j] = j
			}
		}(i)
	}

	wg.Wait()
}
