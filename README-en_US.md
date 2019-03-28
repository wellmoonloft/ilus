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

## Quick start

1. Install golang development environment [Golang v1.12.pkg](https://dl.google.com/go/go1.12.darwin-amd64.pkg)

2. Install beego：`go get github.com/astaxie/beego`

3. Install bee tool：`go get github.com/beego/bee`

4. Install mysql driver for Golang：`go get -u github.com/go-sql-driver/mysql`

5. Install redis module for Golang：`go get github.com/gomodule/redigo/redis`

6. Clone code, and change [mysql] in `app.conf`（ for default `is_cache=false`, you can make it `true` if you need ）

7. Build a new schema in mysql，schema name must bee same with what you code in `app.conf`

8. `run main.go`

>DB table will be auto build

## License

[![license](https://img.shields.io/github/license/ruibaby/halo.svg?style=flat-square)](https://github.com/ruibaby/halo/blob/master/LICENSE)

> ilus uses the GPL-v3.0 protocol to open source

## Thanks

The birth of Halo is inseparable from the following projects：

*frontend：*

- [AdminLTE](https://github.com/ColorlibHQ/AdminLTE)：AdminLTE - Free Premium Admin control Panel Theme Based On Bootstrap 3.x
- [Bootstrap](https://github.com/twbs/bootstrap)：The most popular HTML, CSS, and JavaScript framework for developing responsive, mobile first projects on the web
- [Bootstrap-FileInput](https://github.com/kartik-v/bootstrap-fileinput)：An enhanced HTML 5 file input for Bootstrap 3.x with file preview, multiple selection, and more features
- [Bootstrap-Table](https://github.com/wenzhixin/bootstrap-table)：An extended table to integration with some of the most widely used CSS frameworks
- [EasyMDE](https://github.com/Ionaru/easy-markdown-editor)：A simple, beautiful, and embeddable JavaScript Markdown editor
- [Layer](https://github.com/sentsin/layer)：The most used web pop-up layer solution in China 
- [Font-Awesome](https://github.com/FortAwesome/Font-Awesome)：The iconic SVG, font, and CSS toolkit
- [X-Editable](https://github.com/vitalets/x-editable)：In-place editing with Twitter Bootstrap, jQuery UI or pure jQuery


*backend：*

- [Beego](https://github.com/astaxie/beego)：beego is an open-source, high-performance web framework for the Go programming language 
- [Go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)：Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package
- [Redigo](https://github.com/gomodule/redigo)：Go client for Redis

## Code formats levels
[Code formats](https://goreportcard.com/report/github.com/wellmoonloft/ilus)


