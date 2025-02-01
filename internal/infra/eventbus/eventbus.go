package eventbus

import (
	"log"
	"reflect"
	"sync"

	asaskevichEventBus "github.com/asaskevich/EventBus"
)

// Event 基础事件接口(空接口,仅用于标记)
type Event any

// Handler 泛型事件处理函数
type Handler[T Event] func(event T) error

// Middleware 泛型中间件类型
type Middleware func(next Handler[Event]) Handler[Event]

// EventBus 统一入口的非泛型事件总线
type EventBus struct {
	bus         asaskevichEventBus.Bus
	middlewares []Middleware
	handlerMaps sync.Map // 存储各类型处理函数的包装器(key: 类型字符串)
	mu          sync.RWMutex
}

// New 创建统一事件总线
func New() *EventBus {
	return &EventBus{
		bus:         asaskevichEventBus.New(),
		middlewares: make([]Middleware, 0),
	}
}

// Use 添加全局中间件
func (eb *EventBus) Use(middlewares ...Middleware) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.middlewares = append(eb.middlewares, middlewares...)
}

// Subscribe 订阅特定类型事件(泛型方法)
func Subscribe[T Event](eb *EventBus, handler Handler[T]) error {
	// 生成类型唯一主题
	topic := getTopicName[T]()

	// 包装处理函数以应用中间件
	wrappedHandler := wrapMiddleware[T](eb, handler)

	// 转换为底层总线需要的函数类型
	fn := func(event T) error {
		if err := wrappedHandler(event); err != nil {
			log.Printf("Event handler failed: %v", err)
			return err
		}
		return nil
	}

	// 存储原始处理函数用于后续中间件更新
	eb.handlerMaps.Store(topic, fn)

	return eb.bus.Subscribe(topic, fn)
}

// Publish 发布事件(泛型方法)
func Publish[T Event](eb *EventBus, event T) {
	topic := getTopicName[T]()
	eb.bus.Publish(topic, event)
}

// wrapMiddleware 应用中间件链(类型安全包装)
func wrapMiddleware[T Event](eb *EventBus, handler Handler[T]) Handler[T] {
	// 将泛型处理函数转换为通用接口类型
	genericHandler := func(e Event) error {
		event, ok := e.(T)
		if !ok {
			return nil
		}
		return handler(event)
	}

	// 应用中间件链
	wrapped := genericHandler
	for i := 0; i < len(eb.middlewares); i++ {
		wrapped = eb.middlewares[i](wrapped)
	}

	// 转换回类型安全处理函数
	return func(event T) error {
		return wrapped(event)
	}
}

// 获取类型唯一主题名
func getTopicName[T Event]() string {
	var t T
	return reflect.TypeOf(t).String()
}
