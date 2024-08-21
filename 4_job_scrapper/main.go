package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

// 4-2. extractJob 1
type extractedJob struct {
	id                string
	title             string
	location          string
	career            string
	recruitmentPeriod string
}

// // 4-0. getPages 1 / 4-3. extractJob 2 / 4-6. more channels baby
func main() {
	var jobs []extractedJob // 배열들의 조합 [..] + [..] + .. + [..]

	c := make(chan []extractedJob) // 일자리 정보가 여러 개 전달되므로 타입은 []extractedJob(extractedJob X)

	totalPages := getPages()
	fmt.Println(totalPages)
	// URL를 hit 한다.
	for i := 0; i < totalPages; i++ {
		//extractedJobs := getPage(i, c)
		//jobs = append(jobs, extractedJobs...) // extracntedJobs의 contents를 jobs에 추가한다.
		// 모든 slice들을 더해서 하나로 만든다. [..] + [..] + .. + [..] = [..]
		// [[..] + [..] + .. + [..]] (X)
		go getPage(i, c)
		// getPage -> 중간다리 역, goroutine을 생성해서 일자리 전달 받고, main 함수의 채널로 전송한다.
		// main -> 페이지에서 일자리 정보를 전달 받아 jobs에 모아서 writeJob을 실행한다.
	}
	for i := 0; i < totalPages; i++ {
		extracntedjob := <-c // extractedJob을 getPage의 반환값에서 가져오는 대신 체널로 전달된 메시지 사용
		jobs = append(jobs, extracntedjob...)
	}
	fmt.Println(jobs)
	writeJobs2(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

// 4-1. getPages 2 / 4-2. extractJob 1 / 4-3. extractJob 2 / 4-5. channels time
func getPage(page int, mainC chan<- []extractedJob) { // 해당 페이지의 page 번호를 받는다.

	var jobs []extractedJob // 비어있는 jobs의 slice를 만든다.

	c := make(chan extractedJob)

	// 사람인에서 pagination이 동작하는 방법 - 각 페이지의 다른 위치(숫자)에서 start를 사용한다.
	//pageURL := baseURL + "&start=" + strconv.Itoa(page*50) // 정수 -> 텍스트
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page) + "&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n"
	// 페이지 1, 2, 3, ... 10 요청
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit") // 모든 카드를 가져온다.
	searchCards.Each(func(i int, card *goquery.Selection) {
		// job := extractJob(card, c)  // 찾아낸 각각의 카드에 대해 job 추출해 저장한다.
		// jobs = append(jobs, job) // 추출한 job을 업데이트한다.
		go extractJob(card, c) // goroutine, 값을 전달 받는다.
	})

	// 전달 받을 메시지의 수 = 카드의 개수 (각 카드마다 extractJob 함수가 1번씩 실행된다.)
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c // 채널에서 전송받은 메시지를 job 변수에 저장한다.
		// extractJob에서 보낸 메시지를 받는다.
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

// 4.3 extractJob 2
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	// job들을 id로 구분한다.
	// card: 찾은 각각의 카드(div)
	id, _ := card.Attr("value")                       // div는 value라는 속성을 가진다.
	title := card.Find(".area_job>.job_tit>a").Text() // >: inside
	location := cleanString(card.Find(".area_job>.job_condition>span").Eq(0).Text())
	career := card.Find(".area_job>.job_condition>span").Eq(1).Text()
	recruitmentPeriod := card.Find(".area_job>.job_date>.date").Text()
	c <- extractedJob{ // 채널에 값을 전송한다.
		id:                id,
		title:             title,
		location:          location,
		career:            career,
		recruitmentPeriod: recruitmentPeriod}
	// 각 페이지의 모든 카드에 대한 정보 출력
}

func getPages() int { // 페이지 수 반환
	// 4-1. getPages 2
	pages := 0

	// 4-0. getPages 1
	res, err := http.Get(baseURL) // 웹사이트를 요청한다.

	// goquery가 작동하는 방식
	// 1) Get에서 에러 체크
	// 2) goquery document 만들 때에도 에러 체크
	// (계속해서 에러를 체크해준다)

	// 페이지를 받아오기 위한 에러 체크
	checkErr(err) // request에 에러가 있는지 확인한다.
	checkCode(res)

	defer res.Body.Close() // 함수가 끝났을 때 res.Body를 닫아야 한다. -> 메모리가 새어나가는 걸 막는다.

	// goquery document를 만든다.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	// res.Body를 goquery에게 준다.
	// res.Body - byte, 입력과 출력 (IO)

	checkErr(err)
	fmt.Println(doc)

	// 4-1. getPages 2
	// Find - 모든 페이지를 받아온다.
	// Find the review items (div - class)
	// 찾을 수 있는 모든 각각의 것들에 대해 처리
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html()) // 각 페이지의 HTML이 모두 출력된다.
		pages = s.Find("a").Length() // 링크의 수를 카운트한다.
		// pagination 클래스 안에서 링크를 찾아 링크의 수를 카운트한다.
		// 첫 페이지는 링크(<b>1<b>)가 아니므로 "다음" 링크(첫 페이지 대체)까지 카운트한다.

	})

	return pages

}

