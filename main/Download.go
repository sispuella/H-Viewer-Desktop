// Download
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (site *Json) downloadPic(storage []string) {
	fmt.Println("开始缓存图片...")
	//	storage := site.KeyStorage
	var picUrl string
	for i := 0; i < len(storage); i++ {
		picUrl = "http://img.hb.aicdn.com/" + storage[i]
		fmt.Println(picUrl)
		response, err := http.Get(picUrl)
		if err != nil { //重试直到正确获取图片
			i--
			fmt.Println("超时...重试中")
		}
		file, _ := os.Create("./cache/" + storage[i] + ".jpg")
		io.Copy(file, response.Body)
		defer response.Body.Close()
		fmt.Println("图片缓存完成")
	}
	fmt.Println("--------------------------------------------该页缓存完成--------------------------------------------")
}
func (site *Json) downloadCover(storage []string) {
	fmt.Println("开始缓存封面...")
	//	storage := site.KeyStorage
	var picUrl string
	for i := 0; i < len(storage); i++ {
		picUrl = "http://img.hb.aicdn.com/" + storage[i]
		fmt.Println(picUrl)
		response, err := http.Get(picUrl)
		if err != nil { //重试直到正确获取图片
			i--
			fmt.Println("超时...重试中")
		}
		file, _ := os.Create("./cache/cover/" + storage[i] + ".jpg")
		io.Copy(file, response.Body)
		defer response.Body.Close()
		fmt.Println("封面缓存完成")
	}
	fmt.Println("--------------------------------------------该页缓存完成--------------------------------------------")
}
