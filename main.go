package main

import (
	"log"
	"os"
	"strings"

	cfg "github.com/zlingqu/go-template-tool/config"
	svc "github.com/zlingqu/go-template-tool/service"
)

func GenerateDestfile(templatesFlag ...string) { //不定长参数，可以同时处理多个模板生成
	for _, t := range templatesFlag {
		template, dest := t, ""
		if strings.Contains(t, ":") {
			parts := strings.Split(t, ":")
			if len(parts) != 2 {
				log.Fatalf("bad template argument: %s. expected \"/template:/dest\"", t)
			}
			template, dest = parts[0], parts[1]
		}

		fi, err := os.Stat(template)
		if err != nil {
			log.Fatalf("unable to stat %s, error: %s", template, err)
		}
		if fi.IsDir() {
			svc.GenerateDir(template, dest)
		} else {
			svc.GenerateFile(template, dest)
		}
	}

}

func main() {
	config := cfg.InitConfig("config/k8s.config")
	for k, v := range config {
		os.Setenv(k, v)
	}
	// environ := os.Environ()
	// for i := range environ {
	// 	fmt.Println(environ[i])
	// }

	//支持对目录、单个文件、多个文件进行操作
	GenerateDestfile("src_dir:dest_dir")
	// GenerateDestfile("src_dir/deployment.tmpl:dest_dir/deployment.yml")
	GenerateDestfile("src_dir/deployment.tmpl:dest_dir/deployment.yml","src_dir/svc.tmpl:dest_dir/svc.yml")
}
