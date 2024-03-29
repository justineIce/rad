package dbmeta

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/inflection"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type ModelInfo struct {
	PackageName     string
	StructName      string
	SingName        string
	ShortStructName string
	TableName       string
	TableView       string //使用视图显示
	TableRemark     string
	TableHandle     map[string]bool
	FieldsMap       map[string]bool
	Fields          []string
	PrimaryKey      []string
	Columns         []*ColumnInfo
	IDPrimaryKeyInt bool
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}

var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// Constants for return types of golang
const (
	golangByteArray  = "[]byte"
	gureguNullInt    = "null.Int"
	sqlNullInt       = "sql.NullInt64"
	golangInt        = "int"
	golangInt64      = "int64"
	gureguNullFloat  = "null.Float"
	sqlNullFloat     = "sql.NullFloat64"
	golangFloat      = "float"
	golangFloat32    = "float32"
	golangFloat64    = "float64"
	gureguNullString = "null.String"
	sqlNullString    = "sql.NullString"
	gureguNullTime   = "null.Time"
	golangTime       = "time.Time"
)

type ColumnInfo struct {
	ColumnName              string          `json:"column_name"`              //字段名称
	IsNullable              string          `json:"is_nullable"`              //是否允许为空
	DataType                string          `json:"data_type"`                //字段类型
	ColumnType              string          `json:"column_type"`              //详细字段类型
	CharacterMaximumLength2 string          `json:"character_maximum_length"` //长度
	CharacterOctetLength2   string          `json:"character_octet_length"`   //字符八位字节长度
	NumericPrecision2       string          `json:"numeric_precision"`        //double/floag/numeric 长度
	NumericScale2           string          `json:"numeric_scale"`            //小数点
	CharacterMaximumLength  int             //长度
	CharacterOctetLength    int             //字符八位字节长度
	NumericPrecision        int             //double/floag/numeric 长度
	NumericScale            int             //小数点
	ColumnComment           string          `json:"column_comment"` //字段备注
	ColumnItemValue         []string        //字段项值，如：enum('保密','男','女')
	ColumnCNName            string          //中文名称
	DataTypeLower           string          //小写
	StructName              string          //StructName
	Valid                   []string        //字段验证标记
	FormType                string          //表单类型，text:字符长度小于300输入框，number：数字类型，time:时间类型，date：日期类型，
	AddTestValue            interface{}     //添加测试的，字段test时值
	UpdateTestValue         interface{}     //修改测试的，字段test时值
	PrimaryKey              bool            //是否为主键
	PageTable               map[string]bool //页面控制，search：可搜索，none：不显示，如：Table[none]或Table[search]
	PageForm                map[string]bool //页面控制，none：不显示，如：Form[none]
}

//默认值
func GetTableDesc(db *sql.DB, tableName string) (fieldDef map[string]interface{}, primaryKey []string, err error) {
	var data []map[string]interface{}
	data, err = Query(db, fmt.Sprintf("desc `%s`", tableName))
	if err != nil {
		return
	}
	fieldDef = make(map[string]interface{}, 0)
	for _, value := range data {
		if value["Key"].(string) == "PRI" {
			primaryKey = append(primaryKey, value["Field"].(string))
		}
		if value["Default"] != nil {
			fieldDef[value["Field"].(string)] = value["Default"].(string)
		}
	}
	return
}

//表备注
func GetTableRemark(db *sql.DB, dbName, tableName string) (remark string, err error) {
	var data []map[string]interface{}
	data, err = Query(db, fmt.Sprintf("SELECT TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_NAME='%s' AND table_schema='%s'", tableName, dbName))
	if err != nil {
		return
	}
	if len(data) > 0 {
		remark = data[0]["TABLE_COMMENT"].(string)
	}
	return
}

