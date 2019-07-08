`package test

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"{{.PackageName}}/manage-api/handle"
	"{{.PackageName}}/utils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGet{{.StructName}}(t *testing.T) {
	Login(t, func(e *echo.Echo, token string) {
        f := url.Values{}
        f.Add("name", "1")
        req := httptest.NewRequest(echo.POST, "/api/{{.TableName}}", strings.NewReader(f.Encode()))
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        // 断言
        if assert.NoError(t, handle.Get{{.StructName}}(c)) {
            if assert.Equal(t, http.StatusOK, rec.Code) {
                body := rec.Body.String()
                result := utils.ToResultParam(rec.Body.Bytes())
                if result.Ret == 200 {
                    t.Log(body)
                } else {
                    t.Fatal(body)
                }
            } else {
                t.Fatal("http error")
            }
        }
	})
}

func TestGet{{.StructName}}All(t *testing.T) {
	e := echo.New()
	f := url.Values{}
	f.Add("name", "1")
	req := httptest.NewRequest(echo.POST, "/api/{{.TableName}}/all", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// 断言
	if assert.NoError(t, handle.Get{{.StructName}}All(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			body := rec.Body.String()
			result := utils.ToResultParam(rec.Body.Bytes())
			if result.Ret == 200 {
				t.Log(body)
			} else {
				t.Fatal(body)
			}
		} else {
			t.Fatal("http error")
		}
	}
}

func TestGet{{.StructName}}Page(t *testing.T) {
	e := echo.New()
	f := url.Values{}
	f.Add("name", "1")
	req := httptest.NewRequest(echo.POST, "/api/{{.TableName}}/page", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// 断言
	if assert.NoError(t, handle.Get{{.StructName}}Page(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			body := rec.Body.String()
			result := utils.ToResultParam(rec.Body.Bytes())
			if result.Ret == 200 {
				t.Log(body)
			} else {
				t.Fatal(body)
			}
		} else {
			t.Fatal("http error")
		}
	}
}

func TestAdd{{.StructName}}(t *testing.T) {
	e := echo.New()
	f := url.Values{}
	{{ range .Columns }}{{if not .PrimaryKey }}
	f.Add("{{ .ColumnName }}", "{{ .AddTestValue }}"){{end}}{{end}}
	req := httptest.NewRequest(echo.POST, "/api/{{.TableName}}/add", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// 断言
	if assert.NoError(t, handle.Save{{.StructName}}(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			body := rec.Body.String()
			result := utils.ToResultParam(rec.Body.Bytes())
			if result.Ret == 200 {
				t.Log(body)
			} else {
				t.Fatal(body)
			}
		} else {
			t.Fatal("http error")
		}
	}
}

func TestUpdate{{.StructName}}(t *testing.T) {
	e := echo.New()
	f := url.Values{}
	{{ range .Columns }}
	f.Add("{{ .ColumnName }}", "{{ .UpdateTestValue }}"){{end}}
	req := httptest.NewRequest(echo.POST, "/api/{{.TableName}}/update", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// 断言
	if assert.NoError(t, handle.Save{{.StructName}}(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			body := rec.Body.String()
			result := utils.ToResultParam(rec.Body.Bytes())
			if result.Ret == 200 {
				t.Log(body)
			} else {
				t.Fatal(body)
			}
		} else {
			t.Fatal("http error")
		}
	}
}

func TestDel{{.StructName}}(t *testing.T) {
	e := echo.New()
	f := url.Values{}
	f.Add("id", "1")
	req := httptest.NewRequest(echo.POST, "/api/{{.TableName}}/update", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// 断言
	if assert.NoError(t, handle.Del{{.StructName}}(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			body := rec.Body.String()
			result := utils.ToResultParam(rec.Body.Bytes())
			if result.Ret == 200 {
				t.Log(body)
			} else {
				t.Fatal(body)
			}
		} else {
			t.Fatal("http error")
		}
	}
}`
