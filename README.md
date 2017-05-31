# gohomework

使用beego创建api工程，通过api完成，对etcd的某一个key,value值的添加，获取和删除操作。捕获信号量的方式打印cpu,mem,stack的内容分别到不同的文件。

最终可以通过命令go tool pprof --pdf /tmp/ms-server /tmp/cpu1 > /root/Desktop/${IP}cpu.pdf和go tool pprof --pdf /tmp/ms-server /tmp/mem.1 > /root/Desktop/${IP}mem.pdf，可以看到cpu和mem的分析内容
