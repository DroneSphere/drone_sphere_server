package eventbus_test

import (
	"testing"

	"drone_sphere_server/internal/infra/eventbus"
	"drone_sphere_server/pkg/log"
)

// 自定义事件类型
type UserCreatedEvent struct{ UserID string }
type OrderPaidEvent struct{ OrderID string }

func TestUnifiedEventBus(t *testing.T) {
	eb := eventbus.New()

	// 添加全局中间件
	logger := log.GetLogger()
	eb.Use(
		eventbus.RecoveryMiddleware(),
		eventbus.LoggingMiddleware(logger),
	)

	// 订阅 UserCreatedEvent
	eventbus.Subscribe(eb, func(event UserCreatedEvent) error {
		logger.Info("User created: %s", event.UserID)
		return nil
	})

	// 订阅 OrderPaidEvent
	eventbus.Subscribe(eb, func(event OrderPaidEvent) error {
		logger.Info("Order paid: %s", event.OrderID)
		return nil
	})

	// 发布事件
	eventbus.Publish(eb, UserCreatedEvent{UserID: "123"})
	eventbus.Publish(eb, OrderPaidEvent{OrderID: "456"})

	// 错误类型示例（编译失败）
	// eventbus.Publish(eb, "invalid event") // 编译错误：未实现 Event 接口
}
