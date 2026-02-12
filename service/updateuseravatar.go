package service

import (
	"errors"
	"fmt"
	"go-web/global"
	"go-web/models"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func isImage(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

func SaveUploadedFile(file *multipart.FileHeader, dst string, perm ...fs.FileMode) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var mode os.FileMode = 0o750
	if len(perm) > 0 {
		mode = perm[0]
	}
	dir := filepath.Dir(dst)
	if err = os.MkdirAll(dir, mode); err != nil {
		return err
	}
	if err = os.Chmod(dir, mode); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func UpdateUserAvatar(userID uint, file *multipart.FileHeader) (string, error) {
	//  校验文件类型
	if !isImage(file.Filename) {
		return "", errors.New("仅支持 jpg/png/jpeg 格式")
	}
	// 获取文件后缀
	ext := path.Ext(file.Filename)
	// 拼接文件名
	filename := fmt.Sprintf("image_%d_%d%s", userID, time.Now().Unix(), ext)

	// 拼出“文件最终保存的完整路径”
	saveDir := "uploads/avatar/"
	// path.Join 是用来“拼接路径的”，并且会自动处理 / 的问题
	savePath := path.Join(saveDir, filename)

	// 保存文件 （使用 gin 的工具方法）
	if err := SaveUploadedFile(file, savePath); err != nil {
		return "", err
	}

	//  生成头像访问 URL
	avatarURL := "/static/avatar/" + filename

	// 更新数据库
	err := global.DB.Model(&models.SysUser{}).Where("id = ?", userID).Update("avatar", avatarURL).Error
	if err != nil {
		return "", err
	}

	return avatarURL, nil
}
