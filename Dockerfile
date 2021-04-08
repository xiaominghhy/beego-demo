FROM golang:1.9.1

# Install beego and the bee dev tool*
RUN go get github.com/beego/beego && go get github.com/beego/bee

#指定好工作目录，后面就会以这个作为相对路径
WORKDIR /home/app/
#add app
ADD  beego-demo  beego-demo 

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
# CMD ["sh", "beego-demo "]
ENTRYPOINT [ "sh", "-c", "/home/app/beego-demo "]

# 作者：xhaoxiong
# 链接：https://hacpai.com/article/1526210600840
# 来源：黑客派
# 协议：CC BY-SA 4.0 https://creativecommons.org/licenses/by-sa/4.0/