## 一个简单但是功能完备的命令行词典

### 1. 安装 golang

+ debain
```
$ sudo apt-get install golang-go
```

+ macos
```
$ sudo brew update && brew upgrade && brew install go
```

### 2. 安装 simple-dict

```
$ go get -u github.com/qq1009479218/simple-dict
```

### 3. 申请百度翻译开放api

> http://api.fanyi.baidu.com/api/trans/product/index

### 4. 运行 simple-dict 并且填写 appid 和 secret

```
$ simple-dict
You need baidu secret and appid, apply link: http://api.fanyi.baidu.com/api/trans/product/index // 回车

Input your baidu appid
`YOUR_BAIDU_APPID` // 回车
Input your baidu secret
`YOUR_BAIDU_SECRET` // 回车
```

### 5. 运行效果
```
Input words to translate
word
|> word: 单词
multiple
|> multiple: 倍数
```