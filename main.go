package main

import (
	"flag"
	"github.com/CeriChen/tiny-bilibili-downloader/utils"
	"log"
)

var (
	bv  = flag.String("bv", "", "bv号 不用加bv")
	url = flag.String("url", "", "url -url https://www.bilibili.com/video/BV1qE411M7da")
	p   = flag.String("p", "", "多p视频选择，例如1-5集 -p 1-5")
	q   = flag.Int("q", 80, "6:240p 16:360p 32:480p 64:720p \n74:720p60fps 80:1080p 112:1080p高码率 116:1080p60fps")
	s   = flag.String("s", "", "sessionData,登录b站后cookie中查看SESSDATA字段")
	to  = flag.String("to", "./", "视频保存目录")
)

func main() {
	// 解析命令行参数
	flag.Parse()
	// 创建视频下载配置项
	options, err := utils.ParseOptions(*bv, *url, *p, *q, *s, *to)
	//fmt.Println(*options.D)
	if err != nil {
		return
	}
	// 开始下载
	if err := utils.Download(options); err != nil {
		log.Fatal(err)
	}
}
