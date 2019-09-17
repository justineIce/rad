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
	"sort"
	"strings"
	"time"
)

var (
	sqlType    = goopt.String([]string{"--sqltype"}, "mysql", "数据库类型，默认：mysql")
	dbHost     = goopt.String([]string{"--dbHost"}, "127.0.0.1", "数据库主机地址，默认：127.0.0.1")
	dbPort     = goopt.String([]string{"--dbPort"}, "3306", "数据库端口，默认：3306")
	dbUser     = goopt.String([]string{"--dbUser"}, "root", "数据库的用户名，默认：root")
	dbPassword = goopt.String([]string{"--dbPassword"}, "123456", "数据库的密码，默认：123456")

	rdHost     = goopt.String([]string{"--rdHost"}, "127.0.0.1", "Redis数据库主机地址，默认：127.0.0.1")
	rdPort     = goopt.String([]string{"--rdPort"}, "6379", "Redis数据库端口，默认：3306")
	rdPassword = goopt.String([]string{"--rdPassword"}, "", "Redis数据库的密码，默认：无")
	rdDatabase = goopt.String([]string{"--rdDatabase"}, "", "Redis数据库名，默认：0")

	dbName               = goopt.String([]string{"--dbName"}, "test", "数据库名称，默认：test")
	dbParameters         = goopt.String([]string{"--dbParameters"}, "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true", "数据库连接字符串，默认：charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true")
	sqlTable             = goopt.String([]string{"--table"}, "", "需要转换的表名，默认：*")
	sqlView              = goopt.String([]string{"--view"}, "", "需要转换的表名，默认：*")
	packageName          = goopt.String([]string{"--package"}, "", "指定项目包名，默认：当前执行路径/example")
	target               = goopt.String([]string{"--target"}, "", "指定保存目录")
	frontend             = goopt.String([]string{"--frontend"}, "d2admin", "前端模板框架，默认：d2admin")
	backend              = goopt.String([]string{"--backend"}, "echo", "后端模板框架，默认：echo")
	jsonAnnotation       = goopt.Flag([]string{"--json"}, []string{"--no-json"}, "添加json标记，默认：json", "禁用json标记")
	gormAnnotation       = goopt.Flag([]string{"--gorm"}, []string{}, "添加gorm标记，默认：gorm", "")
	gureguTypes          = goopt.Flag([]string{"--guregu"}, []string{}, "支持可为空值的字段类型，默认：guregu", "")
	projectName          string //项目名称，根据包名生成
	packageNameImportURL string //包导入URL
	targetPath           string //项目存放路径
	DB                   *sql.DB
)

var tableWhiteList = map[string]string{
	"file_log":          "",
	"sys_company":       "",
	"sys_dict":          "",
	"sys_log":           "",
	"sys_menu":          "",
	"sys_menu_btn":      "",
	"sys_office":        "",
	"sys_role":          "",
	"sys_role_menu":     "",
	"sys_role_menu_btn": "",
	"sys_user":          "",
	"sys_user_role":     "",
}

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
	if packageName == nil || *packageName == "" {
		packageNameImportURL = fmt.Sprintf("%s/example", strings.ReplaceAll(curPath, fmt.Sprintf("%s/src/", GOPATH), ""))
	} else {
		packageNameImportURL = *packageName
	}
	projectName = packageNameImportURL[strings.LastIndex(packageNameImportURL, "/")+1:]
	if target == nil || *target == "" {
		targetPath = fmt.Sprintf("%s/%s", curPath, projectName)
	} else {
		targetPath = fmt.Sprintf("%s/%s", *target, projectName)
	}
	var err error
	DB, err = sql.Open(*sqlType, getDSN())
	if err != nil {
		fmt.Println("Error in open database: " + err.Error())
		return
	}
	defer DB.Close()

	generateFrontend(curPath)
	generateBackend(curPath)
}

//生成前端框架
func generateFrontend(curPath string) {
	// 默认d2admin
	_, err := util.Copy(fmt.Sprintf("%s/template/frontend/%s", curPath, *frontend), targetPath+"/manage-web", curPath+"/uncopy.txt")
	if err != nil {
		panic(err)
	}
}