//主键
func GetTablePrimaryKey(db *sql.DB, dbName, tableName string) (primaryKey []string, err error) {
	var data []map[string]interface{}
	data, err = Query(db, fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.`KEY_COLUMN_USAGE` WHERE table_name='%s'  AND table_schema='%s' AND constraint_name='PRIMARY'", tableName, dbName))
	if err != nil {
		return
	}
	if len(data) > 0 {
		for _, value := range data {
			primaryKey = append(primaryKey, value["COLUMN_NAME"].(string))
		}
	}
	return
}

//字段备注
func GetTableColumnRemark(db *sql.DB, dbName, tableName string) (cols []*ColumnInfo, err error) {
	var data []map[string]interface{}
	s := `select column_name,is_nullable,data_type,character_maximum_length,character_octet_length,numeric_precision,
	numeric_scale,column_comment,column_type from information_schema.columns where table_name='%s' AND table_schema='%s'`
	data, err = Query(db, fmt.Sprintf(s, tableName, dbName))
	if err != nil {
		return
	}
	if len(data) > 0 {
		var b []byte
		b, err = json.Marshal(data)
		if err != nil {
			return
		}
		err = json.Unmarshal(b, &cols)
		if err != nil {
			return
		}
	}
	return
}

func Query(db *sql.DB, queryStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(queryStr, args...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	list := []map[string]interface{}{}
	// 这里需要初始化为空数组，否则在查询结果为空的时候，返回的会是一个未初始化的指针
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
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
		return nil, err
	}
	return list, nil
}

// GenerateStruct generates a struct for the given table.
func GenerateStruct(db *sql.DB, dbName string, tableName string, structName string, pkgName string, jsonAnnotation bool, gormAnnotation bool, gureguTypes bool) *ModelInfo {
	cols, _ := GetTableColumnRemark(db, dbName, tableName)
	tableRemark, _ := GetTableRemark(db, dbName, tableName)
	var isView bool
	if tableRemark == "VIEW" {
		isView = true
	}
	//primaryKey, _ := GetTablePrimaryKey(db, dbName, tableName)
	var fieldDef map[string]interface{}
	var primaryKey []string
	var tableView string
	tableHandle := make(map[string]bool, 0)
	fields, idPrimaryKeyInt := generateFieldsTypes(isView, cols, fieldDef, primaryKey, jsonAnnotation, gormAnnotation, gureguTypes)

	if !isView {
		fieldDef, primaryKey, _ = GetTableDesc(db, tableName)
		if tableRemark != "" {
			//提取handle
			handle := extractHandle(tableRemark)
			if handle != "" {
				handles := strings.Split(handle, ",")
				for _, value := range handles {
					tableHandle[value] = true
				}
			}
			//提取使用视图
			tableView = extractView(tableRemark)
			/*if tableView!=""{
				tableView=inflection.Singular(FmtFieldName(tableView))
			}*/
			//逗号分隔
			alias := strings.ReplaceAll(tableRemark, "，", ",")
			tableRemark = strings.Split(alias, ",")[0]
		}
		if len(tableHandle) == 0 {
			tableHandle["select"] = true
			tableHandle["insert"] = true
			tableHandle["delete"] = true
			tableHandle["update"] = true
		}
	}

	fieldsMap := make(map[string]bool, 0)
	for _, v := range cols {
		fieldsMap[v.ColumnName] = true
	}

	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructName:      structName,
		TableName:       tableName,
		TableRemark:     tableRemark,
		TableHandle:     tableHandle,
		TableView:       tableView,
		ShortStructName: strings.ToLower(string(structName[0])),
		SingName:        FmtFieldName2(tableName),
		Fields:          fields,
		PrimaryKey:      primaryKey,
		Columns:         cols,
		IDPrimaryKeyInt: idPrimaryKeyInt,
		FieldsMap:       fieldsMap,
	}
	return modelInfo
}

