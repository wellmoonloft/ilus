![](http://image.igerm.cn/img/20190325095035.png)

> ！！！！！！ For Your Attention  ！！！！！

> ！！！！！！ This Project Is Not Completed  ！！！！！

> ！！！！！！ More function is not done  ！！！！！

------------------------------
🇨🇳[简体中文](README.md) | 🇺🇸English

## Introduction

**ilus** [Estonian], means beauty. 

Fast, concise, and powerful blogging system developed in Golang

The function of this project is located on Blog, CMS and WeChat miniprogram backend. 

The reason of I developed this, it is because I have been developing Wechat miniprogram recently, and I haven't found a 
suitable one on the open source platform for a long time. 

SO, I developed Inus an exercise project based on [sdrms](https://github.com/lhtzbj12/sdrms) use Golang.

## Quick start

1. Install golang development environment [Golang v1.12.pkg](https://dl.google.com/go/go1.12.darwin-amd64.pkg)

2. Install beego：`go get github.com/astaxie/beego`

3. Install bee tool：`go get github.com/beego/bee`

4. Install mysql driver for Golang：`go get -u github.com/go-sql-driver/mysql`

5. Install redis module for Golang：`go get github.com/gomodule/redigo/redis`

6. Clone code, and change [mysql] in `app.conf`（ for default `is_cache=false`, you can make it `true` if you need ）

7. Build a new schema in mysql，schema name must bee same with what you code in `app.conf`

8. `run main.go`



## License

[![license](https://img.shields.io/github/license/ruibaby/halo.svg?style=flat-square)](https://github.com/ruibaby/halo/blob/master/LICENSE)

> ilus uses the GPL-v3.0 protocol to open source

## Thanks

The birth of Halo is inseparable from the following projects：

- [beego](https://github.com/astaxie/beego)：beego is an open-source, high-performance web framework for the Go programming language
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)：Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package
- [Redigo](https://github.com/gomodule/redigo)：Go client for Redis
- [Halo](https://github.com/halo-dev/halo)：Halo Become the best blogging system using Java
- [EasyMDE](https://github.com/Ionaru/easy-markdown-editor)：A simple, beautiful, and embeddable JavaScript Markdown editor.
- [sdrms](https://github.com/lhtzbj12/sdrms)：SDRMS is a lightweight functional privilege management system based on Beego, which is easy to use, easy to expand and user-friendly.


## Code formats levels
[Code formats](https://goreportcard.com/report/github.com/lhtzbj12/sdrms)


