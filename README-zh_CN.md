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

### 2. 设置$GOPATH的环境变量

#### - Bash

> 添加下面两行到 ~/.bash_profile 文件:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> 保存并且运行以下命令使其生效.

``` bash
source ~/.bash_profile
```

#### - Zsh

> 添加下面两行到 ~/.zshrc 文件:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> 保存并且运行以下命令使其生效.

``` bash
source ~/.zshrc
```

### 3. 安装 simple-dict

``` bash
$ go get -u github.com/qq1009479218/simple-dict
```

### 4. 申请百度翻译开放api

> http://api.fanyi.baidu.com/api/trans/product/index

### 5. 运行 simple-dict 并且填写 appid 和 secret

``` bash
$ simple-dict
You need baidu secret and appid, apply link: http://api.fanyi.baidu.com/api/trans/product/index // 回车

Input your baidu appid
`YOUR_BAIDU_APPID` // 回车
Input your baidu secret
`YOUR_BAIDU_SECRET` // 回车
```

### 6. 运行效果
``` bash
Input words to translate
word
|> word: 单词
multiple
|> multiple: 倍数
```