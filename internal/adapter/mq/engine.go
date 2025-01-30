package mq

import (
	"context"
	"drone_sphere_server/internal/infra/emqx"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Engine struct {
	Conf   emqx.ConnectConfig
	Client *emqx.EMQX
}

type Callback func(ctx context.Context, client mqtt.Client) error
