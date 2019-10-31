## a simple but feature-complete command-line dictionary

English | [简体中文](./README-zh_CN.md)

### 1. Install golang

+ debain
``` bash
$ sudo apt-get install golang-go
```

+ macos
``` bash
$ sudo brew update && brew upgrade && brew install go
```

### 2. Setting the environment variable

#### - Bash

> Edit your ~/.bash_profile to add the following line:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> Save and exit your editor. Then, source your ~/.bash_profile.

``` bash
source ~/.bash_profile
```

#### - Zsh

> Edit your ~/.zshrc file to add the following line:

``` bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

> Save and exit your editor. Then, source your ~/.zshrc.

``` bash
source ~/.zshrc
```

### 3. Install the dict

```
$ go get -u github.com/woorui/dict
```

### 4. Apply appid and secret from baidu open api

> http://api.fanyi.baidu.com/api/trans/product/index

### 5. Running the dict and save your appid and secret

```
$ dict
You need baidu secret and appid, apply link: http://api.fanyi.baidu.com/api/trans/product/index // Enter

Input your baidu appid
`YOUR_BAIDU_APPID` // Enter
Input your baidu secret
`YOUR_BAIDU_SECRET` // Enter
```

### 6.Show running results
```
word
> word
 word: 单词
> 单词
 单词: Word
> 
```