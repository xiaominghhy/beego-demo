# beego-demo

beego学习demo4
1.  安装go
2.  安装beego
3.  安装bee
1.  bee new ProjectName
2.  bee run 
3.  bee generate

运行bee run的时候程序就会自动构建打包成一个二进制文件。
#### structure
```
beego-demo
├── controller --微服务注册与发现中心
├── admin-server -- 微服务系统管理与调试
├── common -- 公共，配置文件，脚手架等
├── dao -- 数据持久层
├── entity -- 实体类
├── code_platform -- 代码管理平台
├── engine_platform -- 引擎管理平台
└── analysis_platform -- 结果分析平台 
```

参考：https://gitee.com/truthhun/DocHub
#### 学习内容
1. logs 下面是官网提供的两种解决方案，可以配置邮件，发送指定日志级别的日志
    * 文件与行数显示准确 https://beego.me/docs/mvc/controller/logs.md
    * https://beego.me/docs/module/logs.md 
2. router，控制器的学习与使用
3. docker化应用
4. restful
5. vsc调试
6. ORM操作，指定列查询
7. 文件下载操作
8. 文件夹与package name应该是一致的。也就是同一个文件夹中的package应该都是一样的。
9. anyType定义 interface{}
1. & 的含义是什么？有没有好像都没有问题 取地址？

#### plan
* ORM框架
* JWT
* schedule

#### 安装教程

1.  安装go
2.  安装beego
3.  安装bee


url|用途
---|---
user|用户的增删改查
file|文件的上传
download|文件下载

#### docker化使用

镜像构建
```
docker build -t beego .
```
容器运行
```
docker run -rm -p 8080:8080 beego
```
--rm 表示运行之后就直接删除该容器

目前发现一个问题就是界面渲染出现问题，那个返回json数据没有任何问题。

### log

* 2019/10/26:完成go语言得文件下载与上传，下载很简单，上传参考 controller/file.go。下午完成日志输出到文件，环境隔离，配置文件学习
* 2019/10/25:解决查询数据库，参数从前端传输到后端，利用ORM查询数据返回到前端。ORM框架使用初始化自动建表等
* 2019/10/24：解决app-dev.conf文件包含自己导致栈溢出。