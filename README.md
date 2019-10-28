# beego-demo

#### 介绍
beego学习demo

#### 学习内容
1. logs
2. router，控制器的学习与使用
3. docker化应用
4. restful
5. vsc调试

#### plan
* ORM框架
* JWT
* schedule

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

