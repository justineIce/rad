package handle

import (
	"fmt"
	"github.com/labstack/echo"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/convert"
	"{{.PackageName}}/utils/filetype"
	"io/ioutil"
	"path"
	"strings"
)

var b float64 = 1024

func UploadFile(c echo.Context) error {
	auth := c.Param("auth")
	loginInfo := GetLoginInfo(c)
	//放置目录
	dir := c.FormValue("dir")
	if dir == "" {
		dir = "default"
	}
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		return utils.ErrorNull(c, "获取文件失败")
	}
	if file.Size <= 0 {
		return utils.ErrorNull(c, "空文件")
	}
	if convert.MustInt64(file.Size) > global.Conf.FileUpload.MaxFileSize {
		return utils.ErrorNull(c, fmt.Sprintf("文件过大超出限制%vmb", fmt.Sprintf("%.2f", convert.MustFloat64(global.Conf.FileUpload.MaxFileSize)/b/b)))
	}
	src, err := file.Open()
	if err != nil {
		return utils.ErrorNull(c, "打开文件失败")
	}
	defer src.Close()

	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return utils.ErrorNull(c, "获取文件格式错误")
	}
	ext := path.Ext(file.Filename)

	head := make([]byte, 261)
	_, _ = src.Read(head)
	//格式限制判断
	isExt := strings.Contains(fmt.Sprintf(",%v,", global.Conf.FileUpload.ExtFilter), fmt.Sprintf(",%v%v,", buf[0], buf[1]))
	if !isExt && !filetype.IsImage(buf) && !filetype.IsVideo(buf) && !filetype.IsAudio(buf) && !filetype.IsArchive(buf) && !filetype.IsDocument(buf) {
		return utils.ErrorNull(c, fmt.Sprintf("%v文件格式错误，ext：%v%v", ext, buf[0], buf[1]))
	}
	fileName := convert.ToString(utils.ID())
	path := GetPath(auth, dir, fileName, ext, loginInfo)
	dst, err := utils.CreateFile(path)
	if err != nil {
		return utils.ErrorNull(c, "创建文件失败")
	}
	defer dst.Close()
	if _, err = dst.Write(buf); err != nil {
		return utils.ErrorNull(c, "保存文件失败")
	}
	path = fmt.Sprintf("/%s", path)
	go func() {
		err = global.DB.Create(&model.FileLog{
			ID:        utils.IDString(),
			Name:      file.Filename,
			Path:      path,
			Size:      file.Size,
			Ext:       ext,
			IP:        c.RealIP(),
			CreatedBy: loginInfo.ID,
		}).Error
		if err != nil {
			global.Log.Error(fmt.Sprintf("保存上传文件日志失败，ERROR：%s", err.Error()))
		}
	}()
	return utils.SuccessNullMsg(c, map[string]interface{}{
		"id":   fileName,
		"size": file.Size,
		"path": path,
		"url":  getUrl(path),
		"name": file.Filename,
		"ext":  ext,
	})
}
func GetPath(auth, dir, fileName, suffix string, loginInfo model.SysUserLoginInfo) string {
	if dir == "" {
		dir = "file"
	}
	if auth != "" {
		dir += auth
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s%s", global.Conf.FileUpload.BasePath+global.Conf.FileUpload.Path, loginInfo.ID, dir, strings.Replace(utils.CurrentDate(), "-", "/", -1), fileName, suffix)
}

func getUrl(path string) string {
	if strings.Contains(path, "http") {
		return global.Conf.FileUpload.DoMain + path
	} else {
		return path
	}
}
