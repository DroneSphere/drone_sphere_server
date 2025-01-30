package main

import (
	"drone_sphere_server/internal/adapter/web"
	"drone_sphere_server/internal/infra/emqx"
	"drone_sphere_server/internal/infra/rdb"
	"drone_sphere_server/pkg/log"
)

func main() {
	var err error
	logger := log.GetLogger()

	db := rdb.New()
	logger.Info("Database connection established.")
	
	mq := emqx.New(emqx.ConnectConfig{
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

	webServ := web.New(db)
	logger.Info("Web server created.")
	err = webServ.Init()
	if err != nil {
		logger.Error("Failed to initialize web server: %v", err)
		return
	}
	err = webServ.Start()
	if err != nil {
		logger.Error("Failed to start web server: %v", err)
		return
	}
	defer func() {
		err = webServ.Stop()
		if err != nil {
			logger.Error("Failed to stop web server: %v", err)
		}
	}()

	logger.Info("All engines started successfully.")
}
