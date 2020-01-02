# 引入go运行环境
FROM golang

# 设置变量app_env
ARG app_env
ENV APP_ENV $app_env

# 将当前路径下的文件拷贝到polar文件夹下，并将其设置成工作路径
COPY . /go/src/polar
WORKDIR /go/src/polar

# 创建log文件夹
RUN mkdir /etc/smartlog

# 安装依赖库
# RUN go get ./
# 编译
RUN go build -o polar .

# 运行服务
CMD ./polar

# 对外开放端口
EXPOSE 8088
