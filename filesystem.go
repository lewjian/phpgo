package phpgo

import (
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// Basename 返回某个文件路径的文件名，suffix如果提供了则去掉改后缀
func Basename(filename string, suffix ...string) string {
	name := filepath.Base(filename)
	for _, sf := range suffix {
		name = strings.TrimRight(name, sf)
	}
	return name
}

// Chown 改变文件所属者
func Chown(filename, username string) bool {
	ex := FileExists(filename)
	if !ex {
		return false
	}
	// windows下无意义
	if GetSystemOS() == "linux" {
		sysUser, err := user.Lookup(username)
		if err != nil {
			return false
		}
		uid, err := strconv.Atoi(sysUser.Uid)
		if err != nil {
			return false
		}
		gid, err := strconv.Atoi(sysUser.Gid)
		if err != nil {
			return false
		}
		err = os.Chown(filename, uid, gid)
		if err != nil {
			return false
		}
	}
	return true

}

// FileExists 检查文件是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// GetSystemOS 获取当前运行的操作系统类型
func GetSystemOS() string {
	return runtime.GOOS
}

// Copy 复制文件
func Copy(src, dest string) bool {
	if !FileExists(src) {
		return false
	}
	fileSrc, err := os.Open(src)
	if err != nil {
		return false
	}
	defer fileSrc.Close()
	fileDest, err := os.Create(dest)
	if err != nil {
		return false
	}
	defer fileDest.Close()
	_, err = io.Copy(fileDest, fileSrc)
	if err != nil {
		return false
	}
	return true
}

// UnLink 删除文件
func UnLink(filename string) bool {
	return os.Remove(filename) == nil
}

// Dirname 返回路径中的目录部分
func Dirname(path string) string {
	return filepath.Dir(path)
}

// FileGetContents 获取文件内容
func FileGetContents(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return data
}

// FilePutContents 写入文件内容
func FilePutContents(filename string, content []byte, fileMode os.FileMode) (int, error) {
	handle, err := os.OpenFile(filename, os.O_WRONLY, fileMode)
	if err != nil {
		return 0, err
	}
	defer handle.Close()
	n, err := handle.Write(content)
	if err != nil {
		return n, err
	}
	return n, nil
}

// FileATime 获取文件上次访问时间，返回为秒的时间戳
func FileATime(filename string) int64 {
	info, _ := os.Stat(filename)
	sys := info.Sys()
	if sys != nil {
		fileSys := sys.(*syscall.Win32FileAttributeData)
		return fileSys.LastAccessTime.Nanoseconds() / 1e9
	}
	return 0
}

// FileMTime 获取文件修改时间
func FileMTime(filename string) int64 {
	info, _ := os.Stat(filename)
	return info.ModTime().Unix()
}

// FileSize 获取文件大小
func FileSize(filename string) int64 {
	info, _ := os.Stat(filename)
	return info.Size()
}

// IsDir 判断是否为文件夹
func IsDir(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile 判断是否为文件
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsReadable 判断是否可读
func IsReadable(filename string) bool {
	h, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer h.Close()
	return true
}

// IsWritable 判断文件是否可写
func IsWritable(filename string) bool {
	h, err := os.OpenFile(filename, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return false
	}
	_ = h.Close()
	return true
}

// MkDir 创建文件夹
func MkDir(pathname string, fileMode os.FileMode, recursive bool) bool {
	if recursive {
		return os.MkdirAll(pathname, fileMode) == nil
	}
	return os.Mkdir(pathname, fileMode) == nil
}

// Rename 重命名
func Rename(src, dest string) error {
	return os.Rename(src, dest)
}
