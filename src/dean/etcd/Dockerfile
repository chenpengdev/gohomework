FROM ubuntu:latest
MAINTAINER liuzhiyang <276796194@qq.com>
RUN ["mkdir","/homework"]
COPY etcd /homework/
WORKDIR /homework/
ENTRYPOINT ["/homework/etcd"]
EXPOSE 8080