package emqx

import (
	"drone_sphere_server/pkg/log"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
)

var instance *EMQX

type EMQX struct {
	Client mqtt.Client
	Config ConnectConfig
	Topics []string
}

type ConnectConfig struct {
	Protocol string
	Broker   string
	Port     int
	Username string
	Password string
}

func New(conf ConnectConfig) *EMQX {
	instance = &EMQX{
		Config: conf,
	}
	return instance
}

func GetInstance() *EMQX {
	return instance
}

func (c *EMQX) Init() error {
	connectAddress := fmt.Sprintf("%s://%s:%d", c.Config.Protocol, c.Config.Broker, c.Config.Port)
	log.GetLogger().Info("Address " + connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(c.Config.Username)
	opts.SetPassword(c.Config.Password)
	opts.SetClientID(c.generateClientID())
	opts.SetKeepAlive(time.Second * 60)
	log.GetLogger().Info("Connecting to %v", opts)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.GetLogger().Error("Connect Failed")
		return token.Error()
	}
	log.GetLogger().Info("MQTT Connected")
	c.Client = client
	return nil
}

func (c *EMQX) generateClientID() string {
	return fmt.Sprintf("go-client-%d", rand.Int())
}
