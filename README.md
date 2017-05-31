# gohomework

使用beego创建api工程，通过api完成，对etcd的某一个key,value值的添加，获取和删除操作。捕获信号量的方式打印cpu,mem,stack即可。最终通过命令go tool pprof --pdf /tmp/ms-server /tmp/cpu1 > /root/Desktop/${IP}cpu.pdf
go tool pprof --pdf /tmp/ms-server /tmp/mem.1 > /root/Desktop/${IP}mem.pdf，可以看到cpu和mem的分析图片