//生成后端框架
func generateBackend(curPath string) {
	var err error
	//获取表信息
	var tables, views []string
	if *sqlTable != "" {
		tables = strings.Split(*sqlTable, ",")
	} else {
		tables, err = schema.TableNames(DB)
		if err != nil {
			panic(err)
			return
		}
	}
	sort.Strings(tables)
	if *sqlView != "" {
		views = strings.Split(*sqlView, ",")
	} else {
		views, err = schema.ViewNames(DB)
		if err != nil {
			panic(err)
			return
		}
	}
	sort.Strings(views)
	// 默认echo，复制框架模板
	_, err = util.Copy(fmt.Sprintf("%s/template/backend/%s", curPath, *backend), targetPath, curPath+"/uncopy.txt")
	if err != nil {
		panic(err)
	}
	// 生成代码
	switch *backend {
	case "echo":
		executeBackendEcho(tables, views)
	}

	//生成后端代码
	switch *frontend {
	case "d2admin":
		executeFrontendD2admim(tables, views)
	}
}

func getTargetPath(path string) string {
	return fmt.Sprintf("%s/%s", targetPath, path)
}

func executeBackendEcho(tables, views []string) {
	tempData := map[string]interface{}{
		"PackageName": packageNameImportURL,
		"ProjectName": projectName,
	}
	tc := []util.TemplateConfig{
		{
			SourcePath: "template/backend/echo/manage-api/global/init.go.tpl",
			TargetPath: getTargetPath("manage-api/global/init.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/config/config.toml.tpl",
			TargetPath: getTargetPath("manage-api/config/config.toml"),
			Data: map[string]interface{}{
				"host":       dbHost,
				"port":       dbPort,
				"user":       dbUser,
				"password":   dbPassword,
				"name":       dbName,
				"parameters": dbParameters,
				"rdHost":     rdHost,
				"rdPort":     rdPort,
				"rdDatabase": rdDatabase,
				"rdPassword": rdPassword,
			},
			AfterFunc: func(i []byte) []byte {
				str := string(i)
				return []byte(strings.ReplaceAll(str, "`", ""))
			},
		},
		{
			SourcePath: "template/backend/echo/manage-api/test/init.tpl",
			TargetPath: getTargetPath("manage-api/test/init.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/test/login_test.go.tpl",
			TargetPath: getTargetPath("manage-api/test/login_test.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/file.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/file.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/file_log.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/file_log.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/filter.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/filter.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/login.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/login.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/power.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/power.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_company.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_company.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_dict.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_dict.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_log.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_log.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_menu.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_menu.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_menu_btn.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_menu_btn.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_office.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_office.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_role.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_role.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_role_menu.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_role_menu.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_role_menu_btn.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_role_menu_btn.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_user.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_user.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/handle/sys_user_role.go.tpl",
			TargetPath: getTargetPath("manage-api/handle/sys_user_role.go"),
			Data:       tempData,
		},

		{
			SourcePath: "template/backend/echo/manage-api/router/file_log.go.tpl",
			TargetPath: getTargetPath("manage-api/router/file_log.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_company.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_company.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_dict.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_dict.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_log.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_log.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_menu.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_menu.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_menu_btn.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_menu_btn.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_office.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_office.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_role.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_role.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_role_menu.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_role_menu.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_role_menu_btn.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_role_menu_btn.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_user.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_user.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/manage-api/router/sys_user_role.go.tpl",
			TargetPath: getTargetPath("manage-api/router/sys_user_role.go"),
			Data:       tempData,
		},

		{
			SourcePath: "template/backend/echo/utils/filetype/filetype.go.tpl",
			TargetPath: getTargetPath("utils/filetype/filetype.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/utils/filetype/kind.go.tpl",
			TargetPath: getTargetPath("utils/filetype/kind.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/utils/filetype/match.go.tpl",
			TargetPath: getTargetPath("utils/filetype/match.go"),
			Data:       tempData,
		},
		{
			SourcePath: "template/backend/echo/utils/filetype/matchers/matchers.go.tpl",
			TargetPath: getTargetPath("utils/filetype/matchers/matchers.go"),
			Data:       tempData,
		},
	}
	util.ExecuteTemplateConf(tc)

	// model
	m := util.GetTemplateByPath("template/backend/echo/model/model.tpl")
	// handle
	h := util.GetTemplateByPath("template/backend/echo/manage-api/handle/handle.tpl")
	// router
	r := util.GetTemplateByPath("template/backend/echo/manage-api/router/router.tpl")
	// test
	t := util.GetTemplateByPath("template/backend/echo/manage-api/test/test.tpl")
	// main.go
	ma := util.GetTemplateByPath("template/backend/echo/manage-api/main.tpl")

	var structNames, routers []string
	var modelPath, handlePath, routerPath, testPath string
	var viewsMap = make(map[string]dbmeta.ModelInfo, 0)
	var viewInfo dbmeta.ModelInfo
	//视图
	for _, tableName := range views {
		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)
		modelInfo := dbmeta.GenerateStruct(DB, *dbName, tableName, structName, "model", *jsonAnnotation, *gormAnnotation, *gureguTypes)
		viewsMap[tableName] = *modelInfo
		modelPath = fmt.Sprintf(getTargetPath("model/%s.go"), tableName)
		util.CreateFile(modelPath)
		//model
		util.ExecuteTemplate(m, modelPath, modelInfo)
	}
	//表
	for _, tableName := range tables {
		if _, ok := tableWhiteList[tableName]; ok {
			continue
		}

		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)
		structNames = append(structNames, structName)
		modelInfo := dbmeta.GenerateStruct(DB, *dbName, tableName, structName, "model", *jsonAnnotation, *gormAnnotation, *gureguTypes)
		modelPath = fmt.Sprintf(getTargetPath("model/%s.go"), tableName)
		handlePath = fmt.Sprintf(getTargetPath("manage-api/handle/%s.go"), tableName)
		routerPath = fmt.Sprintf(getTargetPath("manage-api/router/%s.go"), tableName)
		testPath = fmt.Sprintf(getTargetPath("manage-api/test/%s_test.go"), tableName)
		util.CreateFile(modelPath)
		util.CreateFile(handlePath)
		util.CreateFile(routerPath)
		util.CreateFile(testPath)
		viewInfo = dbmeta.ModelInfo{}
		if modelInfo.TableView != "" {
			viewInfo = viewsMap[modelInfo.TableView]
		}
		//model
		util.ExecuteTemplate(m, modelPath, modelInfo)
		//handle
		util.ExecuteTemplate(h, handlePath, map[string]interface{}{"PackageName": packageNameImportURL, "StructName": structName,
			"SingName": modelInfo.SingName, "TableRemark": modelInfo.TableRemark, "IDPrimaryKeyInt": modelInfo.IDPrimaryKeyInt,
			"FieldsMap": modelInfo.FieldsMap, "TableName": tableName, "TableView": modelInfo.TableView, "ViewInfo": viewInfo})
		//test
		util.ExecuteTemplateBase(t, testPath, map[string]interface{}{"PackageName": packageNameImportURL, "StructName": structName,
			"TableName": tableName, "Columns": modelInfo.Columns}, func(i []byte) []byte {
			str := string(i)
			return []byte(strings.ReplaceAll(str, "`", ""))
		})
		//router
		util.ExecuteTemplate(r, routerPath, map[string]string{"PackageName": packageNameImportURL, "StructName": structName, "TableName": tableName})
		// add router
		routers = append(routers, fmt.Sprintf("router.%sRouter(auth)", structName))
	}
	//main
	util.ExecuteTemplateBase(ma, getTargetPath("manage-api/main.go"), map[string]interface{}{"PackageName": packageNameImportURL, "Routers": routers}, func(i []byte) []byte {
		str := string(i)
		return []byte(strings.ReplaceAll(str, "`", ""))
	})
}

//生成d2admin
func executeFrontendD2admim(tables, views []string) {
	// api	生成api
	api := util.GetTemplateByPath("template/frontend/d2admin/src/api/api.js.tpl")
	// pages 生成页面
	page := util.GetTemplateByPath("template/frontend/d2admin/src/pages/page.vue.tpl")
	// router 生成路由
	router := util.GetTemplateByPath("template/frontend/d2admin/src/routerMapComponents/index.js.tpl")
	// menu 生成菜单
	menu := util.GetTemplateByPath("template/frontend/d2admin/src/menu/aside.js.tpl")
	var tableList []*dbmeta.ModelInfo
	/*var structNames, routers []string
	var modelPath, handlePath, routerPath, testPath, singName string
	var fieldsMap map[string]bool*/
	var routerPath, pagePath string
	var viewsMap = make(map[string]dbmeta.ModelInfo, 0)
	var viewInfo dbmeta.ModelInfo
	//视图
	for _, tableName := range views {
		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)
		modelInfo := dbmeta.GenerateStruct(DB, *dbName, tableName, structName, "model", *jsonAnnotation, *gormAnnotation, *gureguTypes)
		viewsMap[tableName] = *modelInfo
	}
	//表
	for _, tableName := range tables {
		if _, ok := tableWhiteList[tableName]; ok {
			continue
		}
		structName := dbmeta.FmtFieldName(tableName)
		structName = inflection.Singular(structName)
		modelInfo := dbmeta.GenerateStruct(DB, *dbName, tableName, structName, "model", *jsonAnnotation, *gormAnnotation, *gureguTypes)
		viewInfo = dbmeta.ModelInfo{}
		if modelInfo.TableView != "" {
			viewInfo = viewsMap[modelInfo.TableView]
		}
		//api生成
		routerPath = fmt.Sprintf(getTargetPath("manage-web/src/api/%s.js"), modelInfo.SingName)
		util.ExecuteTemplateBase(api, routerPath, map[string]interface{}{"modelInfo": modelInfo}, func(i []byte) []byte {
			str := string(i)
			return []byte(strings.ReplaceAll(str, "`", ""))
		})
		//page生成
		pagePath = fmt.Sprintf(getTargetPath("manage-web/src/pages/%s/index.vue"), modelInfo.SingName)
		util.ExecuteTemplate(page, pagePath, map[string]interface{}{"modelInfo": modelInfo, "ViewInfo": viewInfo})
		//生成数据库菜单
		createSysMenu(fmt.Sprintf("%sIndex", modelInfo.SingName), modelInfo.TableRemark, fmt.Sprintf("/%s/index", modelInfo.SingName),
			fmt.Sprintf("pages/%s/index", modelInfo.SingName), modelInfo.SingName)
		// add router
		tableList = append(tableList, modelInfo)
	}

	//main
	util.ExecuteTemplateBase(router, getTargetPath("manage-web/src/routerMapComponents/index.js"),
		map[string]interface{}{"tableList": tableList}, func(i []byte) []byte {
			str := string(i)
			return []byte(strings.ReplaceAll(strings.ReplaceAll(str, "`", ""), "&", "`"))
		})
	//menu
	util.ExecuteTemplateBase(menu, getTargetPath("manage-web/src/menu/aside.js"), map[string]interface{}{"tableList": tableList}, func(i []byte) []byte {
		str := string(i)
		return []byte(strings.ReplaceAll(strings.ReplaceAll(str, "`", ""), "&", "`"))
	})
}

// 创建系统菜单
func createSysMenu(routeName, tableRemark, href, target, singName string) {
	rows, err := Query("SELECT COUNT(1) AS c FROM sys_menu WHERE route_name=?", routeName)
	if err != nil {
		println(err.Error())
		return
	}
	if len(rows) <= 0 {
		return
	}
	count := rows[0]["c"].(int64)
	if count == 0 {
		id := dbmeta.IDString()
		now := time.Now()
		//新建路由
		_, err = DB.Exec("INSERT INTO `sys_menu`(`id`, `parent_id`, `relation_ids`, `name`, `sort`, `href`, `target`, `icon`,"+
			" `route_name`, `created_by`, `created_at`, `updated_by`, `updated_at`, `remarks`, `deleted_at`, `component`) "+
			"VALUES (?, '2000', ?, ?, 0.00, ?, ?, 'fa fa-th-list', ?, '1',?, '1', ?, NULL, NULL, ?);",
			id, fmt.Sprintf("-0-2000-%s-", id), tableRemark, href, target, routeName, now, now, fmt.Sprintf("Business_%sIndex", singName))
		if err != nil {
			println(err.Error())
		}
	}
}

func Query(sqlStr string, val ...interface{}) (list []map[string]interface{}, err error) {
	var rows *sql.Rows
	rows, err = DB.Query(sqlStr, val...)
	if err != nil {
		return
	}
	defer rows.Close()
	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		return
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// 这里需要初始化为空数组，否则在查询结果为空的时候，返回的会是一个未初始化的指针
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}

		ret := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				ret[columns[i]] = nil
			} else {
				switch val := (*scanArgs[i].(*interface{})).(type) {
				case []byte:
					ret[columns[i]] = string(val)
					break
				case time.Time:
					ret[columns[i]] = val.Format("2006-01-02 15:04:05")
					break
				default:
					ret[columns[i]] = val
				}
			}
		}
		list = append(list, ret)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
