# bilibili命令行视频下载工具

## 仅限学习与交流！！！

### 基本使用

+ 单p视频下载(url) -url
  `$ BilibiliDownloader -url https://www.bilibili.com/video/BV1GV411k7UX?spm_id_from=333.337.search-card.all.click`
+ 单p视频下载(bv号) -bv
  `BilibiliDownloader -bv BV1GV411k7UX`
+ 多p视频下载(url)--使用'-'表示范围 -url xx -p x-x
  `BilibiliDownloader -url https://www.bilibili.com/video/BV1GV411k7UX?spm_id_from=333.337.search-card.all.click -p 1-9`
+ 多p视频下载(bv号) -bv -p x-x
  `BilibiliDownloader -bv BV1GV411k7UX -p 1-8`
+ 选择清晰度 -q 80

```text
        6:240p 16:360p 32:480p 64:720p
        74:720p60fps 80:1080p 112:1080p
        高码率 116:1080p60fps 
        (default 80)
```

`BilibiliDownloader -bv BV1GV411k7UX -p 1-8 -q 80`

+ sessData(获取1080p以上需要添加此参数) -s xxxxxxxxx
登录b站后cookie里复制下就行


### 详细 BibiliDownloader -help
