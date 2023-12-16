package main

import (
	"fmt"
	"time"
)

// 고루틴은 병렬이 아님
func main() {

	// 1. one > 2
	// status, response := fetchResource()
	// fmt.Println(status, response)

	// 2. multiple > 2 * 10
	// fetchResource()
	// fetchResource()
	// fetchResource()
	// fetchResource()
	// fetchResource()
	// fetchResource()
	// fetchResource()
	// fetchResource()
	// fetchResource()

	// 3. goroutine (async) -> 이대로는 실행이 보장 되지 않는다.
	// go func() {
	// 	res := fetchResource()
	// 	fmt.Println(res)
	// }()

	// 4. use Channel // deadlock
	// 4.1 unbuffered channel
	// 4.2 buffered channel

	// channle은 애초에 다 차있는 상태이기때문에, 용량을 지정해줘야 한다
	// resultch := make(chan string, 1) // -> unbuffered channel
	// resultch <- "foo"
	// res := <-resultch
	// fmt.Println(res)

	// channel에 용량 지정
	// bufferedCh := make(chan string, 2)
	// bufferedCh <- "foo"
	// bufferedCh <- "foo"
	// // bufferedCh <- "foo" // deadlock (Full)
	// resultBuffer := <-bufferedCh
	// fmt.Println(resultBuffer)

	resCh := make(chan string, 10)
	go func() {
		for i := 0; i < 10; i++ {
			resCh <- fetchResource()
		}
	}()

	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
	fmt.Println(<-resCh)
}

func fetchResource() string {
	time.Sleep(time.Second * 2)

	return fmt.Sprintf("Result %d", 200)
}
