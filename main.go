package main

import (
	_ "expvar"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const maxWorkers = 5

var originalRequest int

type job struct {
	name     string
	duration time.Duration
}

func doWork(id int, j job) {
	fmt.Printf("worker%d: started %s, working for %fs\n", id, j.name, j.duration.Seconds())
	time.Sleep(j.duration)
	fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func main() {

	var currentRequests int

	wg := &sync.WaitGroup{}
	for {
		var size int
		var count int = 0
		jobs := make(chan job)

		currentRequests = 12

		requests := currentRequests

		//////
		fmt.Println(requests)
		{
			originalRequest = requests
			requests = int(math.Floor(float64(requests / 5)))
			requests = requests * 5
			if requests == 0 {
				size = 0
			} else {
				size = maxWorkers
			}

			wg.Add(size)
			for j := 1; j <= size; j++ {
				go func(j int) {
					defer wg.Done()
					for c := range jobs {
						doWork(j, c)
					}
				}(j)
			}

			for i := 0; i < requests; i++ {
				name := fmt.Sprintf("job-%d", i)
				duration := time.Duration(rand.Intn(1000)) * time.Millisecond
				fmt.Printf("adding: %s %s\n", name, duration)
				jobs <- job{name, duration}
			}

			//time.Sleep(3 * time.Second)
			close(jobs)
			wg.Wait()

			////////for the remaining number of requests.....
			fmt.Println("65 left")
			size = originalRequest - requests
			fmt.Println(size)
			wg = &sync.WaitGroup{}
			jobs := make(chan job)
			wg.Add(size)
			for j := 1; j <= size; j++ {
				go func(j int) {
					defer wg.Done()
					for c := range jobs {
						doWork(j, c)
					}
				}(j)

				////////////////////////
				ticker := time.NewTicker(1 * time.Second)
				go func() {
					currentRequests++ //////////////////////////////
					fmt.Println(currentRequests)
					for _ = range ticker.C {
						if size == 5 || count == 10 {
							ticker.Stop()
						}

						// currentRequests = 17 ///

						extra := currentRequests - originalRequest
						if extra == 0 {
							count++
						}

						for size < 5 && extra > 0 {
							size++
							extra--
							originalRequest++

							go func(j int) {

								defer wg.Done()
								for c := range jobs {
									doWork(j, c)
								}

							}(size)

						}

					}
				}()
			}
			// originalRequest = originalRequest + count
			fmt.Println(currentRequests) ///
			for i := requests; i < currentRequests; i++ {
				name := fmt.Sprintf("job-%d", i)
				duration := time.Duration(rand.Intn(1000)) * time.Millisecond
				fmt.Printf("adding: %s %s\n", name, duration)
				jobs <- job{name, duration}
			}
			time.Sleep(15 * time.Second)
			close(jobs)

		}
	}
	// wg.Wait()

}
