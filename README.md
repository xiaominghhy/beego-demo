# beego-demo

#### 介绍
beego学习demo

#### 学习内容
1. logs
2. router，控制器的学习与使用
3. docker化应用
4. restful
5. vsc调试
6. ORM操作
7. 文件下载操作
8. 文件夹与package name应该是一致的。也就是同一个文件夹中的package应该都是一样的。
9. anyType定义



#### 安装教程

1.  安装go
2.  安装beego
3.  安装bee

#### 使用说明

1.  bee new ProjectName
2.  bee run 
3.  bee generate

运行bee run的时候程序就会自动构建打包成一个二进制文件。

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
* 2019/10/25:解决查询数据库，参数从前端传输到后端，利用ORM查询数据返回到前端。ORM框架使用初始化自动建表等
* 2019/10/24：解决app-dev.conf文件包含自己导致栈溢出。