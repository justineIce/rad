package main

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/devopsmi/rad/dbmeta"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/droundy/goopt"
	"github.com/jimsmart/schema"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/inflection"
	_ "github.com/lib/pq"
	"github.com/serenize/snaker"
)

var (
	sqlType        = goopt.String([]string{"--sqltype"}, "mysql", "数据库类型，默认：mysql")
	dbHost         = goopt.String([]string{"--dbHost"}, "127.0.0.1", "数据库主机地址，默认：127.0.0.1")
	dbPort         = goopt.String([]string{"--dbPort"}, "3306", "数据库端口，默认：3306")
	dbUser         = goopt.String([]string{"--dbUser"}, "root", "数据库的用户名，默认：root")
	dbPassword     = goopt.String([]string{"--dbPassword"}, "123456", "数据库的密码，默认：123456")
	dbName         = goopt.String([]string{"--dbName"}, "test", "数据库名称，默认：test")
	dbParameters   = goopt.String([]string{"--dbParameters"}, "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true", "数据库连接字符串，默认：charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true")
	sqlTable       = goopt.String([]string{"--table"}, "", "需要转换的表名，默认：*")
	packageName    = goopt.String([]string{"--package"}, "", "指定项目包名，默认：当前执行路径/example")
	jsonAnnotation = goopt.Flag([]string{"--json"}, []string{"--no-json"}, "添加json标记，默认：json", "禁用json标记")
	gormAnnotation = goopt.Flag([]string{"--gorm"}, []string{}, "添加gorm标记，默认：gorm", "")
	gureguTypes    = goopt.Flag([]string{"--guregu"}, []string{}, "支持可为空值的字段类型，默认：guregu", "")
)

func init() {
	// Setup goopts
	goopt.Description = func() string {
		return "Automatic generation of echo + model + API + MySQL Web Framework."
	}
	goopt.Version = "0.1"
	goopt.Summary = `rad [-v] [--dbHost] [--dbPort] [--dbUser] [--dbPassword] [--dbName] [--dbParameters] --package pkgName --table tableName [--json] [--gorm] [--guregu]`
	//Parse options
	goopt.Parse(nil)
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", *dbUser, *dbPassword, *dbHost, *dbPort, *dbName, *dbParameters)
}

