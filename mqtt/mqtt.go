package mqtt

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Client struct
type Client struct {
	conn MQTT.Client
	opts *MQTT.ClientOptions
}

// NewClient : simple mqtt client
// url - localhost:1883
func NewClient(url string, clientID string, id string, password string) *Client {
	c := &Client{}
	opts := MQTT.NewClientOptions().AddBroker("tcp://" + url)
	opts.SetClientID(clientID)
	opts.SetUsername(id)
	opts.SetPassword(password)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetPingTimeout(1 * time.Second)

	c.opts = opts

	return c
}

// Connet : simple mqtt connect
func (c *Client) Connet() error {
	c.conn = MQTT.NewClient(c.opts)
	if token := c.conn.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Publish : simple mqtt publish
// msg - data
func (c *Client) Publish(t string, msg interface{}) {
	c.conn.Publish(t, 0, false, msg)
}

// Subscribe : simple mqtt Subscribe
// f - MQTT.MessageHandler (MQTT.Client, MQTT.Message)
func (c *Client) Subscribe(topic string, f MQTT.MessageHandler) error {
	if token := c.conn.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// OnConnect : simple onConnect. This func called when client connect or reconnect.
// f - MQTT.MessageHandler (MQTT.Client, MQTT.Message)
func (c *Client) OnConnect(topic string, f MQTT.MessageHandler) *Client {
	c.opts.OnConnect = func(client MQTT.Client) {
		if token := client.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	return c
}
