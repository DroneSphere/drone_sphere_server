package web

import (
	web "drone_sphere_server/internal/adapter/web/router"
	userapp "drone_sphere_server/internal/domain/user/application"
	"drone_sphere_server/internal/infra/eventbus"
	"drone_sphere_server/internal/infra/rdb"
	"drone_sphere_server/pkg/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Engine 代表使用 Fiber 框架的 Web 服务器引擎。
type Engine struct {
	fiber    *fiber.App
	RDB      *rdb.RDB
	EventBus *eventbus.EventBus
}

// New 创建一个新的 Engine 实例，并包含一个新的 Fiber 应用。
// 返回值:
//
//	*Engine: 指向新创建的 Engine 实例的指针。
func New(rdb *rdb.RDB, eventBus *eventbus.EventBus) *Engine {
	return &Engine{
		fiber:    fiber.New(),
		RDB:      rdb,
		EventBus: eventBus,
	}
}

// Init 初始化 Engine。目前它不执行任何操作，但将来可以扩展。
// 返回值:
//
//	error: 始终返回 nil。
func (e *Engine) Init() error {
	e.fiber.Use(cors.New())
	e.fiber.Use(recover.New())
	return nil
}

// RegisterApps 注册应用程序到 Engine。
// 参数:
//
//	apps: map[string]interface{}: 包含应用程序的 map。
//
// 返回值:
//
//	error: 如果注册失败，则返回错误。
func (e *Engine) RegisterApps(apps map[string]interface{}) {
	group := e.fiber.Group("/api/v1")
	for name, app := range apps {
		switch name {
		case "user":
			web.RegisterUserRoutes(group.Group("/user"), app.(*userapp.Application))
		case "product":
			log.GetLogger().Info("product application is not implemented yet")
		default:
			panic("Unknown application: " + name)
		}
	}
}

// Start 启动 Fiber 应用并监听 8080 端口。
// 返回值:
//
//	error: 如果 Fiber 应用启动失败，则返回错误。
func (e *Engine) Start() error {
	return e.fiber.Listen(":10086")
}
