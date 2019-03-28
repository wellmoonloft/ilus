![](http://image.igerm.cn/img/20190325095035.png)

> ！！！！！！  请注意，本项目未完成  ！！！！！！！！！

> ！！！！！！    多个功能仍未实现  ！！！！！！！！！

> ！！！！！！    开放成公共项目主要为了代码质量审查，请不要下载  ！！！！！！！！！

------------------------------
🇨🇳简体中文 | 🇺🇸[English](README-en_US.md)

## 简介

**ilus** [爱沙尼亚语]，意为靓！

轻快，简洁，功能强大，使用**Go语言**开发的博客系统

本项目的功能定位为博客、CMS及微信小程序后台

由于自己近期一直在研究微信小程序，在开源平台选了很久也没有找到一款合适自己的，正好近期在学习Golang，
所以就基于[sdrms](https://github.com/lhtzbj12/sdrms)这个项目，加上自己的实际需求开发了ilus作为练习项目


## 快速开始

1. 安装go语言开发环境 [Golang v1.12.pkg](https://dl.google.com/go/go1.12.darwin-amd64.pkg)

2. 安装beego：`go get github.com/astaxie/beego`

3. 安装bee工具：`go get github.com/beego/bee`

4. 安装mysql驱动：`go get -u github.com/go-sql-driver/mysql`

5. 安装redis模块：`go get github.com/gomodule/redigo/redis`

6. 下载源码，并在`app.conf`中配置数据库信息（默认不开启redis缓存、可自行配置）

7. 在mysql数据库中创建实例，实例名需与配置一致

8. `run main.go`

## 许可证

[![license](https://img.shields.io/github/license/ruibaby/halo.svg?style=flat-square)](https://github.com/ruibaby/halo/blob/master/LICENSE)

> ilus 使用 GPL-v3.0 协议开源，请尽量遵守开源协议，即便是在中国。

## 感谢

ilus 的诞生离不开下面这些项目：

*前端：*

- [AdminLTE](https://github.com/ColorlibHQ/AdminLTE)：AdminLTE - Free Premium Admin control Panel Theme Based On Bootstrap 3.x
- [Bootstrap](https://github.com/twbs/bootstrap)：The most popular HTML, CSS, and JavaScript framework for developing responsive, mobile first projects on the web
- [Bootstrap-FileInput](https://github.com/kartik-v/bootstrap-fileinput)：An enhanced HTML 5 file input for Bootstrap 3.x with file preview, multiple selection, and more features
- [Bootstrap-Table](https://github.com/wenzhixin/bootstrap-table)：An extended table to integration with some of the most widely used CSS frameworks
- [EasyMDE](https://github.com/Ionaru/easy-markdown-editor)：A simple, beautiful, and embeddable JavaScript Markdown editor
- [Layer](https://github.com/sentsin/layer)：目前layer已经成为国内乃至全世界最多人使用的Web弹层解决方案
- [Font-Awesome](https://github.com/FortAwesome/Font-Awesome)：The iconic SVG, font, and CSS toolkit
- [X-Editable](https://github.com/vitalets/x-editable)：In-place editing with Twitter Bootstrap, jQuery UI or pure jQuery


*后端：*

- [Beego](https://github.com/astaxie/beego)：beego is an open-source, high-performance web framework for the Go programming language 
- [Sdrms](https://github.com/lhtzbj12/sdrms)：SDRMS是基于Beego开发的易用、易扩展、界面友好的轻量级功能权限管理系统
- [Go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)：Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package
- [Redigo](https://github.com/gomodule/redigo)：Go client for Redis




## 代码质量查看
[代码质量](https://goreportcard.com/report/github.com/wellmoonloft/ilus)

