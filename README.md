# A simple but feature-complete command-line dictionary

English | [简体中文](./README-zh_CN.md)

## 1. Install golang

debain

``` bash
sudo apt-get install golang-go
```

macos

``` bash
sudo brew update && brew upgrade && brew install go
```

## 2. Setting the environment variable

### Bash

Edit your ~/.bash_profile to add the following line:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Save and exit your editor. Then, source your ~/.bash_profile.

``` bash
source ~/.bash_profile
```

### Zsh

Edit your ~/.zshrc file to add the following line:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Save and exit your editor. Then, source your ~/.zshrc.

``` bash
source ~/.zshrc
```

### 3. Install the dict

``` bash
go get -u github.com/woorui/dict
```

### 4. Apply appid and secret from baidu and youdao open api

baidu

> <http://api.fanyi.baidu.com/api/trans/product/index>

youdao

> <https://ai.youdao.com/>

### 5. Write your config

``` yaml
-
  key: baidu
  value: YOUR_BAIDU_APPID-YOUR_BAIDU_SECRET
-
  key: youdao
  value: YOUR_YOUDAO_APPKEY-YOUR_YOUDAO_ADDSECRET
```

note. key is api source, value is `strings.Join([]string{APPID, SECRET}, "-"})`

### 6. Running the dict and specify the config file

``` bash
dict -file TYPING_YOUR_CONFIG_FILE_WHEN_FIRST_USING
```

### 7.Show running results

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
