package main

import (
	//"net/http"

	"os"

	"github.com/asjasj3964/learngo/5_web_server_with_echo/scrapper"
	"github.com/labstack/echo/v4"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html") // 에코 서버가 템플릿 파일을 응답하도록 한다.
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := scrapper.CleanString(c.FormValue("term"))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName) // 첨부 파일을 반환하는 기능, jobs.csv 파일을 리턴한다.
	// 사용자가 파일을 다운로드 하면 서버에서는 파일을 삭제하고 싶어한다. (요청이 다른데 같은 파일을 저장하는 건 좋지 않다.)
}

func main() {
	e := echo.New()        // 에코를 이용해 서버를 만든다.
	e.GET("/", handleHome) // echo에 url을 설정한다. (url, 함수)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
