package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	j := &Json{}
	j.GetUrl()
	typename := j.Url
	if strings.Contains(typename, "all") {
		fmt.Println("检测到all类型...开始获取图片")
		j.GetJson()
		j.GetKey()
		j.GetPage()
		j.downloadPic(j.Key)
		for i := 0; i < 3; i++ {
			j.ScrollPageA() //获取翻页后的URL
			j.GetJson()     //从翻页后URL重新获取一次Json数据
			j.GetKey()      //再次获取Key数据
			j.GetPage()     //获取页码
			j.downloadPic(j.Key)
		}
		os.Exit(0)
	}
	if strings.Contains(typename, "boards") {
		fmt.Println("检测到boards类型...开始获取图片")
		j.GetJson()
		j.GetBoardID()
		j.GetBoardPage()
		j.GetBoard()
		j.downloadPic(j.BoardsKey)
		for i := 0; i < 3; i++ {
			j.ScrollPageB()
			j.GetJson()
			j.GetBoardID()
			j.GetBoardPage()
			j.GetBoard()
			j.downloadPic(j.BoardsKey)
		}
		os.Exit(0)
	}
	if strings.Contains(typename, "explore") {
		fmt.Println("检测到explore类型...开始获取图片")
		j.GetJson()
		j.GetExploreName()
		j.GetExplorePage()
		j.GetExplore()
		j.downloadCover(j.ExploreKey)
		for i := 0; i < 3; i++ {
			j.ScrollPageE()
			j.GetJson()
			j.GetExploreName()
			j.GetExplorePage()
			j.GetExplore()
			j.downloadCover(j.ExploreKey)
		}
		os.Exit(0)
	}

	//	j := &Json{}
	//	j.GetUrl()
	//	j.GetJson()
	//	j.GetBoardID()
	//	j.GetBoardPage()
	//	j.GetBoard()
	//	j.downloadPic(j.BoardsKey)
	//	for i := 0; i < 3; i++ {
	//		j.ScrollPageB()
	//		j.GetJson()
	//		j.GetBoardPage()
	//		j.GetBoard()
	//		j.downloadPic(j.BoardsKey)
	//	}
	//	//	j := &Json{}
	//	//	j.GetUrl()
	//	//	j.GetJson()
	//	//	j.GetKey()
	//	//	j.downloadPic(j.Key)

}
