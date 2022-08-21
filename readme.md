# bilibili命令行视频下载工具

## 仅限学习与交流！！！

### 基本使用

单p视频下载(url) -url
  `$ BilibiliDownloader -url="https://www.bilibili.com/video/BV1GV411k7UX"`  
单p视频下载(bv号) -bv
  `BilibiliDownloader -bv="BV1GV411k7UX"`  
多p视频下载(url) -p
  `BilibiliDownloader -url="https://www.bilibili.com/video/BV1GV411k7UX" -p="1-9"`  
多p视频下载(bv号)
  `BilibiliDownloader -bv="BV1GV411k7UX" -p="1-8"`  
选择视频清晰度 -q 80 `BilibiliDownloader -bv="BV1GV411k7UX" -q=80`
```text
以下是清晰度参数
6:240p 16:360p 32:480p 64:720p
74:720p60fps 80:1080p 112:1080p
116:1080p60fps 
默认为80
```
>sessData(获取1080p以上视频需要添加此参数) -s xxxxxxxxx
登录b站后cookie里复制下就行


### 详细 BibiliDownloader -help
```text
  -bv string
        视频bv号，如果已添加url则不需要
  -p string
        多p视频选择，例如1-5集 -p="1-5"
  -path string
        指定视频保存目录 (default "./")
  -q int
        6:240p 16:360p 32:480p 64:720p 
        74:720p60fps 80:1080p 112:1080p高码率 116:1080p60fps (default 80)
  -s string
        sessionData，登录b站后cookie中查看SESSDATA字段，可下载高清及以上视频
  -url string
        例如 https://www.bilibili.com/video/BV1qE411M7da

```