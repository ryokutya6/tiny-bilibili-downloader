package models

import "net/http"

// BiliVideoDownloadInfo 视频下载信息
type BiliVideoDownloadInfo struct {
	Code    int           `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Ttl     int8          `json:"ttl,omitempty"`
	Data    *DownloadData `json:"data,omitempty"`
}

type DownloadData struct {
	From              string         `json:"from,omitempty"`
	Result            string         `json:"result,omitempty"`
	Message           string         `json:"message,omitempty"`
	Quality           int            `json:"quality,omitempty"`
	Format            string         `json:"format,omitempty"`
	TimeLength        int            `json:"timelength,omitempty"`
	AcceptFormat      string         `json:"accept_format,omitempty"`
	AcceptDescription []string       `json:"accept_description,omitempty"`
	AcceptQuality     []int          `json:"accept_quality,omitempty"`
	VideoCodecid      int            `json:"video_codecid,omitempty"`
	Durl              []*DurlItem    `json:"durl,omitempty"`
	SupportFormats    []*VideoFormat `json:"support_formats,omitempty"`
	// TODO dash
	//Dash []*DashItem
}

type DurlItem struct {
	Order     int      `json:"order,omitempty"`
	Length    int      `json:"length,omitempty"`
	Size      int      `json:"size,omitempty"`
	Ahead     string   `json:"ahead,omitempty"`
	Vhead     string   `json:"vhead,omitempty"`
	Url       string   `json:"url,omitempty"`
	BackupUrl []string `json:"backup_url,omitempty"`
}

type VideoFormat struct {
	Codecs         interface{} `json:"codecs,omitempty"`
	Quality        int         `json:"quality,omitempty"`
	Format         string      `json:"format,omitempty"`
	NewDescription string      `json:"new_description,omitempty"`
	DisplayDesc    string      `json:"display_desc,omitempty"`
	Superscript    string      `json:"superscript,omitempty"`
}

// VideoInfo 视频基本信息
type VideoInfo struct {
	Code    int
	Ttl     int
	Message string
	Data    *VideoData
}

type VideoData struct {
	Bvid      string `json:"bvid,omitempty"`
	Aid       int    `json:"aid,omitempty"`
	Videos    int    `json:"videos,omitempty"`    // 稿件分p数
	Tid       int    `json:"tid,omitempty"`       // 分区tid
	Tname     string `json:"tname,omitempty"`     // 子分区名称
	Copyright int8   `json:"copyright,omitempty"` // 原创1 转载2
	Pic       string `json:"pic,omitempty"`       // 稿件封面图片
	Title     string `json:"title,omitempty"`     // 标题
	Cid       int    `json:"cid,omitempty"`       // 当前视频cid
	Pages     []*Page
	// TODO 还没写完
}

type Page struct {
	Cid      int    `json:"cid,omitempty"`
	Page     int    `json:"page,omitempty"`
	Part     string `json:"part,omitempty"`     // 当前分p标题
	Duration int    `json:"duration,omitempty"` // 时间
}

// DownloadOptions 视频下载配置
type DownloadOptions struct {
	Quality     int
	Bv          string
	Url         string
	Cli         *http.Client
	PageChoice  []int
	SessionData string
	SavePath    string
	DD          *DownloadData
	VD          *VideoData
}
