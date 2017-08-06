#!/bin/sh
go install etcd
mv ../../bin/etcd .
docker build -t homework/etcd .
docker run -d -p 8990:8080 --restart=always homework/etcd