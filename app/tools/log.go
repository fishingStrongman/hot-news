package tools

//按照时间划分日志
import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

const (
	status200 = 42
	status404 = 43
	status500 = 41

	methodGET = 44
)

func (hook FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook FileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02_15-04")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
	filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

type MyFormatter struct {
}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// 设置buffer 缓冲区
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	// 设置格式
	fmt.Fprintf(b, "%s\n", entry.Message)

	return b.Bytes(), nil
}

func InitFile(logPath, appName string) {
	logrus.SetFormatter(&MyFormatter{})

	fileDate := time.Now().Format("2006-01-02_15-04")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	fileHook := FileDateHook{file, logPath, fileDate, appName}

	logrus.AddHook(&fileHook)
}
