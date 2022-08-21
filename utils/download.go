package utils

import (
	"fmt"
	"github.com/CeriChen/tiny-bilibili-downloader/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	videoQuality = map[int]string{
		6:   "240p",
		16:  "360p",
		32:  "480p",
		64:  "720p",
		74:  "720p60fps",
		80:  "1080p",
		112: "1080p high bit rate",
		116: "1080p 60fps",
	}

	B  = 1
	KB = B << 10
	MB = KB << 10

	token chan interface{} //令牌桶

	allFiles, curFiles int

	mutex sync.Mutex
)

func Download(options *models.DownloadOptions) (err error) {
	// 创建令牌桶
	CreateToken(min(options.PageChoice[1]-options.PageChoice[0]+1, 3))
	// 获取视频下载信息
	if err := GetVideoDownloadInfo(options); err != nil {
		return err
	}
	return DownloadVideo(options)
}

func DownloadVideo(options *models.DownloadOptions) (err error) {
	if options.Bv == "" {
		fmt.Println("need bv id.")
		return
	}
	if err = CreateDirAndToDir(options); err != nil {
		fmt.Println("create dir failed.")
		return
	}

	if options.PageChoice == nil {
		allFiles = 1
		fmt.Printf("[START DOWNLOADING] ALL %d VIDEO(S).\n", allFiles)
		err = DownloadOneVideo(options.VD.Title, options, options.DD)
		return
	}
	allFiles = options.PageChoice[1] - options.PageChoice[0] + 1
	fmt.Printf("[START DOWNLOADING] ALL %d VIDEO(S).\n", allFiles)
	err = DownloadMultiVideo(options)
	return
}

// DownloadOneVideo 下载单一视频
func DownloadOneVideo(name string, options *models.DownloadOptions, dd *models.DownloadData) error {
	// 处理文件格式
	suffix := options.DD.Format
	if strings.Contains(suffix, "flv") {
		suffix = "flv"
	}
	fileName := fmt.Sprintf("%s.%s", name, suffix)
	// 设置请求头
	req, err := http.NewRequest(http.MethodGet, options.DD.Durl[0].Url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip,deflate,br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Referer", "https://api.bilibili.com/")

	// 读取视频流
	all, err := options.Cli.Do(req)
	if all.StatusCode != http.StatusOK {
		fmt.Printf("status:%d,err:%v\n", all.StatusCode, err)
		return err
	}
	if err != nil {
		return err
	}
	defer all.Body.Close()
	// 创建目标视频文件
	var file *os.File
	defer file.Close()
	file, err = os.Create(fileName)
	if os.IsExist(err) {
		err = nil
		file, err = os.Create(options.Bv + ".flv")
		if err != nil {
			return err
		}
	}
	duration, _ := time.ParseDuration(strconv.Itoa(dd.TimeLength) + "ms")
	fmt.Printf("[TASK] Start downloading: '%s' -> '%s' <%s| %dMB |%s>\n",
		name,
		file.Name(),
		videoQuality[options.DD.Quality],
		dd.Durl[0].Size/MB,
		duration.String(),
	)
	mutex.Unlock()
	// 文件拷贝
	_, err = io.Copy(file, all.Body)
	if err != nil {
		fmt.Printf("save as video failed..err:%v\n", err)
		return err
	}
	curFiles++
	fmt.Printf("[COMPLETED %d/%d] '%s' -> %s\n", curFiles, allFiles, name, file.Name())
	return err
}

// DownloadMultiVideo 多线程下载多p
func DownloadMultiVideo(options *models.DownloadOptions) error {
	ok := make(chan interface{}, options.PageChoice[1]-options.PageChoice[0]+1)
	for _, p := range options.VD.Pages {
		p := p
		if p.Page < options.PageChoice[0] || p.Page > options.PageChoice[1] {
			continue
		}
		select {
		case <-token:
			go func() {
				mutex.Lock()
				options.VD.Cid = p.Cid
				err := GetVideoDownloadInfo(options)
				if err != nil {
					panic(err)
				}
				var namePrefix = "[P" + strconv.Itoa(p.Page) + "]"
				err = DownloadOneVideo(namePrefix+p.Part, options, options.DD)
				if err != nil {
					panic(err)
				}
				token <- 1
				ok <- 1
			}()
		}
	}
	for i := 0; i < options.PageChoice[1]-options.PageChoice[0]+1; i++ {
		<-ok
	}
	fmt.Println("♥all download tasks were done!")
	return nil
}

// CreateToken 创建令牌桶
func CreateToken(t int) {
	token = make(chan interface{}, t)
	for i := 0; i < t; i++ {
		token <- 1
	}
}

func min(i, c int) int {
	if i < c {
		return i
	}
	return c
}
