## a simple but feature-complete command-line dictionary

English | [简体中文](./README-zh_CN.md)

### 1. install golang

+ debain
```
$ sudo apt-get install golang-go
```

+ macos
```
$ sudo brew update && brew upgrade && brew install go
```

### 2. install the simple-dict

```
$ go get -u github.com/qq1009479218/simple-dict
```

### 3. apply appid and secret from baidu open api

> http://api.fanyi.baidu.com/api/trans/product/index

### 4. run simple-dict and save your appid and secret

```
$ simple-dict
You need baidu secret and appid, apply link: http://api.fanyi.baidu.com/api/trans/product/index // Enter

Input your baidu appid
`YOUR_BAIDU_APPID` // Enter
Input your baidu secret
`YOUR_BAIDU_SECRET` // Enter
```

### 5. show running results
```
Input words to translate
word
|> word: 单词
multiple
|> multiple: 倍数
```