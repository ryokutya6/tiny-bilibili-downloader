package utils

import (
	"fmt"
	"github.com/CeriChen/tiny-bilibili-downloader/models"
	"net/http"
	"strconv"
	"strings"
)

// ParseOptions 解析视频下载配置
func ParseOptions(bv string, url string, p string, q int, sd string, path string) (options *models.DownloadOptions, err error) {
	options = &models.DownloadOptions{
		Cli:         http.DefaultClient,
		Bv:          bv,
		Url:         url,
		PageChoice:  []int{1, 1},
		Quality:     q,
		SessionData: sd,
		SavePath:    path,
	}
	// 从bv号或者url中获取视频信息
	if err = GetVideoInfo(options); err != nil {
		return options, err
	}
	// 解析多p
	if p != "" {
		index := strings.Index(p, "-")
		firstPage, err := strconv.Atoi(p[:index])
		if err != nil {
			fmt.Printf("strconv.Atoi(p[:index]) failed..err: %v", err)
			return nil, err
		}
		endPage, err := strconv.Atoi(p[index+1:])
		if err != nil {
			fmt.Printf("strconv.Atoi(p[index+1:]) failed..err: %v", err)
			return nil, err
		}
		options.PageChoice = []int{firstPage, endPage}
	}

	return options, nil
}
