# 使用go 模板生成yaml文件
## 使用方法
配置文件：config/k8s.config, 支持多端口，支持多路径挂载


模板文件：src_dest/*

查看main.go文件，修改相关内容后执行以下操作

```
go run main.go
```
进入dest_dir目录即可查看生成的yaml文件

