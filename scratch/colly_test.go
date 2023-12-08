package scratch

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/gocolly/colly/v2"
)

const (
	//url = "https://www.baidu.com/"
	url = "https://wiki.zhiyinlou.com/pages/viewpage.action?pageId=168697370"

	shimo = "https://yach-doc-shimo.zhiyinlou.com/docs/Ee32MaOMaESQvNA2/ <分组优化服务端技术文档>"

	//url      = "https://wiki.zhiyinlou.com/plugins/servlet/mobile?contentId=177348620#content/view/177348620"
	username = "fanzhihao"
	password = "...."
)

func TestWik(t *testing.T) {

	c := colly.NewCollector()

	// 在访问每个链接之前调用的回调函数
	// request在colly中表示一个http请求的所有内容
	c.OnRequest(func(r *colly.Request) {
		authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
		r.Headers.Set("Authorization", authHeader)
	})

	// 在收到HTTP响应时调用的回调函数
	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("Received response for", r.Request.URL)
	//	fmt.Println("Status Code:", r.StatusCode)
	//	fmt.Println("Body:", string(r.Body))
	//})

	// 在访问HTML元素（这里是标题）时调用的回调函数
	//c.OnHTML("title", func(e *colly.HTMLElement) {
	//	fmt.Println("Found title:", e.Text)
	//})

	// 在访问链接时调用的回调函数
	// 只有一个Head元素
	// 只有一个body元素
	// h1:
	//Found: 版本修订说明
	//Found: 一. 项目概述
	//Found: 二. 项目价值
	//Found: 三. 功能需求
	//Found: 研发问题沟通结论：

	//data := make(map[string]string)
	//c.OnHTML("div.table-wrap", func(e *colly.HTMLElement) {
	//	if e.Text == "" {
	//		return
	//	}
	//
	//	tableHTML := e.
	//
	//	// 将 HTML 内容转换成 Markdown 格式
	//	markdown := HTMLToMarkdown(tableHTML)
	//
	//	// 为什么不能使用md的形式展示表格呢
	//	fmt.Println("Found:", e.Text)
	//
	//	// 提取标题文本
	//	//title := e.Text
	//	//// 初始化用于存储 <p> 元素文本的切片
	//	//var paragraphs []string
	//	//// 查找与当前标题相邻的所有 <p> 元素
	//	//e.ForEach("p", func(i int, elem *colly.HTMLElement) {
	//	//	// 提取每个 <p> 元素的文本并添加到切片中
	//	//	paragraphs = append(paragraphs, elem.Text)
	//	//})
	//	//// 将 <p> 元素的文本拼接成一个字符串，以换行符分隔
	//	//content := strings.Join(paragraphs, "。")
	//	//// 将标题和内容添加到 map 中
	//	//
	//	//if content != "" {
	//	//	data[title] = content
	//	//}
	//
	//	//fmt.Println("Found:", e.Text)
	//})

	// 访问 URL 并找到表格
	c.OnHTML("body", func(e *colly.HTMLElement) {

		// 找到第一个元素
		// 在这个元素中，遍历所有的p元素

		//fmt.Println("table start")
		fmt.Println(e.Text)
		//fmt.Println(e.ChildTexts("p"))

		//resP := e.ChildTexts("body")
		//fmt.Println(resP)

		//tableHtml, _ := e.DOM.Html() // 获取表格的 HTML

		//fmt.Println(tableHtml)
	})

	c.Visit(url)
	//fmt.Println(len(data))
	//fmt.Println(data)
	//
	//for key, d := range data {
	//	fmt.Println("line1:", key, ":", d)
	//}
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

func TestShimo(t *testing.T) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
		r.Headers.Set("Authorization", authHeader)

		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("", func(e *colly.HTMLElement) {

		// 找到第一个元素
		// 在这个元素中，遍历所有的p元素

		//fmt.Println("table start")
		//fmt.Println(e.Text)
		//fmt.Println(e.ChildTexts("p"))

		resP := e.ChildTexts("p")
		fmt.Println(resP)

		//tableHtml, _ := e.DOM.Html() // 获取表格的 HTML

		//fmt.Println(tableHtml)
	})

	c.Visit(shimo)
}
