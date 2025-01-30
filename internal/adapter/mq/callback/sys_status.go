package callback

import (
	"context"
	app "drone_sphere_server/internal/domain/product/app"
	"drone_sphere_server/internal/infra/emqx"
	"drone_sphere_server/pkg/log"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func UpdateTopo(ctx context.Context, client mqtt.Client) error {
	// 从 ctx 中提取 SN
	sn := ctx.Value("sn").(string)
	if len(sn) <= 0 {
		log.GetLogger().Warn("SN is empty")
		return nil
	}

	topic := fmt.Sprintf("sys/product/%s/status", sn)
	client.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		log.GetLogger().Info(fmt.Sprintf("Received message: %s from topic: %s\n", message.Payload(), message.Topic()))

		payload := struct {
			emqx.CommonModel
			Data app.UpdateTopoCommand `json:"data"`
		}{}
		err := json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			log.GetLogger().Error(fmt.Sprintf("Unmarshal payload error: %s", err.Error()))
			return
		}
		log.GetLogger().Info(fmt.Sprintf("payload: %+v", payload))

		resp := struct {
			emqx.CommonModel
			Data struct {
				Result int `json:"result"`
			}
		}{}
		resp.CommonModel = payload.CommonModel
		resp.Data.Result = 0
		respBytes, err := json.Marshal(resp)
		if err != nil {
			log.GetLogger().Error(fmt.Sprintf("Marshal response error: %s", err.Error()))
			return
		}

		replyTopic := fmt.Sprintf("sys/product/%s/status_reply", sn)
		token := client.Publish(replyTopic, 1, false, respBytes)
		token.Wait()
		log.GetLogger().Info(fmt.Sprintf("token: %s", token))
	})
	return nil
}
