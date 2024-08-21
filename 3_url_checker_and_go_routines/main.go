package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

// 3-6. URLChecker + Goroutines
type requestResult struct {
	url    string
	status string
}

func main() {
	// 3-0. hitURL / 3-1. slow URLChecker / 3-6. URLChecker + Goroutines
	// var results map[string]string // -> panic
	// var results = map[string]string{}
	results := make(map[string]string) // empty map 생성
	// -> 이렇게 하지 않으면 map이 nil이 되고, nil인 map엔 값을 넣을 수 없다.
	c2 := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	//results["hello"] = "hello" // panic(컴파일러가 못 찾아내는 에러) 발생
	// -> 초기화되지 않은 map에 값을 넣을 수 없다.
	for _, url := range urls {
		// url에 대해 성공인지 실패인지
		result := "OK"
		err := hitURL1(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
	fmt.Println("--------------------")

	// 3-6. URLChecker + Goroutines
	for _, url := range urls {
		go hitURL2(url, c2)
	}

	// 3-7. fast URLChecker
	for i := 0; i < len(urls); i++ {
		result := <-c2
		// fmt.Println(result) // 받고 있는 메시지들
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
	fmt.Println("--------------------")

	// 3-2. Goroutiones / 3-3.channels
	/* people := [2]string{"seongjin", "beomgeun"}
	for _, person := range people {
		go isSexy(person)
	} // 메인 함수가 이미 끝나 작동하지 않는다.
	go sexyCount("seongjin") // go를 붙여 작업을 병렬적으로 진행한다.
	sexyCount("beomgeun")    // 이 경우 beomgeun이 카운트 되므로 동시 작업이 가능하다.
	// go sexyCount("beomgeun") // 여기에 go를 붙이면 아무 일도 일어나지 않는다(프로그램(메인 함수) 종료).
	// -> goroutines은 프로그램이 작동하는 동안만 유효하다(메인 함수 실행동안만).
	// -> goroutines를 마치면 더 이상 진행할 작업이 없게 된다. 메인 함수는 다른 goroutines를 기다려주지 않는다.
	// 메인 함수의 끝 = goroutines 소멸
	time.Sleep(time.Second * 3) // goroutines가 5초 동안 살아있는다.
	fmt.Println("--------------------") */

	// 3-3.channels / 3-4. channels recap
	// goroutines와 메인 함수 간의 정보를 주고 받는 방법
	// (메인 함수 - 결과를 저장하는 곳)
	// channels - goroutione과 메인 함수 사이에 정보를 전달하기 위한 방법 or goroutines에서 다른 goroutines로 커퓨니케이션
	c1 := make(chan string) // make - channel 생성 / (chan 보낼 정보의 타입)
	people := [4]string{"seongjin", "beomgeun", "edword", "bella"}
	for _, person := range people {
		go isSexy(person, c1)
	} // 파이프 만들기 -> isSexy 함수가 seongjin이나 beomgeun이 섹시한지 체크하고 그 결과를 알려줄 수 있게 만든다.
	// 메인이 메시지를 받게 한다.
	/* fmt.Println("Waiting for message")
	resultOne := <-c1
	resultTwo := <-c1
	// resultThree := <-c1
	fmt.Println("Received this message:", resultOne) // 메시지 하나 받을 때까지 기다리고(blocking operation: 작업이 끝날 때까지 멈춘다.) 받으면 계속해서 메시지를 프린트한다.
	fmt.Println("Received this message:", resultTwo)
	// fmt.Println("Received this message:", resultThree) // 메시지를 계속 기다리고 있는데 goroutines는 끝나버려서 deadlock이 발생한다.
	// time.Sleep(time.Second * 10) */
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i, ": ")
		fmt.Println(<-c1) // 동시에 일어난다. (순서대로 출력되지 않는다)
	}

	// 메시지를 받을 곳이 없어도 메시지를 보낼 수 있다.
}

// 3-0. hitURL / 3-1. slow URLChecker
// 웹사이트로 접속(hit: 인터넷 웹 서버의 파일 1개에 접속하는 것)하고 그 결과를 알려주는 함수
func hitURL1(url string) error {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)                // 웹 사이트 접속 (request)
	if err != nil || resp.StatusCode >= 400 { // 400 이상부턴 뭔가 문제가 발생했다는 뜻
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

// 3-2. Goroutiones
// Go의 최적화 방법 - 동시에 작업한다.
// Goroutiones: 기본적으로 다른 함수와 동시에 실행시키는 함수
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second) // 1초동안 sleep
	} // top-down (일반적 프로그래밍 방식)
}

// 3-3. channels
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	// fmt.Println(person)
	// return true
	c <- person + " is sexy" // channel에 true라는 메시지를 보낸다.
}

// 정리 !!
// 하나의 채널을 만든다.
// 그 채널에 두 함수(isSexy seongjin, isSexy beomgeun)로 보낸다.
// 이 두 함수는 5초 뒤에 true라는 2개의 메시지를 보내준다.
//   main
//     |
//    / \
//  url url  -> 하나의 채널이 있고 함수들이 메인 함수와 커뮤니케이션 하기 위해 그 하나의 채널을 같이 쓴다.

// 3-5. one more recap
// < 채널의 룰 >
// 채널과 goroutine이 먼저이다.
// 메인 함수가 끝나면 goroutine은 끝났던 아니던간에 무의미해진다.
// 받을 데이터와 채널을 통해 보낼 데이터에 대해서 어떤 타입을 받을 것인지를 구체적으로 지정해줘야 한다.
// 메시지를 보내는 방법은 <-(메시지)

// 3-6. URLChecker + Goroutines
func hitURL2(url string, c chan<- requestResult) { // Send Only (보내기만 가능)
	// fmt.Println("Checking:", url)
	// c <- result{}    // 채널로 보낸다.
	// fmt.Println(<-c) // 채널로부터 받는다. -> 실행 X
	resp, err := http.Get(url)
	status := "OK"
	// 에러가 있으면 체널에 result(struct)를 보낸다.
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
} // 이 함수에 데이터를 보내기만 하고 받을 순 없는 채널이 있다.
