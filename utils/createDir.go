package utils

import (
	"os"
	"tiny-bilibili-downloader/models"
)

func CreateDirAndToDir(data *models.VideoData) {
	downloadDir = data.Bvid
	// 查看文件夹是否存在
	var err error
	_, err = os.Stat(data.Bvid)
	if os.IsNotExist(err) {
		// 不存在则创建文件夹
		_ = os.Mkdir(data.Bvid, os.ModePerm)
	}
	_ = os.Chdir(downloadDir)
}
