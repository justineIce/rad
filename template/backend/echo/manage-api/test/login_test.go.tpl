package test

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"{{.PackageName}}/manage-api/handle"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/convert"
	"{{.PackageName}}/utils/encrypt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	Login(t, func(e *echo.Echo, token string) {
		t.Log(token)
	})
}

func Login(t *testing.T, back func(e *echo.Echo, token string)) {
	e := echo.New()
	f := url.Values{}
	f.Add("username", "test")
	f.Add("password", encrypt.Md5("666666"))
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// 断言
	if assert.NoError(t, handle.Login(c)) {
		if assert.Equal(t, http.StatusOK, rec.Code) {
			body := rec.Body.String()
			result := utils.ToResultParam(rec.Body.Bytes())
			if result.Ret == 200 {
				t.Log(body)
				var token model.SysToken
				_ = json.Unmarshal(convert.MustJson(result.Data), &token)
				back(e, token.Token)
			} else {
				t.Fatal(body)
			}
		} else {
			t.Fatal("http error")
		}
	}
}
