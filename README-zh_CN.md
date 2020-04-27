# 简单但功能完备的命令行词典

## 1. 安装 golang

debain

``` bash
sudo apt-get install golang-go
```

macos

``` bash
sudo brew update && brew upgrade && brew install go
```

## 2. 设置$GOPATH的环境变量

### Bash

> 添加下面两行到 ~/.bash_profile 文件:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> 保存并且运行以下命令使其生效.

``` bash
source ~/.bash_profile
```

### - Zsh

> 添加下面两行到 ~/.zshrc 文件:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> 保存并且运行以下命令使其生效.

``` bash
source ~/.zshrc
```

### 3. 安装 dict

``` bash
go get -u github.com/woorui/dict
```

### 4. 申请百度和有道翻译开放api

百度

> <http://api.fanyi.baidu.com/api/trans/product/index>

有道

> <https://ai.youdao.com/>

### 5. 生成一个配置文件

``` yaml
-
  key: baidu
  value: YOUR_BAIDU_APPID-YOUR_BAIDU_SECRET
-
  key: youdao
  value: YOUR_YOUDAO_APPKEY-YOUR_YOUDAO_ADDSECRET
```

注意. value 由代码 `strings.Join([]string{APPID, SECRET}, "-"})` 生成

### 6. 运行dict, 并指定config文件

``` bash
dict -file TYPING_YOUR_CONFIG_FILE_WHEN_FIRST_USING
```

### 7. 运行效果

``` bash
> hello
来源    原文    译文    音标    详情
百度    hello   你好
有道    hello   你好    həˈləʊ  int. 喂；哈罗，你好，您好, n. 表示问候，
                                惊奇或唤起注意时的用语, n. (Hello) 人名；（法）埃洛
> 你好
来源    原文    译文    音标    详情
百度    你好    Hello
有道    你好    hello           hello, hi, how do you do
```
