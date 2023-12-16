package main

import (
	"fmt"
	"time"
)

func main() {

	// 사실 이 용량을 결정하기가 굉장이 어려움.. (메모리 낭비로 이어질수 있음)
	// produce -> consume 하는 구조로 하면 편해짐...
	msgCh := make(chan string, 128) // 일단 지정

	for _, v := range []string{"A", "B", "C", "D"} {
		msgCh <- fmt.Sprintf("msg %s", v)
	}
	close(msgCh)

	// deadlock (실제 channel에 존재하는 값들만 출력하고싶다면?)
	// close를 사용해서 produce한 것들만 consume

	// first for
	// for msg := range msgCh {
	// 	fmt.Println(msg)
	// }

	for {
		msg, ok := <-msgCh
		if !ok {
			break
		}

		fmt.Println(msg)
	}

	fmt.Println("done reading all the message")

}

func fetchResponse() string {
	time.Sleep(time.Second * 2)
	return "hello world"
}
