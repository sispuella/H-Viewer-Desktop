package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	//"log"
	"bufio"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	//"github.com/tidwall/gjson"
)

type Json struct {
	//通用
	SiteResponse http.Response
	Url          string
	SiteJson     string
	KeyStorage   []string
	//画板
	BoardsKey   []string
	Boardspinid []string
	BoardsPage  []string
	BoardsCover string
	BoardID     string
	//首页
	Key   []string
	Pinid []string
	Page  []string
	//发现
	ExploreKey   []string
	ExplorePage  []string
	ExploreCover string
	ExploreName  string
}

func (site *Json) GetUrl() {
	urlType := bufio.NewReader(os.Stdin)
	var input string
	fmt.Printf("输入URL:")
	input, _ = urlType.ReadString('\n')
	site.Url = strings.TrimSpace(input)
}

//func (address Json) Url() {
//	address.Url
//}
func (site *Json) GetJson() {
	fmt.Printf("获取请求...")
	client := &http.Client{Timeout: 10 * time.Second}
	address := site.Url
	//	fmt.Println("$$$$$$$$$", address)
	var passedResponse *http.Response
	for retry := 0; retry < 3; retry++ { //请求重试次数设定
		request, _ := http.NewRequest("GET", address, nil)
		request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:60.0) Gecko/20100101 Firefox/60.0")
		request.Header.Set("X-Request", "JSON")
		request.Header.Set("X-Requested-With", "XMLHttpRequest")
		response, err := client.Do(request)
		if err == nil { //成功
			passedResponse = response
			fmt.Println("成功")

			break
		}
		if retry == 2 {
			fmt.Println("超时...")
			os.Exit(0)
		}
	}
	var json string
	fmt.Printf("获取json数据...")
	for {
		if passedResponse.StatusCode == 200 {
			dom, err := goquery.NewDocumentFromResponse(passedResponse)
			if err != nil {
				log.Fatal("请求错误：", passedResponse.StatusCode)
			}
			dom.Find("html").Each(func(i int, content *goquery.Selection) {
				json = content.Find("body").Text()
				site.SiteJson = json
			})
			fmt.Println("成功", site.SiteJson, "$$$$$$$$$$")
		}
		break
	}
}

func (site *Json) GetKey() { //从首页中获取图片信息
	json := site.SiteJson
	fmt.Printf("解析key数据...")
	storage := make([]string, 0, 100)
	key := gjson.Get(json, "pins.#.file.key")
	for _, result := range key.Array() {
		storage = append(storage, result.String())
	}
	site.Key = storage
	fmt.Println("成功")
	//	fmt.Println("############", site.Key)
}
func (site *Json) GetPin() {
	json := site.SiteJson
	fmt.Printf("解析pinid...")
	storage := make([]string, 0, 100)
	pin_id := gjson.Get(json, "pins.19.pin_id")
	for _, result := range pin_id.Array() {
		storage = append(storage, result.String())
	}
	site.Pinid = storage
	fmt.Println("成功")
	fmt.Println("----------------------------------------------------------------------------------------")
}
func (site *Json) GetBoard() { //获取当前huaban.com/boards/xxxxxx/页面下的所有图片Key
	json := site.SiteJson
	fmt.Printf("解析key数据...")
	storage := make([]string, 0, 100)
	key := gjson.Get(json, "board.pins.#.file.key")
	for _, result := range key.Array() {
		storage = append(storage, result.String())
	}
	site.BoardsKey = storage
	fmt.Println("成功")
}
func (site *Json) GetBoardPage() { //获取当前huaban.com/boards/xxxxxx/页面下最后一张图片pinid以实现翻页
	json := site.SiteJson
	fmt.Printf("获取页码...")
	storage := make([]string, 0, 1)
	key := gjson.Get(json, "board.pins.19.pin_id")
	for _, result := range key.Array() {
		storage = append(storage, result.String())
	}
	fmt.Println(storage)
	site.BoardsPage = storage
	fmt.Println("成功")
}
func (site *Json) GetBoardID() { //获取当前huaban.com/boards/xxxxxx/中的XXXXXX
	json := site.SiteJson
	boardid := gjson.Get(json, "board.board_id")
	site.BoardID = boardid.String()
	fmt.Println("获取Board ID成功")
}
func (site *Json) ScrollPageB() { //翻页
	var url string
	boardID := site.BoardID
	page := site.BoardsPage
	fmt.Println(page)
	url = "http://huaban.com/boards/" + boardID + "/?jhsic5tc&max=" + page[0] + "&limit=20&wfl=1"
	site.Url = url
	fmt.Println("!!!!!!!!!!", site.Url, "!!!!!!!!!")
}
func (site *Json) GetExplore() { //获取当前huaban.com/explore/xxxxxx/页面下的所有图片Key
	json := site.SiteJson
	fmt.Printf("解析key数据...")
	storage := make([]string, 0, 100)
	key := gjson.Get(json, "pins.#.file.key")
	for _, result := range key.Array() {
		storage = append(storage, result.String())
	}
	site.ExploreKey = storage
	fmt.Println(site.ExploreKey)
	fmt.Println("成功")
}
func (site *Json) GetExplorePage() { //获取当前huaban.com/explore/xxxxxx/页面下最后一张图片pinid以实现翻页
	json := site.SiteJson
	fmt.Printf("获取页码...")
	storage := make([]string, 0, 1)
	key := gjson.Get(json, "pins.19.pin_id")
	for _, result := range key.Array() {
		storage = append(storage, result.String())
	}
	//fmt.Println(storage)
	site.ExplorePage = storage
	fmt.Println("成功")
}
func (site *Json) GetExploreName() { //获取当前huaban.com/explore/xxxxxx/中的XXXXXX
	json := site.SiteJson
	exploreName := gjson.Get(json, "urlname")
	site.ExploreName = exploreName.String()
	fmt.Println("获取Board ID成功")
}
func (site *Json) ScrollPageE() { //翻页
	var url string
	exploreName := site.ExploreName
	page := site.ExplorePage
	fmt.Println(page)
	url = "http://huaban.com/explore/" + exploreName + "?&max=" + page[0]
	site.Url = url
	fmt.Println("!!!!!!!!!!", site.Url, "!!!!!!!!!")
}
func (site *Json) GetPage() { //获取当前huaban.com/all页面下最后一张图片pinid以实现翻页
	json := site.SiteJson
	fmt.Printf("获取页码...")
	storage := make([]string, 0, 1)
	key := gjson.Get(json, "pins.19.pin_id")
	for _, result := range key.Array() {
		storage = append(storage, result.String())
	}
	//fmt.Println(storage)
	site.Page = storage
	fmt.Println("成功")
}
func (site *Json) ScrollPageA() { //翻页
	var url string
	page := site.Page
	fmt.Println(page)
	url = "http://huaban.com/all/" + "/?jhsic5tc&max=" + page[0] + "&limit=20&wfl=1"
	site.Url = url
	fmt.Println("!!!!!!!!!!", site.Url, "!!!!!!!!!")
}

//func (site Json) PassKey() []string {
//	site.KeyStorage = site.Key
//	return site.KeyStorage
//}
