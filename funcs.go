package toolkit

// 通用方法
import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

/**
* 检测一个文件是否存在
* 前提是你有权限访问该文件
 */
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}

/**
* 获得当前执行程序的目录
 */
func GetExecDir() string {
	file, _ := exec.LookPath(os.Args[0])
	absfile, _ := filepath.Abs(file)
	return filepath.Dir(absfile)
}

/**
* 获得当前时间戳
 */
func GetUnixTime() int64 {
	return time.Now().Unix()
}

/**
* 获得当前时间，格式为yyyy-mm-dd HH:ii:ss
 */
func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// @title 获得远程连接的内容
// @param url string 远程地址
// @return io.Reader
func GetRemoteFile(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// save content to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.String(), nil
}
