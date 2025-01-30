package web

import (
	"drone_sphere_server/internal/adapter/web/router"
	"drone_sphere_server/internal/domain/user/app"
	"drone_sphere_server/internal/domain/user/repo"
	"drone_sphere_server/internal/infra/rdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Engine 代表使用 Fiber 框架的 Web 服务器引擎。
type Engine struct {
	fiber *fiber.App
	RDB   *rdb.RDB
}

// New 创建一个新的 Engine 实例，并包含一个新的 Fiber 应用。
// 返回值:
//
//	*Engine: 指向新创建的 Engine 实例的指针。
func New(rdb *rdb.RDB) *Engine {
	return &Engine{
		fiber: fiber.New(),
		RDB:   rdb,
	}
}

// Init 初始化 Engine。目前它不执行任何操作，但将来可以扩展。
// 返回值:
//
//	error: 始终返回 nil。
func (e *Engine) Init() error {
	var err error
	api := e.fiber.Group("/api/v1")
	api.Use(cors.New())

	userRouter := api.Group("/user")
	userApp := user_app.NewApplication(repo.NewRepository(e.RDB))
	err = router.UserRoutes(userRouter, userApp)
	if err != nil {
		panic(err)
	}

	return nil
}

// Start 启动 Fiber 应用并监听 8080 端口。
// 返回值:
//
//	error: 如果 Fiber 应用启动失败，则返回错误。
func (e *Engine) Start() error {
	return e.fiber.Listen(":10086")
}

// Stop 优雅地关闭 Fiber 应用。
// 返回值:
//
//	error: 如果 Fiber 应用关闭失败，则返回错误。
func (e *Engine) Stop() error {
	return e.fiber.Shutdown()
}
