package logging

import (
	"io"
	"log/slog"
)

func defaultOptions() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				s := a.Value.Any().(*slog.Source)
				s.File = Shortpath(s.File, 2)
			}
			return a
		}}
}

// NewSFDailyLogger 天切片结构化日志
func NewSFDailyLogger(path string, level slog.Level) (*slog.Logger, io.Closer, error) {
	trf, err := NewTimedRotatingFile(path, WhenDaily, LogFileNameDailyFormat)
	if err != nil {
		return nil, nil, err
	}
	opts := defaultOptions()
	opts.Level = level
	logger := slog.New(slog.NewTextHandler(trf, opts))
	return logger, trf, nil
}

// NewSFHourlyLogger 小时切片结构化日志
func NewSFHourlyLogger(path string, level slog.Level) (*slog.Logger, io.Closer, error) {
	trf, err := NewTimedRotatingFile(path, WhenHourly, LogFileNameHourlyFormat)
	if err != nil {
		return nil, nil, err
	}

	opts := defaultOptions()
	opts.Level = level
	logger := slog.New(slog.NewTextHandler(trf, opts))
	return logger, trf, nil
}
