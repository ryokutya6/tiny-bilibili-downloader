package utils

import (
	"encoding/json"
	"fmt"
	"github.com/CeriChen/tiny-bilibili-downloader/models"
	"io/ioutil"
	"net/http"
)

// GetVideoDownloadInfo 获取下载信息--高清晰视频需要输入sessionData配置项
func GetVideoDownloadInfo(options *models.DownloadOptions) (err error) {
	var req *http.Request
	req, err = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&qn=%d",
			options.Bv,
			options.VD.Cid,
			options.Quality,
		),
		nil)
	if err != nil {
		fmt.Printf("http.NewRequest() failed..err: %v", err)
		return
	}
	req.AddCookie(&http.Cookie{
		Name:  "SESSDATA",
		Value: options.SessionData,
	})
	var resp *http.Response
	resp, err = options.Cli.Do(req)
	if err != nil {
		return
	}
	body := resp.Body
	defer body.Close()
	all, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	var info models.BiliVideoDownloadInfo
	err = json.Unmarshal(all, &info)
	if err != nil {
		return err
	}
	options.DD = info.Data
	return err
}
