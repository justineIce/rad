package dbmeta

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jimsmart/schema"
)

type ModelInfo struct {
	PackageName     string
	StructName      string
	ShortStructName string
	TableName       string
	TableRemark     string
	Fields          []string
	PrimaryKey      []string
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
func GetTableColumnRemark(db *sql.DB, dbName, tableName string) (remark map[string]string, err error) {
	var data []map[string]interface{}
	data, err = Query(db, fmt.Sprintf("SELECT COLUMN_NAME,column_comment FROM INFORMATION_SCHEMA.Columns WHERE TABLE_NAME='%s' AND table_schema='%s'", tableName, dbName))
	if err != nil {
		return
	}
	if len(data) > 0 {
		//var cols []map[string]interface{}
		remark = make(map[string]string)
		for _, value := range data {
			name := value["COLUMN_NAME"]
			comment := value["column_comment"]
			remark[fmt.Sprintf("%v", name)] = fmt.Sprintf("%v", comment)
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
	cols, _ := schema.Table(db, tableName)
	colRemark, _ := GetTableColumnRemark(db, dbName, tableName)
	tableRemark, _ := GetTableRemark(db, dbName, tableName)
	primaryKey, _ := GetTablePrimaryKey(db, dbName, tableName)
	fields := generateFieldsTypes(cols, colRemark, primaryKey, 0, jsonAnnotation, gormAnnotation, gureguTypes)

	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructName:      structName,
		TableName:       tableName,
		TableRemark:     tableRemark,
		ShortStructName: strings.ToLower(string(structName[0])),
		Fields:          fields,
		PrimaryKey:      primaryKey,
	}

	return modelInfo
}

// Generate fields string
func generateFieldsTypes(columns []*sql.ColumnType, colRemark map[string]string, primaryKey []string, depth int, jsonAnnotation bool,
	gormAnnotation bool, gureguTypes bool) (fields []string) {
	var field = ""
	for _, c := range columns {
		nullable, _ := c.Nullable()
		key := c.Name()
		valueType := sqlTypeToGoType(strings.ToLower(c.DatabaseTypeName()), nullable, gureguTypes)
		if valueType == "" { // unknown type
			continue
		}
		fieldName := FmtFieldName(stringifyFirstChar(key))

		var annotations []string
		if gormAnnotation == true {
			if Contains(primaryKey, key) {
				annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s;primary_key\"", key))
			} else {
				annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s\"", key))
			}
		}
		if jsonAnnotation == true {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", key))
		}
		if len(annotations) > 0 {
			field = fmt.Sprintf("%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			field = fmt.Sprintf("%s %s",
				fieldName,
				valueType)
		}
		if colRemark != nil {
			field += fmt.Sprintf(" // %s", colRemark[key])
		}
		fields = append(fields, field)
	}
	return fields
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

func sqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}
