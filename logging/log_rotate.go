package logging

import (
	"os"
	"time"
)

const (
	// WhenMinutely 分钟级
	WhenMinutely = 1
	// WhenHourly 小时级
	WhenHourly = 2
	// WhenDaily 天级
	WhenDaily = 3
)

var (
	// LogFileNameDailyFormat 天级文件名模式
	LogFileNameDailyFormat = "2006-01-02"
	// LogFileNameHourlyFormat 小时级文件名模式
	LogFileNameHourlyFormat = "2006-01-02T15"
	// LogFileNameMinutelyFormat 分钟级文件名模式
	LogFileNameMinutelyFormat = "2006-01-02T15:04"
)

// TimedRotatingFile implement io.Writer
//	时间切分文件, 不能多进程使用
//	log.Logger模块在Output()时本身会加锁,因此这里没有额外加锁
type TimedRotatingFile struct {
	fp       *os.File
	when     int
	path     string
	format   string
	lastTime time.Time
}

// NewTimedRotatingFile 创建时间切分的文件
func NewTimedRotatingFile(path string, when int, format string) (*TimedRotatingFile, error) {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return &TimedRotatingFile{fp: fp, when: when, lastTime: time.Now(), path: path, format: format}, nil
}

// NewDailyRotatingFile 创建按天切分的文件
func NewDailyRotatingFile(path string) (*TimedRotatingFile, error) {
	return NewTimedRotatingFile(path, WhenDaily, LogFileNameDailyFormat)
}

// NewHourlyRotatingFile 创建按小时切分的文件
func NewHourlyRotatingFile(path string) (*TimedRotatingFile, error) {
	return NewTimedRotatingFile(path, WhenHourly, LogFileNameHourlyFormat)
}

// Write 文件写入
func (trf *TimedRotatingFile) Write(buf []byte) (int, error) {
	now := time.Now()
	changed := false
	switch trf.when {
	case WhenMinutely:
		changed = now.Minute() != trf.lastTime.Minute()
	case WhenHourly:
		changed = now.Hour() != trf.lastTime.Hour()
	case WhenDaily:
		changed = now.Day() != trf.lastTime.Day()
	default:
	}
	trf.lastTime = now

	if changed {
		err := trf.fp.Close()
		if err != nil {
			return 0, err
		}
		err = os.Rename(trf.path, trf.path+"."+trf.lastTime.Format(trf.format))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		if err != nil {
			return 0, err
		}
		fp, err := os.OpenFile(trf.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return 0, err
		}
		trf.fp = fp
	}
	return trf.fp.Write(buf)
}

// Close 关闭文件
func (trf *TimedRotatingFile) Close() error {
	return trf.fp.Close()
}
