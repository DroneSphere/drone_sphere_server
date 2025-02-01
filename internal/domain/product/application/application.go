package application

import (
	"context"
	"drone_sphere_server/internal/core/adapter"
	userapp "drone_sphere_server/internal/domain/user/application"
	"drone_sphere_server/internal/infra/eventbus"
	mqtt2 "drone_sphere_server/internal/infra/mqtt"
	"drone_sphere_server/pkg/log"
	"encoding/json"
	"fmt"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Application represents the product application instance
type Application struct {
	bus *eventbus.EventBus
	mq  *mqtt2.MQTT
}

// New creates a new Application instance
func New(bus *eventbus.EventBus, mq *mqtt2.MQTT) *Application {
	app := &Application{
		bus: bus,
		mq:  mq,
	}

	eventbus.Subscribe(bus, app.handleLoginSuccess)

	return app
}

type SNKeyType string

const SNKey SNKeyType = "sn"

// handleLoginSuccess handles the login success event
func (a *Application) handleLoginSuccess(event userapp.LoginSuccessEvent) error {
	log.GetLogger().Info("Product application received login success event")
	log.GetLogger().Info("Event: ", slog.Any("event", event))

	ctx := context.WithValue(context.Background(), SNKey, event.SN)
	topic := fmt.Sprintf("sys/product/%s/status", event.SN)
	replyTopic := fmt.Sprintf("sys/product/%s/status_reply", event.SN)

	err := a.mq.SubscribeTopic(ctx, topic, func(ctx context.Context, message mqtt.Message) error {
		sn := ctx.Value(SNKey).(SNKeyType)
		if len(sn) <= 0 {
			log.GetLogger().Warn("SN is empty")
			return nil
		}

		payload := struct {
			adapter.CommonModel
			Data UpdateTopoCommand `json:"data"`
		}{}
		err := json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			log.GetLogger().Error("Unmarshal payload error: ", slog.Any("error", err.Error()))
			return nil
		}
		log.GetLogger().Info("Payload: ", slog.Any("payload", payload))

		resp := struct {
			adapter.CommonModel
			Data struct {
				Result int `json:"result"`
			} `json:"data"`
		}{
			CommonModel: payload.CommonModel,
			Data: struct {
				Result int `json:"result"`
			}{Result: 0},
		}
		log.GetLogger().Info("Response: ", slog.Any("response", resp))

		bytes, err := json.Marshal(resp)
		if err != nil {
			log.GetLogger().Error("Marshal response error: ", slog.Any("error", err.Error()))
			return nil
		}

		token := a.mq.Client.Publish(replyTopic, 1, false, bytes)
		if token.Wait() && token.Error() != nil {
			log.GetLogger().Error("Publish reply error: ", slog.Any("error", token.Error()))
			return nil
		}
		log.GetLogger().Info("Successfully published reply message")

		return nil
	})

	if err != nil {
		log.GetLogger().Error("Subscribe topic error: ", slog.Any("error", err.Error()))
		return err
	}

	return nil
}
