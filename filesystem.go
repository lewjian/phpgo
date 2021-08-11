package phpgo

import (
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

// FileATime 获取文件上次访问时间
// 对应windows下使用：
// fileSys := sys.(*syscall.Win32FileAttributeData)
// second := fileSys.LastAccessTime.Nanoseconds() / 1e9
// 对于linux下使用
// fileSys := sys.(*syscall.Stat_t)
// second := fileSys.Atim.Sec
func FileATime(filename string) (sys interface{}) {
	info, _ := os.Stat(filename)
	sys = info.Sys()
	return
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

// ChMod 改变文件权限
func ChMod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// ChDir 改变当前目录
func ChDir(dir string) bool {
	return os.Chdir(dir) == nil
}

// GetCWD 取得当前工作目录
func GetCWD() string {
	dir, _ := os.Getwd()
	return dir
}

// ReadDir 读取目录，golang里面不用opendir
func ReadDir(dir string) ([]os.DirEntry, error) {
	return os.ReadDir(dir)
}
