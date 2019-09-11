package util

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/inflection"
	"github.com/serenize/snaker"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type TemplateConfig struct {
	SourcePath string
	TargetPath string
	Data       map[string]interface{}
	AfterFunc  func(i []byte) []byte
}

func GetTemplateByPath(path string) *template.Template {
	data, err := ReadAll(path)
	if err != nil {
		panic(err)
	}
	return GetTemplate(string(data))
}

func ExecuteTemplateConf(temps []TemplateConfig) {
	var data []byte
	var err error
	var tm *template.Template
	for _, value := range temps {
		data, err = ReadAll(value.SourcePath)
		if err != nil {
			panic(err)
		}
		tm = GetTemplate(string(data))
		if value.AfterFunc != nil {
			ExecuteTemplateBase(tm, value.TargetPath, value.Data, value.AfterFunc)
		} else {
			ExecuteTemplate(tm, value.TargetPath, value.Data)
		}
	}
}

func ExecuteTemplate(t *template.Template, path string, m interface{}) {
	ExecuteTemplateBase(t, path, m, nil)
	return
}

func ExecuteTemplateBase(t *template.Template, path string, m interface{}, before func([]byte) []byte) {
	defer func() {
		if err := recover(); err != nil {
			println(fmt.Sprintf("path：%s，ExecuteTemplateBase：%v", path, err))
		}
	}()
	//判断是否存在文件
	flag, _ := PathExists(path)
	if !flag {
		f, err := CreateFile(path)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	var buf bytes.Buffer
	err := t.Execute(&buf, m)
	if err != nil {
		panic(err)
		return
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
		return
	}
	if before != nil {
		b = before(b)
	}
	err = ioutil.WriteFile(path, b, 0777)
	return
}

func GetTemplate(t string) (tmpl *template.Template) {
	var funcMap = template.FuncMap{
		"pluralize":        inflection.Plural,
		"title":            strings.Title,
		"toLower":          strings.ToLower,
		"toLowerCamelCase": CamelToLowerCamel,
		"toSnakeCase":      snaker.CamelToSnake,
	}
	tmpl, err := template.New("model").Funcs(funcMap).Parse(t)
	if err != nil {
		panic(err)
		return
	}
	return
}

func CamelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

func CreateFile(path string) (*os.File, error) {
	i := strings.LastIndex(path, "/")
	dir := string([]rune(path)[0:i])
	mErr := MkdirAll(dir)
	if mErr != nil {
		return nil, mErr
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CreateFileContent(path, data string) (err error) {
	f, err := CreateFile(path)
	defer f.Close()
	if err != nil {
		return
	}
	f.WriteString(data)
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkdirAll(path string) error {
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = strings.Replace(path, "/", "\\", -1)
	}
	flog, err := PathExists(path)
	if err != nil {
		return err
	}
	if flog {
		return nil
	}
	err2 := os.MkdirAll(path, os.ModePerm)

	if err2 != nil {
		return err2
	}
	return nil
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
