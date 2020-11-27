package main

import (
	"log"

	// cfg "github.com/zlingqu/go-template-tool/config"
	"github.com/zlingqu/go-template-tool/cmd"
	// svc "github.com/zlingqu/go-template-tool/service"
)

func main() {
	// config := cfg.InitConfig("config/k8s.config")
	// for k, v := range config {
	// 	os.Setenv(k, v)
	// }

	newGenetateYml := cmd.NewGenetateYmlCommand()
	if err := newGenetateYml.Execute(); err != nil {
		log.Fatal(err)
	}
	// environ := os.Environ()
	// for i := range environ {
	// 	fmt.Println(environ[i])
	// }

	// //支持对目录、单个文件、多个文件进行操作
	// GenerateDestfile("src_dir:dest_dir")
	// // GenerateDestfile("src_dir/deployment.tmpl:dest_dir/deployment.yml")
	// GenerateDestfile("src_dir/deployment.tmpl:dest_dir/deployment.yml","src_dir/svc.tmpl:dest_dir/svc.yml")
}
