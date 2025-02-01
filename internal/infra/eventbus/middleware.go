package eventbus

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"
)

// LoggingMiddleware 通用日志中间件
func LoggingMiddleware(logger *slog.Logger) Middleware {
	return func(next Handler[Event]) Handler[Event] {
		return func(event Event) error {
			start := time.Now()
			topic := reflect.TypeOf(event).String()
			logger.Info(fmt.Sprintf("[EventBus] Handling event: %v", event))

			err := next(event)

			duration := time.Since(start)
			if err != nil {
				logger.Warn(fmt.Sprintf("[EventBus] %s failed in %v: %v", topic, duration, err.Error()))
			} else {
				logger.Info("[EventBus] %s handled in %v", topic, duration)
			}
			return err
		}
	}
}

// RecoveryMiddleware 通用错误恢复中间件
func RecoveryMiddleware() Middleware {
	return func(next Handler[Event]) Handler[Event] {
		return func(event Event) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic recovered: %v", r)
				}
			}()
			return next(event)
		}
	}
}
