package service

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

//
// keyValueConfigToMap key\value格式的配置文件转换成map
func keyValueConfigToMap(path string) map[string]string {
	config := make(map[string]string)

	fi, err := os.Stat(path)
	if err != nil || fi.IsDir() {
		log.Fatalf("bad file: %s, error1: %s", path, err)
	}

	// basePath, _ := os.Getwd()
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	for {
		b, _, err := r.ReadLine() //逐行读取
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=") //如果有多个=号，也取第一个
		if index < 0 {                 //未出现=号，跳过
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 { //判断行首是=号，跳过
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue //判断行尾是=号，跳过
		}
		config[key] = value
	}
	return config
}
