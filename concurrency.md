# Goroutine-Channel-Concurrency

## Goroutine

- 1. 기본적인 고루틴을 사용해서는 병렬처리를 할 수 없음

  ```go
    go func() {
   	    res := fetchResource()
   	    fmt.Println(res)
    }()
  ```

- 2. 고루틴을 사용하기 위해서는 무조건 Channel을 사용해야 함

  ```go
    resultch := make(chan string)
  	resultch <- "foo" // <- Full (Block)
  	res := <-resultch
  	fmt.Println(res)
  ```

  - 하지만, 위와같은 방법으로 구성하면 <br>Deadlock</br> 이 발생함
  - 데드락이 발생하는 이유는 make(chan string) 현재 1이고 Full 즉 Block이 발생하기 때문에

- 3. Buffered Channel을 사용해서 진행해야 함

  ```go
    bufferedCh := make(chan string, 2)
    bufferedCh <- "foo"
    bufferedCh <- "foo"
    bufferedCh <- "foo" // deadlock (Full)
    resultBuffer := <-bufferedCh
    fmt.Println(resultBuffer)
  ```

  - 위와같이 chan 용량이 지정해준다면 에러가 안난다.

- 4. 결과론적으로 코드 라인은 goroutine을 사용하되, 결과를 channel에 받도록 진행해야 함

  ```go
  resCh := make(chan string, 1)
  go func() {
  	resCh <- fetchResource()
  }()

  res := <-resCh
  fmt.Println(res)
  ```

- 5. close를 사용해서 Produce 그만하기

  ```go

  // ...
  close(msgCh)

  // 첫번째 방법
  for msg := range msgCh {
  	fmt.Println(msg)
  }

  // 두번째 방법
  for {
    msg, ok := <-msgCh
    if !ok {
      break
    }

    fmt.Println(msg)
  }

  fmt.Println("done reading all the message")
  ```
