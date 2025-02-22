package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var Logobj *logrus.Logger

type MarkdownFormatter struct {
	TimestampFormat string
}

func InitLogger() {
	// 如果 Logobj 已初始化，直接返回
	if Logobj != nil {
		return
	}
	// 初始化 Logobj
	Logobj = logrus.New()
	Logobj.SetLevel(logrus.InfoLevel)

	wd, err := os.Getwd()
	if err != nil {
		// 如果获取工作目录失败，使用标准输出并记录错误
		Logobj.SetOutput(os.Stderr)
		Logobj.Errorf("Failed to get working directory: %v", err)
		return
	}

	currentTime := time.Now().Format("2006-01-02")
	logDir := filepath.Join(wd, "static", "log")
	logFileName := filepath.Join(logDir, currentTime+".log")

	// 确保日志目录存在
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			// 如果创建目录失败，记录到标准错误输出
			Logobj.SetOutput(os.Stderr)
			Logobj.Errorf("Failed to create log directory: %v", err)
			return
		}
	}

	// 打开或创建日志文件
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果打开文件失败，记录到标准错误输出
		Logobj.SetOutput(os.Stderr)
		Logobj.Errorf("Failed to open log file: %v", err)
		return
	}
	Logobj.SetOutput(file)

	// 设置日志格式
	Logobj.SetFormatter(&MarkdownFormatter{
		//键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	})
}

func (m *MarkdownFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(m.TimestampFormat)
	level := entry.Level.String()
	message := entry.Message

	// 创建Markdown格式的日志行
	markdownLog := fmt.Sprintf("## %s\n**%s** %s\n", level, timestamp, message)

	// 遍历字段，将它们也格式化为Markdown
	for key, value := range entry.Data {
		markdownLog += fmt.Sprintf("- **%s**: `%v`\n", key, value)
	}

	markdownLog += "\n" // 添加换行符以分隔日志条目

	return []byte(markdownLog), nil
}
