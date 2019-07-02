package dbmeta

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type ModelInfo struct {
	PackageName     string
	StructName      string
	ShortStructName string
	TableName       string
	TableRemark     string
	Fields          []string
	PrimaryKey      []string
	FieldDefVal     map[string]interface{}
	Columns         []*ColumnInfo
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
	ColumnName             string      `json:"column_name"`              //字段名称
	IsNullable             string      `json:"is_nullable"`              //是否允许为空
	DataType               string      `json:"data_type"`                //字段类型
	CharacterMaximumLength string      `json:"character_maximum_length"` //长度
	CharacterOctetLength   string      `json:"character_octet_length"`   //字符八位字节长度
	NumericPrecision       string      `json:"numeric_precision"`        //double/floag/numeric 长度
	NumericScale           string      `json:"numeric_scale"`            //小数点
	ColumnComment          string      `json:"column_comment"`           //字段备注
	AddTestValue           interface{} //字段test时值
	UpdateTestValue        interface{} //字段test时值
	PrimaryKey             bool        //是否为主键
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
	numeric_scale,column_comment from information_schema.columns where table_name='%s' AND table_schema='%s'`
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
	//primaryKey, _ := GetTablePrimaryKey(db, dbName, tableName)
	fieldDef, primaryKey, _ := GetTableDesc(db, tableName)
	fields := generateFieldsTypes(cols, fieldDef, primaryKey, jsonAnnotation, gormAnnotation, gureguTypes)
	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructName:      structName,
		TableName:       tableName,
		TableRemark:     tableRemark,
		ShortStructName: strings.ToLower(string(structName[0])),
		Fields:          fields,
		PrimaryKey:      primaryKey,
		Columns:         cols,
	}
	return modelInfo
}

// Generate fields string
func generateFieldsTypes(columns []*ColumnInfo, fieldDef map[string]interface{}, primaryKey []string, jsonAnnotation bool,
	gormAnnotation bool, gureguTypes bool) (fields []string) {
	var field = ""
	for _, c := range columns {
		key := c.ColumnName
		valueType, valid, addDefVal, updateDefVal := sqlTypeToGoType(c, gureguTypes)
		if valueType == "" { // unknown type
			continue
		}
		if addDefVal != nil {
			c.AddTestValue = addDefVal
		}
		if updateDefVal != nil {
			c.UpdateTestValue = updateDefVal
		}
		if Contains(primaryKey, key) {
			c.PrimaryKey = Contains(primaryKey, key)
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
			annotations = append(annotations, fmt.Sprintf("json:\"%s\" form:\"%s\" query:\"%s\"", key, key, key))
		}
		if len(valid) > 0 {
			annotations = append(annotations, fmt.Sprintf("valid:\"%s\"", strings.Join(valid, ";")))
		}
		if c.ColumnComment != "" {
			//逗号分隔
			alias := strings.ReplaceAll(c.ColumnComment, "，", ",")
			comm := strings.Split(alias, ",")
			annotations = append(annotations, fmt.Sprintf("validAlias:\"%s\"", comm[0]))
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

func Contains(str []string, s string) (flag bool) {
	for _, value := range str {
		if value == s {
			flag = true
			return
		}
	}
	return
}

func sqlTypeToGoType(c *ColumnInfo, gureguTypes bool) (varType string, valid []string, addDefVal interface{}, updateDefVal interface{}) {
	var nullable bool
	if strings.ToLower(c.IsNullable) == "yes" {
		nullable = true
	} else {
		if c.ColumnName != "updated_at" && c.ColumnName != "created_at" {
			valid = append(valid, "Required")
		}
	}
	mysqlType := strings.ToLower(c.DataType)
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		addDefVal = GetRandom(StrToInt(c.NumericPrecision) / 2)
		updateDefVal = GetRandom(StrToInt(c.NumericPrecision) / 2)
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
		addDefVal = GetRandom(StrToInt(c.NumericPrecision) / 2)
		updateDefVal = GetRandom(StrToInt(c.NumericPrecision) / 2)
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
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		addDefVal = GetRandomString(StrToInt(c.CharacterMaximumLength) / 2)
		updateDefVal = GetRandomString(StrToInt(c.CharacterMaximumLength) / 2)
		valid = append(valid, fmt.Sprintf("MaxSize(%s)", c.CharacterMaximumLength))
		if nullable {
			valid = append(valid, "omitempty")
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
		addDefVal = fmt.Sprintf("%d.%d", GetRandom(StrToInt(c.NumericPrecision)/2), GetRandom(StrToInt(c.NumericScale)-1))
		updateDefVal = fmt.Sprintf("%d.%d", GetRandom(StrToInt(c.NumericPrecision)/2), GetRandom(StrToInt(c.NumericScale)-1))
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
		addDefVal = fmt.Sprintf("%d.%d", GetRandom(StrToInt(c.NumericPrecision)/2), GetRandom(StrToInt(c.NumericScale)-1))
		updateDefVal = fmt.Sprintf("%d.%d", GetRandom(StrToInt(c.NumericPrecision)/2), GetRandom(StrToInt(c.NumericScale)-1))
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
		addDefVal = []byte(GetRandomString(StrToInt(c.CharacterMaximumLength) / 2))
		updateDefVal = []byte(GetRandomString(StrToInt(c.CharacterMaximumLength) / 2))
		varType = golangByteArray
		return
	}
	return
}
