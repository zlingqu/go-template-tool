package svc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/jwilder/gojq"
)

type Context struct {
}

func (c *Context) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func contains(item map[string]string, key string) bool {
	if _, ok := item[key]; ok {
		return true
	}
	return false
}

func defaultValue(args ...interface{}) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("default called with no values!")
	}

	if len(args) > 0 {
		if args[0] != nil {
			return args[0].(string), nil
		}
	}

	if len(args) > 1 {
		if args[1] == nil {
			return "", fmt.Errorf("default called with nil default value!")
		}

		if _, ok := args[1].(string); !ok {
			return "", fmt.Errorf("default is not a string value. hint: surround it w/ double quotes.")
		}

		return args[1].(string), nil
	}

	return "", fmt.Errorf("default called with no default value")
}

func parseUrl(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("unable to parse url %s: %s", rawurl, err)
	}
	return u
}

func add(arg1, arg2 int) int {
	return arg1 + arg2
}

func isTrue(s string) bool {
	b, err := strconv.ParseBool(strings.ToLower(s))
	if err == nil {
		return b
	}
	return false
}

func jsonQuery(jsonObj string, query string) (interface{}, error) {
	parser, err := gojq.NewStringQuery(jsonObj)
	if err != nil {
		return "", err
	}
	res, err := parser.Query(query)
	if err != nil {
		return "", err
	}
	return res, nil
}

func loop(args ...int) (<-chan int, error) {
	var start, stop, step int
	switch len(args) {
	case 1:
		start, stop, step = 0, args[0], 1
	case 2:
		start, stop, step = args[0], args[1], 1
	case 3:
		start, stop, step = args[0], args[1], args[2]
	default:
		return nil, fmt.Errorf("wrong number of arguments, expected 1-3"+
			", but got %d", len(args))
	}

	c := make(chan int)
	go func() {
		for i := start; i < stop; i += step {
			c <- i
		}
		close(c)
	}()
	return c, nil
}

func GenerateFile(templatePath, destPath string) bool {
	tmpl := template.New(filepath.Base(templatePath)).Funcs(template.FuncMap{ //自定义了很多方法，可以在模板中直接使用
		"contains":  contains,
		"exists":    exists,
		"split":     strings.Split,
		"replace":   strings.Replace,
		"default":   defaultValue,
		"parseUrl":  parseUrl,
		"atoi":      strconv.Atoi,
		"add":       add,
		"isTrue":    isTrue,
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"jsonQuery": jsonQuery,
		"loop":      loop,
	})

	tmpl, err := tmpl.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("unable to parse template: %s", err)
	}

	if _, err := os.Stat(destPath); err == nil {
		// return false        // 如果存在这个文件，则退出函数，不做操作
		os.Remove(destPath) //如果文件存在就删除
	}

	dest := os.Stdout
	if destPath != "" {

		dest, err = os.Create(destPath)
		if err != nil {
			log.Fatalf("unable to create %s", err)
		}
		defer dest.Close()
	}

	err = tmpl.ExecuteTemplate(dest, filepath.Base(templatePath), &Context{})
	if err != nil {
		log.Fatalf("template error: %s\n", err)
	}

	return true
}

func GenerateDir(templateDir, destDir string) bool {
	if destDir != "" {
		fiDest, err := os.Stat(destDir)
		if err != nil {
			log.Fatalf("unable to stat %s, error: %s", destDir, err)
		}
		if !fiDest.IsDir() {
			log.Fatalf("if template is a directory, dest must also be a directory (or stdout)")
		}
	}

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatalf("bad directory: %s, error: %s", templateDir, err)
	}

	for _, file := range files {
		fileName := strings.Split(file.Name(), ".")[0]
		fileName = fileName + ".yml"
		if destDir == "" {
			GenerateFile(filepath.Join(templateDir, fileName), "")
		} else {
			GenerateFile(filepath.Join(templateDir, file.Name()), filepath.Join(destDir, fileName))
		}
	}

	return true
}
