![](http://image.igerm.cn/img/20190325095035.png)

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

> ！！！！！！ For Your Attention  ！！！！！

> ！！！！！！ This Project Is Not Completed  ！！！！！

> ！！！！！！ More function is not done  ！！！！！

------------------------------
🇨🇳[简体中文](README.md) | 🇺🇸English

## Introduction

**ilus** [Estonian], means beauty 

Blogging system developed in Golang

Fork from [sdrms](https://github.com/lhtzbj12/sdrms)

Function&UI is draw from [Halo](https://github.com/halo-dev/halo)

The function of this project is located on Blog, CMS and WeChat miniprogram backend

Do not support < IE9

## Quick start

1. Install golang development environment [Golang v1.12.pkg](https://dl.google.com/go/go1.12.darwin-amd64.pkg)

2. Install beego：`go get github.com/astaxie/beego`

3. Install bee tool：`go get github.com/beego/bee`

4. Install yaml.v2：`go get gopkg.in/yaml.v2`

5. Install mysql driver for Golang：`go get -u github.com/go-sql-driver/mysql`

6. Install redis module for Golang：`go get github.com/gomodule/redigo/redis`

7. Clone code, and change [mysql] in `app.conf`（ for default `is_cache=false`, you can make it `true` if you need ）

8. Build a new schema in mysql，schema name must bee same with what you code in `app.conf`

9. `run main.go`

>DB table will be auto build

## License

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

## Dependents



**frontend：**

- [AdminLTE.v2.4.10](https://github.com/ColorlibHQ/AdminLTE)：AdminLTE - Free Premium Admin control Panel Theme Based On Bootstrap 3.x
- [Bootstrap-FileInput.v5.0.0](https://github.com/kartik-v/bootstrap-fileinput)：An enhanced HTML 5 file input for Bootstrap 3.x with file preview, multiple selection, and more features
- [Bootstrap-Select.v1.13.9](https://github.com/snapappointments/bootstrap-select)：The jQuery plugin that brings select elements into the 21st century with intuitive multiselection, searching, and much more
- [Bootstrap-Table.v1.11.1](https://github.com/wenzhixin/bootstrap-table)：An extended table to integration with some of the most widely used CSS frameworks
- [EasyMDE.v2.5.1](https://github.com/Ionaru/easy-markdown-editor)：A simple, beautiful, and embeddable JavaScript Markdown editor
- [jQuery-Validation.v1.19.0](https://github.com/jquery-validation/jquery-validation)：jQuery Validation plugin
- [jQuery-treetable.v3.2.0](https://github.com/ludo/jquery-treetable)：jQuery Dynamic Tree Table Plug-in
- [JavaScript-Cookie.v2.2.0](https://github.com/js-cookie/js-cookie)：Operation Cookie Plug-in
- [Layer.v3.1.0](https://github.com/sentsin/layer)：The most used web pop-up layer solution in China 
- [X-Editable.v1.5.3](https://github.com/vitalets/x-editable)：In-place editing with Twitter Bootstrap, jQuery UI or pure jQuery


**backend：**

- [Beego.v1.11.1](https://github.com/astaxie/beego)：Web framework for the Golang
- [Go-sql-driver/mysql.v1.4.1](https://github.com/go-sql-driver/mysql)：MySQL driver for Golang
- [Redigo.v1.7.0](https://github.com/gomodule/redigo)：Go client for Redis
- [Yaml.v2.2.2](https://github.com/go-yaml/yaml)：YAML support for the Go language

## Code formats levels
[Code formats](https://goreportcard.com/report/github.com/wellmoonloft/ilus)


