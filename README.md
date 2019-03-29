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


## 快速开始

1. 安装go语言开发环境 [Golang v1.12.pkg](https://dl.google.com/go/go1.12.darwin-amd64.pkg)

2. 安装beego：`go get github.com/astaxie/beego`

3. 安装bee工具：`go get github.com/beego/bee`

4. 安装mysql驱动：`go get -u github.com/go-sql-driver/mysql`

5. 安装redis模块：`go get github.com/gomodule/redigo/redis`

6. 下载源码，并在`app.conf`中配置数据库信息（默认不开启redis缓存、可自行配置）

7. 在mysql数据库中创建实例，实例名需与配置一致

8. `run main.go`

> 系统会自动建表

## 许可证

[![license](https://img.shields.io/github/license/ruibaby/halo.svg?style=flat-square)](https://github.com/ruibaby/halo/blob/master/LICENSE)

> ilus 使用 GPL-v3.0 协议开源，请尽量遵守开源协议，即便是在中国。

## 感谢

ilus 的诞生离不开下面这些项目：

*前端：*

- [AdminLTE.v2.4.10](https://github.com/ColorlibHQ/AdminLTE)：免费的又好看的后台模版
- [Bootstrap-FileInput.v5.0.0](https://github.com/kartik-v/bootstrap-fileinput)：An enhanced HTML 5 file input for Bootstrap 3.x with file preview, multiple selection, and more features
- [Bootstrap-Select.v1.13.9](https://github.com/snapappointments/bootstrap-select)：The jQuery plugin that brings select elements into the 21st century with intuitive multiselection, searching, and much more
- [Bootstrap-Table.v1.11.1](https://github.com/wenzhixin/bootstrap-table)：An extended table to integration with some of the most widely used CSS frameworks
- [EasyMDE.v2.5.1](https://github.com/Ionaru/easy-markdown-editor)：MarkDown编辑器，Fork自SimplyMDE
- [Layer.v3.1.0](https://github.com/sentsin/layer)：国内最流行的弹出插件 
- [X-Editable.v1.5.3](https://github.com/vitalets/x-editable)：In-place editing with Twitter Bootstrap, jQuery UI or pure jQuery


*后端：*

- [Beego.v1.11.1](https://github.com/astaxie/beego)：beego is an open-source, high-performance web framework for the Go programming language 
- [Go-sql-driver/mysql.v1.4.1](https://github.com/go-sql-driver/mysql)：Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package
- [Redigo.v1.7.0](https://github.com/gomodule/redigo)：Go client for Redis




## 代码质量查看
[代码质量](https://goreportcard.com/report/github.com/wellmoonloft/ilus)

