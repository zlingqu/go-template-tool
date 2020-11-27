package service

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//GenetateYml 判断是文件还是目录，生成yml
func GenetateYml(config, src, dest string) {

	b1, tp1 := checkFileType(config)
	b2, tp2 := checkFileType(src)
	b3, tp3 := checkFileType(dest)

	if !b1 || tp1 == "dir" {
		log.Fatalf("bad 参数: %s", config)
	} else { //存在且是文件就转换成环境变量
		cfgMap := keyValueConfigToMap(config)
		for k, v := range cfgMap {
			os.Setenv(k, v)
		}
	}

	if !b2 {
		log.Fatalf("bad 参数: %s", src) //src不能不存在
	}

	if tp2 == "dir" {
		if tp3 == "file" { //dest不能是已经存在的文件
			log.Fatalf("bad 参数: %s", dest)
		}

		if !b3 {
			log.Printf("目录%s不存在，自动创建", dest)
			os.Mkdir(dest, os.ModePerm)
		}

		files, _ := ioutil.ReadDir(src)
		for _, file := range files {
			srcPath := filepath.Join(src, file.Name())
			if !strings.HasSuffix(file.Name(), ".tmpl") {
				continue //如果不是以.tmpl结尾跳过循环
			}
			fileName := strings.Split(file.Name(), ".")[0]
			fileName = fileName + ".yml"
			destPath := filepath.Join(dest, fileName)
			log.Printf("生成目的文件%s", destPath)
			generateFile(srcPath, destPath)
		}
	}
	// file, _ := ioutil.ReadFile(src)
	if !strings.HasSuffix(src, ".tmpl") {
		return
	}
	generateFile(src, dest)

}

// checkFileType 判断文件类型
func checkFileType(path string) (bool, string) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, ""
	}
	if fi.IsDir() {
		return true, "dir"
	}

	return true, "file"
}