// Generate fields string
func generateFieldsTypes(isView bool, columns []*ColumnInfo, fieldDef map[string]interface{}, primaryKey []string, jsonAnnotation bool,
	gormAnnotation bool, gureguTypes bool) (fields []string, idPrimaryKeyInt bool) {
	var field = ""
	for _, c := range columns {
		key := c.ColumnName
		c.PageTable = make(map[string]bool, 0)
		c.PageForm = make(map[string]bool, 0)
		if c.ColumnComment != "" {
			//逗号分隔
			alias := strings.ReplaceAll(c.ColumnComment, "，", ",")
			comm := strings.Split(alias, ",")
			c.ColumnCNName = comm[0]

			//抽取PageTable标记
			pageTable := extractPageTable(c.ColumnComment)
			if pageTable != "" {
				pts := strings.Split(pageTable, ",")
				for _, value := range pts {
					c.PageTable[value] = true
				}
			}
			//抽取PageForm标记
			pageForm := extractPageForm(c.ColumnComment)
			if pageForm != "" {
				pfs := strings.Split(pageForm, ",")
				for _, value := range pfs {
					c.PageForm[value] = true
				}
			}
		}
		c.StructName = inflection.Singular(FmtFieldName(c.ColumnName))
		c.DataTypeLower = strings.ToLower(c.DataType)
		c.CharacterMaximumLength = StrToInt(c.CharacterMaximumLength2)
		c.CharacterOctetLength = StrToInt(c.CharacterOctetLength2)
		c.NumericScale = StrToInt(c.NumericScale2)
		c.NumericPrecision = StrToInt(c.NumericPrecision2)
		if strings.ToLower(key) == "id" {
			switch c.DataTypeLower {
			case "tinyint", "int", "smallint", "mediumint", "bigint":
				idPrimaryKeyInt = true
				break
			}
		}

		valueType, valid, addDefVal, updateDefVal := sqlTypeToGoType(c, gureguTypes, isView)
		if valueType == "" { // unknown type
			continue
		}
		if addDefVal != nil {
			c.AddTestValue = addDefVal
		}
		if updateDefVal != nil {
			c.UpdateTestValue = updateDefVal
		}
		if contains(primaryKey, key) {
			c.PrimaryKey = contains(primaryKey, key)
		}
		fieldName := FmtFieldName(stringifyFirstChar(key))
		var annotations []string
		if gormAnnotation == true {
			var gormTags []string
			gormTags = append(gormTags, fmt.Sprintf("column:%s", key))
			if c.PrimaryKey {
				gormTags = append(gormTags, "primary_key")
			}
			if fieldDef[key] != nil {
				gormTags = append(gormTags, fmt.Sprintf("default:'%v'", fieldDef[key]))
			}
			annotations = append(annotations, fmt.Sprintf("gorm:\"%s\"", strings.Join(gormTags, ";")))
		}
		if jsonAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\" form:\"%s\" query:\"%s\"", c.StructName, key, key))
		}
		if !isView {
			if len(valid) > 0 {
				v := extractValid(c.ColumnComment)
				if v != "" {
					c.Valid = strings.Split(v, ";")
					valid = append(valid, v)
				}
				annotations = append(annotations, fmt.Sprintf("valid:\"%s\"", strings.Join(valid, ";")))
			}
			if c.ColumnCNName != "" {
				annotations = append(annotations, fmt.Sprintf("validAlias:\"%s\"", c.ColumnCNName))
			}
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`", fieldName, valueType, strings.Join(annotations, " "))

		} else {
			field = fmt.Sprintf("%s %s", fieldName, valueType)
		}
		if c.ColumnComment != "" {
			field += fmt.Sprintf(" // %s", c.ColumnComment)
		}
		fields = append(fields, field)
	}
	return
}

func extractValid(str string) string {
	return extract("Valid", str)
}
func extractView(str string) string {
	return extract("View", str)
}
func extractHandle(str string) string {
	return extract("Handle", str)
}
func extractPageTable(str string) string {
	return extract("Table", str)
}
func extractPageForm(str string) string {
	return extract("Form", str)
}
func extract(flag, str string) (v string) {
	var reg = regexp.MustCompile(fmt.Sprintf(`%s\[(.*?)]`, flag))
	params := reg.FindStringSubmatch(str)
	if len(params) > 1 {
		v = strings.TrimRight(params[1], ";")
		return
	}
	return
}

func EnumValue(v string) (val []string) {
	reg := regexp.MustCompile(`'(.*?)'`)
	params := reg.FindAllString(v, -1)
	for _, value := range params {
		val = append(val, strings.Trim(value, "'"))
	}
	return
}

func contains(str []string, s string) (flag bool) {
	for _, value := range str {
		if value == s {
			flag = true
			return
		}
	}
	return
}

func sqlTypeToGoType(c *ColumnInfo, gureguTypes, isView bool) (varType string, valid []string, addDefVal interface{}, updateDefVal interface{}) {
	var nullable bool
	if strings.ToLower(c.IsNullable) == "yes" && !isView {
		nullable = true
	} else {
		if c.ColumnName != "updated_at" && c.ColumnName != "created_at" {
			valid = append(valid, "Required")
		}
	}
	mysqlType := c.DataTypeLower
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		addDefVal = GetRandom(c.NumericPrecision / 2)
		updateDefVal = GetRandom(c.NumericPrecision / 2)
		if nullable {
			if gureguTypes {
				varType = gureguNullInt
				return
			}
			varType = sqlNullInt
			return
		}
		varType = golangInt
		return
	case "bigint":
		addDefVal = GetRandom(c.NumericPrecision / 2)
		updateDefVal = GetRandom(c.NumericPrecision / 2)
		if nullable {
			if gureguTypes {
				varType = gureguNullInt
				return
			}
			varType = sqlNullInt
			return
		}
		varType = golangInt64
		return
	case "char", "varchar", "longtext", "mediumtext", "text", "tinytext":
		addDefVal = GetRandomString(c.CharacterMaximumLength / 2)
		updateDefVal = GetRandomString(c.CharacterMaximumLength / 2)
		valid = append(valid, fmt.Sprintf("MaxSize(%v)", c.CharacterMaximumLength))
		if nullable {
			if gureguTypes {
				varType = gureguNullString
				return
			}
			varType = sqlNullString
			return
		}
		varType = "string"
		return
	case "enum":
		c.ColumnItemValue = EnumValue(c.ColumnType)
		if len(c.ColumnItemValue) > 0 {
			addDefVal = c.ColumnItemValue[0]
		}
		if len(c.ColumnItemValue) > 1 {
			updateDefVal = c.ColumnItemValue[1]
		}
		if nullable {
			if gureguTypes {
				varType = gureguNullString
				return
			}
			varType = sqlNullString
			return
		}
		varType = "string"
		return
	case "date", "datetime", "time", "timestamp":
		addDefVal = time.Now().AddDate(0, 0, rand.Intn(100)).Format("2006-01-02 15:04:05")
		updateDefVal = time.Now().AddDate(0, 0, rand.Intn(100)).Format("2006-01-02 15:04:05")
		if nullable && gureguTypes {
			varType = gureguNullTime
			return
		}
		varType = golangTime
		return
	case "decimal", "double":
		addDefVal = fmt.Sprintf("%d.%d", GetRandom(c.NumericPrecision/2), GetRandom(c.NumericScale-1))
		updateDefVal = fmt.Sprintf("%d.%d", GetRandom(c.NumericPrecision/2), GetRandom(c.NumericScale-1))
		if nullable {
			if gureguTypes {
				varType = gureguNullFloat
				return
			}
			varType = sqlNullFloat
			return
		}
		varType = golangFloat64
		return
	case "float":
		addDefVal = fmt.Sprintf("%d.%d", GetRandom(c.NumericPrecision/2), GetRandom(c.NumericScale-1))
		updateDefVal = fmt.Sprintf("%d.%d", GetRandom(c.NumericPrecision/2), GetRandom(c.NumericScale-1))
		if nullable {
			if gureguTypes {
				varType = gureguNullFloat
				return
			}
			varType = sqlNullFloat
			return
		}
		varType = golangFloat32
		return
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		addDefVal = []byte(GetRandomString(c.CharacterMaximumLength / 2))
		updateDefVal = []byte(GetRandomString(c.CharacterMaximumLength / 2))
		varType = golangByteArray
		return
	}
	return
}
