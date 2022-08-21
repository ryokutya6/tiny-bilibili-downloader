package utils

import (
	"fmt"
	"github.com/CeriChen/tiny-bilibili-downloader/models"
	"os"
)

func CreateDirAndToDir(options *models.DownloadOptions) (err error) {
	options.SavePath = fmt.Sprintf("%s/%s", options.SavePath, options.VD.Bvid)
	// 查看文件夹是否存在
	_, err = os.Stat(options.SavePath)
	if os.IsNotExist(err) {
		// 不存在则创建文件夹
		if err = os.Mkdir(options.SavePath, os.ModePerm); err != nil {
			fmt.Println("文件夹创建失败，请检查路径是否存在。")
			return
		}
	}
	_ = os.Chdir(options.SavePath)
	return
}
