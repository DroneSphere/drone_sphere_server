package mqtt

import (
	"context"
	"drone_sphere_server/pkg/log"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MessageHandler func(ctx context.Context, message mqtt.Message) error

type Handler interface {
	Subscribe(ctx context.Context, handler MessageHandler) error
	Unsubscribe(ctx context.Context) error
	Publish(ctx context.Context, payload interface{}) error
	Size() int
}

type Config struct {
	Protocol string
	Broker   string
	Port     int
	Username string
	Password string
}

type MQTT struct {
	Client        mqtt.Client
	Config        Config
	TopicHandlers map[string]Handler
}

func New(conf Config) *MQTT {
	return &MQTT{
		Config:        conf,
		TopicHandlers: make(map[string]Handler),
	}
}

func (c *MQTT) Init() error {
	opts := mqtt.NewClientOptions()
	connectAddress := fmt.Sprintf("%s://%s:%d", c.Config.Protocol, c.Config.Broker, c.Config.Port)
	opts.AddBroker(connectAddress)
	opts.SetUsername(c.Config.Username)
	opts.SetPassword(c.Config.Password)
	opts.SetClientID(c.generateClientID())
	opts.SetKeepAlive(time.Second * 60)
	log.GetLogger().Info(fmt.Sprintf("Connecting to %v", fmt.Sprintf("%+v", opts)))

	c.Client = mqtt.NewClient(opts)
	token := c.Client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.GetLogger().Error("Connect Failed")
		return token.Error()
	}

	return nil
}

func (c *MQTT) generateClientID() string {
	return fmt.Sprintf("go-client-%d", rand.Int())
}

func (c *MQTT) SubscribeTopic(ctx context.Context, topic string, handler MessageHandler) error {
	if _, ok := c.TopicHandlers[topic]; !ok {
		c.TopicHandlers[topic] = &TopicHandler{
			Topic:    topic,
			Client:   c.Client,
			Handlers: make(map[string]MessageHandler),
		}
		c.TopicHandlers[topic].Subscribe(ctx, handler)
	}

	return nil
}

func (c *MQTT) UnsubscribeTopic(ctx context.Context, topic string) error {
	if _, ok := c.TopicHandlers[topic]; ok {
		c.TopicHandlers[topic].Unsubscribe(ctx)
		// 当没有handler时，删除topic
		if c.TopicHandlers[topic].Size() == 0 {
			delete(c.TopicHandlers, topic)
		}
	}

	if len(c.TopicHandlers) == 0 {
		c.Client.Disconnect(250)
	}

	return nil
}

func (c *MQTT) PublishTopic(ctx context.Context, topic string, payload interface{}) error {
	// 如果topic存在，则发布消息
	if _, ok := c.TopicHandlers[topic]; ok {
		return c.TopicHandlers[topic].Publish(ctx, payload)
	}

	return nil
}

type TopicHandler struct {
	Topic    string
	Client   mqtt.Client
	Handlers map[string]MessageHandler
}

func (t *TopicHandler) Size() int {
	return len(t.Handlers)
}

func (t *TopicHandler) Subscribe(ctx context.Context, handler MessageHandler) error {
	if token := t.Client.Subscribe(t.Topic, 1, func(client mqtt.Client, message mqtt.Message) {
		err := handler(ctx, message)
		if err != nil {
			log.GetLogger().Error(fmt.Sprintf("Handler for %s Failed", t.Topic))
			log.GetLogger().Error(fmt.Sprintf("Error: %v", err))
		}
	}); token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.GetLogger().Error(fmt.Sprintf("Subscribe to %s Failed", t.Topic))
		log.GetLogger().Error(fmt.Sprintf("Error: %v", token.Error()))
		return token.Error()
	}
	log.GetLogger().Info(fmt.Sprintf("Subscribed to %s", t.Topic))

	return nil
}

func (t *TopicHandler) Unsubscribe(ctx context.Context) error {
	if token := t.Client.Unsubscribe(t.Topic); token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.GetLogger().Error("Unsubscribe Failed")
		return token.Error()
	}

	return nil
}

func (t *TopicHandler) Publish(ctx context.Context, payload interface{}) error {
	json := fmt.Sprintf("%v", payload)
	log.GetLogger().Info(fmt.Sprintf("Publishing to %s: %s", t.Topic, json))
	if token := t.Client.Publish(t.Topic, 1, true, json); token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.GetLogger().Error("Publish Failed")
		log.GetLogger().Error(fmt.Sprintf("Error: %v", token.Error()))
		return token.Error()
	}
	return nil
}
