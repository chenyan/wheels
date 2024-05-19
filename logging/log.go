package logging

import (
	"io"
	"log"
	"os"
	"runtime"
)

var (
	// LogCallDepth 打印调用深度
	LogCallDepth = 2
)

// SimpleLogger 便捷方法
func SimpleLogger(path string, prefix string) (*log.Logger, *os.File, error) {
	return SimpleFLogger(path, prefix, log.Llongfile|log.LstdFlags)
}

// SimpleFLogger ...
func SimpleFLogger(path string, prefix string, flag int) (*log.Logger, *os.File, error) {
	var fp *os.File
	var err error
	if path != "" {
		fp, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return nil, nil, err
		}
	} else {
		fp = os.Stdout
	}
	logger := log.New(fp, prefix, flag)
	return logger, fp, nil
}

// Std 创建一个打在标准输出的日志
func Std(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, log.Llongfile|log.LstdFlags)
}

// NewFDailyLogger 天切片日志
func NewFDailyLogger(path string, prefix string, flag int) (*log.Logger, io.Closer, error) {
	trf, err := NewTimedRotatingFile(path, WhenDaily, LogFileNameDailyFormat)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(trf, prefix, flag)
	return logger, trf, nil
}

// NewFHourlyLogger 小时切片日志
func NewFHourlyLogger(path string, prefix string, flag int) (*log.Logger, io.Closer, error) {
	trf, err := NewTimedRotatingFile(path, WhenHourly, LogFileNameHourlyFormat)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(trf, prefix, flag)
	return logger, trf, nil
}

// CallerInfo 获取运行时位置
//	return:
//		0: function name
//		1: file
//		2: line no
func CallerInfo(skip int) (string, string, int) {
	pc, f, line, ok := runtime.Caller(skip)
	if !ok {
		return "", "", 0
	}
	funcname := runtime.FuncForPC(pc).Name()
	return funcname, f, line
}
