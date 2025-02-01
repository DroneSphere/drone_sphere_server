package main

import (
	"drone_sphere_server/internal/adapter/web"
	productapp "drone_sphere_server/internal/domain/product/application"
	userapp "drone_sphere_server/internal/domain/user/application"
	"drone_sphere_server/internal/domain/user/repo"
	"drone_sphere_server/internal/infra/eventbus"
	"drone_sphere_server/internal/infra/mqtt"
	"drone_sphere_server/internal/infra/rdb"
	"drone_sphere_server/pkg/log"
)

func main() {
	var err error
	logger := log.GetLogger()

	// Infra 组件的初始化
	db := rdb.New()
	logger.Info("Database connection established.")

	eb := eventbus.New()
	eb.Use(eventbus.LoggingMiddleware(logger))
	logger.Info("Event bus created.")

	mq := mqtt.New(mqtt.Config{
		Protocol: "tcp",
		Broker:   "47.245.40.222",
		Port:     1883,
		Username: "server",
		Password: "server",
	})
	err = mq.Init()
	if err != nil {
		panic(err)
	}

	// 领域应用Map初始化
	var apps = make(map[string]interface{})
	apps["user"] = userapp.New(repo.NewRepository(db), eb)
	apps["product"] = productapp.New(eb, mq)

	// Web 服务器初始化
	webServ := web.New(db, eb)
	logger.Debug("Web server created.")
	err = webServ.Init()
	if err != nil {
		panic(err)
	}
	webServ.RegisterApps(apps)
	err = webServ.Start()
	if err != nil {
		panic(err)
	}

	logger.Info("All engines started successfully.")
}
