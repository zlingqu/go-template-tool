# 使用go 模板生成yaml文件
## 使用方法
配置文件示例：k8s.config


模板文件示例：src_dest/*，k8s.tmpl

查看帮助
```bash
go run main.go --help
```

例如：
```bash
go run main.go --config k8s.config --src k8s.tmpl --dest k8s.yml
```