// 4-0. getPages 1
func checkErr(err error) { // 에러가 있는지 확인한다.
	if err != nil {
		log.Fatalln(err) // 프로그램 종료
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

// 4-2. extractJob 1
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
	// 1) TrimSpace: 전달한 문자열의 스페이스를 제거한다.
	// 2) Fields: 텍스트로만 이뤄진 배열을 만든다.
	// -> 모든 공백이 제거된(clean) 텍스트로 된 배열을 만든다.
	// ex. "hello     world     " -> "hello","world"
	// 3) join: 배열을 가져와서 합친다. seperator 이용
	// ex. "hello","world" -> hello world
}

// 4-4 writing jobs
// job에 대한 정보를 csv 파일로 저장한다.
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)      // csv에서 한글이 깨지지 않도록 해준다.
	w := csv.NewWriter(file) // 파일(writer) 생성
	// w := csv.NewWriter(os.Stdout) // 실행 시 파일이 생성되지 않고 콘솔에만 결과물이 표시됨
	defer w.Flush() // Flush - 함수가 끝나는 시점에 작성된 모든 데이터를 파일에 입력해 저장한다.
	headers := []string{"Link", "Title", "Location", "Career", "Recruitment period"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs { // 모든 페이지에서 가져온 job 정보가 입력된다.
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.title, job.location, job.career, job.recruitmentPeriod}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

// csv write 부분을 goroutine으로 구현하면 runtime panic에 빠진다.
// -> Go가 General하게 제공하는 csv Write은 goroutine-safe 하지 않다.
// 1) concurrency 작업을 지원하는 라이브러리를 가져다 구현한다.
// 2) Slice를 묶어 한번에 Batch 처리한다. (WriteAll 사용)
// 3) goroutine마다 약간의 딜레이를 준다. (추천 X)
func writeJobs2(jobs []extractedJob) {
	file, err := ccsv.NewCsvWriter("jobs.csv")
	checkErr(err)
	defer file.Close()
	headers := []string{"Link", "Title", "Location", "Career", "Recruitment period"}
	file.Write(headers)
	c := make(chan bool)
	for _, job := range jobs {
		go func(job extractedJob) {
			file.Write([]string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.title, job.location, job.career, job.recruitmentPeriod})
			c <- true
		}(job)
	}
	for i := 0; i < len(jobs); i++ {
		<-c
	}
}

// 4-5. channels time
// 1) 총 페이지 수를 가져온다.
// 2) 각 페이지 별로 goroutine을 생성한다.
// 3) getPage는 각 일자리 정보 별로 goroutine을 생성한다.
//  - 각 페이지에 40개의 일자리 정보가 있고 페이지는 10개가 있으므로 (10 * 40(일자리)) + 10(페이지))개의 goroutine이 만들어진다.
// 일자리 정보를 추출하는 함수는 getPage 함수와 channel을 이용해 정보를 주고 받는다.
//  - getPage 함수는 main 함수와 channel을 이용해 커뮤니케이션 한다.
//  - 즉, 2개의 channel 생성: main <- getPage, getPage <- extractJob
