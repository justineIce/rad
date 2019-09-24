package dbmeta

import (
	"github.com/jinzhu/inflection"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
)

type Employee struct {
	EmpNo     int       `gorm:"column:emp_no;primary_key" json:"emp_no"`
	BirthDate time.Time `gorm:"column:birth_date" json:"birth_date"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Gender    string    `gorm:"column:gender" json:"gender"`
	HireDate  time.Time `gorm:"column:hire_date" json:"hire_date"`
}

func Test_Copy(t *testing.T) {
	now := time.Now()
	dst := &Employee{
		EmpNo:     10001,
		BirthDate: now,
		FirstName: "Tom",
	}

	src := &Employee{
		EmpNo:     10001,
		BirthDate: now.Add(3600 * time.Second),
		FirstName: "Jerry",
		Gender:    "Male",
	}

	err := Copy(dst, src)
	if err != nil {
		t.Fatal(err)
	}

	expected := &Employee{
		EmpNo:     10001,
		BirthDate: now.Add(3600 * time.Second),
		FirstName: "Jerry",
		Gender:    "Male",
	}

	if !reflect.DeepEqual(expected, dst) {
		t.Errorf("expect: %+v, but got %+x", expected, dst)
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(GetRandom(10))
	}
}

func TestRandomStr(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(GetRandomString(10))
	}
}

func TestRegexpStr(t *testing.T) {
	var reg = regexp.MustCompile(`enum\((.*?)\)`)
	params := reg.FindStringSubmatch("enum('保密','男','女')")
	t.Log(params)
	if len(params) > 1 {
		v := strings.TrimRight(params[1], ";")
		if v != "" {
			reg = regexp.MustCompile(`'(.*?)'`)
			params = reg.FindAllString("enum('保密','男','女')", -1)
			t.Log(params)
		}
	}
}

func TestIDStr(t *testing.T) {
	t.Log(inflection.Singular(FmtFieldName("id")))
}
