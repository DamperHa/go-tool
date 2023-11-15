package scratch

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/gocolly/colly/v2"
)

const (
	//url = "https://www.baidu.com/"
	url = "https://wiki.zhiyinlou.com/pages/viewpage.action?pageId=177348620"

	//url      = "https://wiki.zhiyinlou.com/plugins/servlet/mobile?contentId=177348620#content/view/177348620"
	username = "fanzhihao"
	password = "HAOweilai11."
)

func TestWik(t *testing.T) {

	c := colly.NewCollector()

	// 在访问每个链接之前调用的回调函数
	// request在colly中表示一个http请求的所有内容
	c.OnRequest(func(r *colly.Request) {
		authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
		r.Headers.Set("Authorization", authHeader)

		fmt.Println("Visiting", r.URL)
	})

	// 在收到HTTP响应时调用的回调函数
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Received response for", r.Request.URL)
		fmt.Println("Status Code:", r.StatusCode)
		fmt.Println("Body:", string(r.Body))
	})

	// 在访问HTML元素（这里是标题）时调用的回调函数
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Found title:", e.Text)
	})

	// 在访问链接时调用的回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Found link:", link)
	})

	c.Visit(url)
}

//func main() {
//	// 创建一个新的Collector
//	c := colly.NewCollector()
//
//	// 在访问每个链接之前调用的回调函数
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL)
//	})
//
//	// 在访问HTML元素（这里是标题）时调用的回调函数
//	c.OnHTML("title", func(e *colly.HTMLElement) {
//		fmt.Println("Found title:", e.Text)
//	})
//
//	// 在访问链接时调用的回调函数
//	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//		link := e.Attr("href")
//		fmt.Println("Found link:", link)
//	})
//
//	// 访问目标网页
//	err := c.Visit("http://example.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//}
