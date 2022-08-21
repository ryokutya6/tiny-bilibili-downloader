package utils

import (
	"encoding/json"
	"github.com/CeriChen/tiny-bilibili-downloader/models"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetVideoInfo 获取视频信息，输入bv号或者url
func GetVideoInfo(options *models.DownloadOptions) (err error) {
	if options.Url != "" {
		s := strings.Index(options.Url, "BV")
		options.Bv = options.Url[s+2 : s+12]
	}
	path := `https://api.bilibili.com/x/web-interface/view?bvid=` + options.Bv
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, path, nil); err != nil {
		return err
	}
	var resp *http.Response
	if resp, err = options.Cli.Do(req); err != nil {
		return err
	}
	body := resp.Body
	defer body.Close()
	var all []byte
	if all, err = ioutil.ReadAll(body); err != nil {
		return err
	}
	var vi models.VideoInfo
	err = json.Unmarshal(all, &vi)
	if err != nil {
		return err
	}
	options.VD = vi.Data
	return
}
