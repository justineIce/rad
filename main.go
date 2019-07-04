package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/devopsmi/rad/dbmeta"
	"github.com/devopsmi/rad/util"
	"github.com/droundy/goopt"
	"github.com/jimsmart/schema"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/inflection"
	"os"
	"strings"
	"text/template"
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
	manageStyle    = goopt.String([]string{"--manageStyle"}, "d2admin", "后台模板框架，默认：d2admin")
	jsonAnnotation = goopt.Flag([]string{"--json"}, []string{"--no-json"}, "添加json标记，默认：json", "禁用json标记")
	gormAnnotation = goopt.Flag([]string{"--gorm"}, []string{}, "添加gorm标记，默认：gorm", "")
	gureguTypes    = goopt.Flag([]string{"--guregu"}, []string{}, "支持可为空值的字段类型，默认：guregu", "")
	projectName    string //项目名称，根据包名生成
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
	var packageNameImportUrl string
	if packageName == nil || *packageName == "" {
		packageNameImportUrl = fmt.Sprintf("%s/example", strings.ReplaceAll(curPath, fmt.Sprintf("%s/src/", GOPATH), ""))
	}
	projectName = packageNameImportUrl[strings.LastIndex(packageNameImportUrl, "/")+1:]

	//获取表信息
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

	generateFrontend(curPath,curPath+"/"+projectName)
	generateBackend(curPath,curPath+"/"+projectName)



	tempData := map[string]interface{}{
		"PackageName": packageNameImportUrl,
	}
	tc := []tempConfig{
		{
			sourcePath: "template/globalInit.tpl",
			targetPath: "example/manage-api/global/init.go",
			data:       tempData,
		},
		{
			sourcePath: "template/configtoml.tpl",
			targetPath: "example/manage-api/config/config.toml",
			data: map[string]interface{}{
				"host":       dbHost,
				"port":       dbPort,
				"user":       dbUser,
				"password":   dbPassword,
				"name":       dbName,
				"parameters": dbParameters,
			},
			afterFunc: func(i []byte) []byte {
				str := string(i)
				return []byte(strings.ReplaceAll(str, "`", ""))
			},
		},
		{
			sourcePath: "template/testInit.tpl",
			targetPath: "example/manage-api/test/init.go",
			data:       tempData,
		},
	}
	executeTemplate(tc)

	// model
	m := getTemplate("template/model.tpl")
	// handle
	h := getTemplate("template/handle.tpl")
	// router
	r := getTemplate("template/router.tpl")
	// test
	t := getTemplate("template/test.tpl")
	// main.go
	ma := getTemplate("template/main.tpl")

	var structNames, routers []string
	var modelPath, handlePath, routerPath, testPath, singName string
	// generate go files for each table
	for _, tableName := range tables {
		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)
		structNames = append(structNames, structName)
		modelInfo := dbmeta.GenerateStruct(db, *dbName, tableName, structName, "model", *jsonAnnotation, *gormAnnotation, *gureguTypes)
		singName = inflection.Singular(tableName)
		modelPath = fmt.Sprintf("example/manage-api/model/%s.go", singName)
		handlePath = fmt.Sprintf("example/manage-api/handle/%s.go", singName)
		routerPath = fmt.Sprintf("example/manage-api/router/%s.go", singName)
		testPath = fmt.Sprintf("example/manage-api/test/%s_test.go", singName)
		util.CreateFile(modelPath)
		util.CreateFile(handlePath)
		util.CreateFile(routerPath)
		util.CreateFile(testPath)
		//model
		util.ExecuteTemplate(m, modelPath, modelInfo)
		//handle
		util.ExecuteTemplate(h, handlePath, map[string]string{"PackageName": packageNameImportUrl, "StructName": structName,
			"singName": dbmeta.FmtFieldName2(tableName), "TableRemark": modelInfo.TableRemark})
		//test
		util.ExecuteTemplate(t, testPath, map[string]interface{}{"PackageName": packageNameImportUrl, "StructName": structName,
			"SingName": singName, "FieldDefVal": modelInfo.FieldDefVal, "Columns": modelInfo.Columns})
		//router
		util.ExecuteTemplate(r, routerPath, map[string]string{"PackageName": packageNameImportUrl, "StructName": structName, "SingName": singName})
		// add router
		routers = append(routers, fmt.Sprintf("router.%sRouter(api)", structName))
	}
	//main
	util.ExecuteTemplateBase(ma, "example/manage-api/main.go", map[string]interface{}{"PackageName": packageNameImportUrl, "Routers": routers}, func(i []byte) []byte {
		str := string(i)
		return []byte(strings.ReplaceAll(str, "`", ""))
	})
}

//生成前端框架
func generateFrontend(curPath, targetPath string) {
	// 默认d2admin
	_, err := util.Copy(fmt.Sprintf("%s/template/frontend/%s", curPath, *manageStyle), targetPath+"/example/manage", curPath+"/uncopy.txt")
	if err != nil {
		panic(err)
	}
}

//生成后端框架
func generateBackend(curPath string, targetPath string) {
	// 默认echo
	_, err := util.Copy(fmt.Sprintf("%s/template/backend/", curPath), curPath+"/example/manage", curPath+"/uncopy.txt")
	if err != nil {
		panic(err)
	}
}

func getTemplate(tempPath string) *template.Template {
	data, err := util.ReadAll(tempPath)
	if err != nil {
		panic(err)
	}
	return util.GetTemplate(string(data))
}

func executeTemplate(temps []tempConfig) {
	var data []byte
	var err error
	var tm *template.Template
	for _, value := range temps {
		data, err = util.ReadAll(value.sourcePath)
		if err != nil {
			panic(err)
		}
		tm = util.GetTemplate(string(data))
		if value.afterFunc != nil {
			util.ExecuteTemplateBase(tm, value.targetPath, value.data, value.afterFunc)
		} else {
			util.ExecuteTemplate(tm, value.targetPath, value.data)
		}
	}
}

type tempConfig struct {
	sourcePath string
	targetPath string
	data       map[string]interface{}
	afterFunc  func(i []byte) []byte
}
