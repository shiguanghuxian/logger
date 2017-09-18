/*
 * @Author: 时光弧线
 * @Date: 2017-09-08 16:14:30
 * @Last Modified by: 时光弧线
 * @Last Modified time: 2017-09-08 17:12:06
 */
package logger

import (
	"io"
	"os"

	glogger "github.com/shiguanghuxian/google-logger"
)

// Logger 日志记录对象
type Logger struct {
	name           string         // 用于 syslog tag
	verbose        bool           // 是否输出到终端
	systemLog      bool           // 是否启用syslog
	maxFileSize    int64          // 日志文件的最大byte数
	timedTask      int            // 定时检查文件时间(秒)
	formatFileName FormatFileName // 格式化日志文件名（文件路径）
	logFile        io.Writer      // 日志输出对象
}

const (
	// DefaultMaxFileSize 默认日志文件最大值(10M)
	DefaultMaxFileSize = 10485760
	// DefaultTimedTask 默认定时检查文件时间
	DefaultTimedTask = 60
)

var (
	logger *Logger
	lf     *os.File
)

func init() {
	// 创建日志记录对象
	logger = &Logger{
		name:           "ErrorLog",
		verbose:        true,
		systemLog:      false,
		maxFileSize:    DefaultMaxFileSize,
		timedTask:      DefaultTimedTask,
		formatFileName: defaultFormatFileName,
	}
}

// Init 初始化日志记录对象
func Init(name string, verbose, systemLog bool, formatFileName ...FormatFileName) (err error) {
	if name != "" {
		logger.name = name
	}
	logger.verbose = verbose
	logger.systemLog = systemLog
	// 判断是否自定义了日志文件生成函数
	if len(formatFileName) != 0 {
		logger.formatFileName = formatFileName[0]
	}
	// 开启定时任务，每夜更新
	initCron()
	// 调用初始化日志记录文件对象
	return ResetLogger()
}

// Close 结束日志纪录，关闭打开的日志文件
func Close() (err error) {
	// 文件如果已经打开过日志文件，则先关闭文件
	if lf != nil {
		err = lf.Close()
	}
	return
}

// ResetLogger 刷新生成日志文件
func ResetLogger() error {
	fieName, err := logger.formatFileName()
	if err != nil {
		return err
	}
	// 文件如果已经打开过日志文件，则先关闭文件
	Close()
	lf, err = os.OpenFile(fieName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	// 初始化谷歌日志对象
	glogger.Init(logger.name, logger.verbose, logger.systemLog, lf)
	return nil
}

// Info 输出info类型的日志,不回车换行fmt.Print
func Info(v ...interface{}) {
	glogger.Info(v...)
}

// Infoln 输出info类型的日志,回车换行fmt.Println
func Infoln(v ...interface{}) {
	glogger.Infoln(v...)
}

// Infof 输出info类型的日志,可以使用fmt.Sprintf使用占位符
func Infof(format string, v ...interface{}) {
	glogger.Infof(format, v...)
}

// Warn 输出warn类型的日志,不回车换行fmt.Print
func Warn(v ...interface{}) {
	glogger.Warn(v...)
}

// Warnln 输出warn类型的日志,回车换行fmt.Println
func Warnln(v ...interface{}) {
	glogger.Warnln(v...)
}

// Warnf 输出warn类型的日志,可以使用fmt.Sprintf使用占位符
func Warnf(format string, v ...interface{}) {
	glogger.Warnf(format, v...)
}

// Error 输出error类型的日志,不回车换行fmt.Print
func Error(v ...interface{}) {
	glogger.Error(v...)
}

// Errorln 输出error类型的日志,回车换行fmt.Println
func Errorln(v ...interface{}) {
	glogger.Errorln(v...)
}

// Errorf 输出error类型的日志,可以使用fmt.Sprintf使用占位符
func Errorf(format string, v ...interface{}) {
	glogger.Errorf(format, v...)
}

// Fatal 输出fatal类型的日志,不回车换行fmt.Print,并调用 os.Exit(1).
func Fatal(v ...interface{}) {
	glogger.Fatal(v...)
}

// Fatalln 输出fatal类型的日志,回车换行fmt.Println,并调用 os.Exit(1).
func Fatalln(v ...interface{}) {
	glogger.Fatalln(v...)
}

// Fatalf 输出fatal类型的日志,可以使用fmt.Sprintf使用占位符,并调用 os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	glogger.Fatalf(format, v...)
}
