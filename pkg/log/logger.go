package log

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"sync"
)

var (
	loggerInstance *slog.Logger
	once           sync.Once
	Logger         = GetLogger()
)

// GetLogger 返回一个单例的 slog.Logger 实例，并使用易于人类阅读理解的 formatter。
func GetLogger() *slog.Logger {
	// 使用 sync.Once 来确保 loggerInstance 只被初始化一次。
	once.Do(func() {
		loggerInstance = slog.New(tint.NewHandler(os.Stdout, nil))
	})
	return loggerInstance
}
