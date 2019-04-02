![](http://image.igerm.cn/img/20190325095035.png)

> ！！！！！！  请注意，本项目未完成  ！！！！！！！！！

> ！！！！！！    多个功能仍未实现  ！！！！！！！！！

> ！！！！！！    开放成公共项目主要为了代码质量审查，请不要下载  ！！！！！！！！！

------------------------------
🇨🇳简体中文 | 🇺🇸[English](README-en_US.md)

## 简介

**ilus** 【爱沙尼亚语】，意思是漂亮

使用**Go语言**开发的博客系统，博客部分的功能和外观模仿[Halo](https://github.com/halo-dev/halo)开发

本项目的功能定位为博客、CMS及微信小程序后台【已去掉<IE9的支持】



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

[GNU.v3.0](https://github.com/wellmoonloft/ilus/blob/master/LICENSE)


## 依赖

ilus 依赖的开源项目：

**前端：**

- [AdminLTE.v2.4.10](https://github.com/ColorlibHQ/AdminLTE)：管理后台整合包
- [Bootstrap-FileInput.v5.0.0](https://github.com/kartik-v/bootstrap-fileinput)：图片上传插件
- [Bootstrap-Select.v1.13.9](https://github.com/snapappointments/bootstrap-select)：下拉菜单插件
- [Bootstrap-Table.v1.11.1](https://github.com/wenzhixin/bootstrap-table)：动态表格插件
- [EasyMDE.v2.5.1](https://github.com/Ionaru/easy-markdown-editor)：MarkDown编辑器
- [jQuery-Validation.v1.19.0](https://github.com/jquery-validation/jquery-validation)：内容校验插件
- [jQuery-treetable.v3.2.0](https://github.com/ludo/jquery-treetable)：动态树状表格插件
- [JavaScript-Cookie.v2.2.0](https://github.com/js-cookie/js-cookie)：使用JS操作Cookie的插件
- [Layer.v3.1.0](https://github.com/sentsin/layer)：国内最流行的弹出插件 
- [X-Editable.v1.5.3](https://github.com/vitalets/x-editable)：支持弹出和行内编辑的插件
- [Inline-Attachment.v2.0.3](https://github.com/Rovak/InlineAttachment)：粘贴图片/文件到编辑器内的插件
- [Pagination.v2.1.4](https://github.com/superRaytin/paginationjs)：JS控制的分页插件，支持json数据源


**后端：**

- [Beego.v1.11.1](https://github.com/astaxie/beego)：基于Go语言的web框架
- [Go-sql-driver/mysql.v1.4.1](https://github.com/go-sql-driver/mysql)：基于Go语言的Mysql驱动
- [Redigo.v1.7.0](https://github.com/gomodule/redigo)：基于Go语言的Redis模块
- [Imaging.v1.6.0](https://github.com/disintegration/imaging)：基于Go语言的图片处理工具



## 代码质量查看
[代码质量](https://goreportcard.com/report/github.com/wellmoonloft/ilus)