func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		println(errors.New("无有效的GOPATH工种目录"))
	} else {
		GOPATH = strings.Replace(GOPATH, "\\", "/", -1)
	}
	curPath := getCurrentPath()
	var db, err = sql.Open(*sqlType, getDSN())
	if err != nil {
		fmt.Println("Error in open database: " + err.Error())
		return
	}
	defer db.Close()
	// parse or read tables
	var tables []string
	if *sqlTable != "" {
		tables = strings.Split(*sqlTable, ",")
	} else {
		tables, err = schema.TableNames(db)
		if err != nil {
			fmt.Println("Error in fetching tables information from mysql information schema")
			return
		}
	}
	// if packageName is not set we need to default it
	if packageName == nil || *packageName == "" {
		*packageName = fmt.Sprintf("%s/example", strings.ReplaceAll(curPath, fmt.Sprintf("%s/src/", GOPATH), ""))
	}
	// global init.go
	data, err := readAll("template/globalInit.tpl")
	if err != nil {
		panic(err)
	}
	gi := getTemplate(string(data))
	ExecuteTemplate(gi, "example/global/init.go", map[string]interface{}{
		"PackageName": *packageName,
	})
	// utils uuid.go
	data, err = readAll("template/uuid.tpl")
	if err != nil {
		panic(err)
	}
	createFileContent("example/utils/uuid.go", string(data))
	// echo bind.go
	data, err = readAll("template/bind.tpl")
	if err != nil {
		panic(err)
	}
	createFileContent("example/handle/bind.go", string(data))
	//result
	data, err = readAll("template/result.tpl")
	if err != nil {
		panic(err)
	}
	createFileContent("example/utils/result.go", string(data))
	//gorm
	data, err = readAll("template/gorm.tpl")
	if err != nil {
		panic(err)
	}
	createFileContent("example/utils/gorm.go", string(data))
	//page
	data, err = readAll("template/page.tpl")
	if err != nil {
		panic(err)
	}
	createFileContent("example/model/page.go", string(data))
	// config.go
	data, err = readAll("template/config.tpl")
	if err != nil {
		panic(err)
	}
	createFileContent("example/config/config.go", string(data))
	// config.toml
	data, err = readAll("template/configtoml.tpl")
	if err != nil {
		panic(err)
	}
	cft := getTemplate(string(data))
	ExecuteTemplateBase(cft, "example/config/config.toml", map[string]interface{}{
		"host":       dbHost,
		"port":       dbPort,
		"user":       dbUser,
		"password":   dbPassword,
		"name":       dbName,
		"parameters": dbParameters,
	}, func(i []byte) []byte {
		str := string(i)
		return []byte(strings.ReplaceAll(str, "`", ""))
	})
	// model
	data, err = readAll("template/model.tpl")
	if err != nil {
		panic(err)
	}
	m := getTemplate(string(data))
	// handle
	data, err = readAll("template/handle.tpl")
	if err != nil {
		panic(err)
	}
	h := getTemplate(string(data))
	// router
	data, err = readAll("template/router.tpl")
	if err != nil {
		panic(err)
	}
	r := getTemplate(string(data))
	// test
	data, err = readAll("template/test.tpl")
	if err != nil {
		panic(err)
	}
	t := getTemplate(string(data))
	// test init.go
	data, err = readAll("template/testInit.tpl")
	if err != nil {
		panic(err)
	}
	ti := getTemplate(string(data))
	ExecuteTemplate(ti, "example/test/init.go", map[string]interface{}{
		"PackageName": *packageName,
	})
	// main.go
	data, err = readAll("template/main.tpl")
	if err != nil {
		panic(err)
	}
	ma := getTemplate(string(data))

	var structNames, routers []string
	var modelPath, handlePath, routerPath, testPath, singName string
	// generate go files for each table
	for _, tableName := range tables {
		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)
		structNames = append(structNames, structName)
		modelInfo := dbmeta.GenerateStruct(db, *dbName, tableName, structName, "model", *jsonAnnotation, *gormAnnotation, *gureguTypes)
		singName = inflection.Singular(tableName)
		modelPath = fmt.Sprintf("example/model/%s.go", singName)
		handlePath = fmt.Sprintf("example/handle/%s.go", singName)
		routerPath = fmt.Sprintf("example/router/%s.go", singName)
		testPath = fmt.Sprintf("example/test/%s_test.go", singName)
		createFile(modelPath)
		createFile(handlePath)
		//model
		ExecuteTemplate(m, modelPath, modelInfo)
		//handle
		ExecuteTemplate(h, handlePath, map[string]string{"PackageName": *packageName, "StructName": structName,
			"singName": dbmeta.FmtFieldName2(tableName), "TableRemark": modelInfo.TableRemark})
		//test
		ExecuteTemplate(t, testPath, map[string]interface{}{"PackageName": *packageName, "StructName": structName,
			"SingName": singName, "FieldDefVal": modelInfo.FieldDefVal, "Columns": modelInfo.Columns})
		//router
		ExecuteTemplate(r, routerPath, map[string]string{"PackageName": *packageName, "StructName": structName, "SingName": singName})
		// add router
		routers = append(routers, fmt.Sprintf("router.%sRouter(api)", structName))
	}
	//main
	ExecuteTemplateBase(ma, "example/main.go", map[string]interface{}{"PackageName": *packageName, "Routers": routers}, func(i []byte) []byte {
		str := string(i)
		return []byte(strings.ReplaceAll(str, "`", ""))
	})
}

func ExecuteTemplate(t *template.Template, path string, m interface{}) {
	ExecuteTemplateBase(t, path, m, nil)
	return
}

func ExecuteTemplateBase(t *template.Template, path string, m interface{}, before func([]byte) []byte) {
	//判断是否存在文件
	flag, _ := pathExists(path)
	if !flag {
		f, err := createFile(path)
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

func getTemplate(t string) (tmpl *template.Template) {
	var funcMap = template.FuncMap{
		"pluralize":        inflection.Plural,
		"title":            strings.Title,
		"toLower":          strings.ToLower,
		"toLowerCamelCase": camelToLowerCamel,
		"toSnakeCase":      snaker.CamelToSnake,
	}
	tmpl, err := template.New("model").Funcs(funcMap).Parse(t)
	if err != nil {
		panic(err)
		return
	}
	return
}

func camelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

func createFile(path string) (*os.File, error) {
	i := strings.LastIndex(path, "/")
	dir := string([]rune(path)[0:i])
	mErr := mkdirAll(dir)
	if mErr != nil {
		return nil, mErr
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func createFileContent(path, data string) (err error) {
	f, err := createFile(path)
	defer f.Close()
	if err != nil {
		return
	}
	f.WriteString(data)
	return
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkdirAll(path string) error {
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = strings.Replace(path, "/", "\\", -1)
	}
	flog, err := pathExists(path)
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

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
