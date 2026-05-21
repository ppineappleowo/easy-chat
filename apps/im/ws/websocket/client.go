package websocket

import (
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/json"
	"net/url"
)

type Client interface {
	Close() error

	Send(v any) error
	Read(v any) error
}
type client struct {
	*websocket.Conn
	host string

	opt dailOption
}

func NewClient(host string, opts ...DailOptions) *client {
	opt := newDailOption(opts...)

	c := client{
		Conn: nil,
		opt:  opt,
		host: host,
	}

	conn, err := c.dail()
	if err != nil {
		panic(err)
	}
	c.Conn = conn
	return &c
}

func (c *client) dail() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: c.host, Path: c.opt.pattern}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), c.opt.header)
	return conn, err
}

func (c *client) Send(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}

	//todo:增加重连发送
	conn, err := c.dail()
	if err != nil {
		return err
	}
	c.Conn = conn
	return c.WriteMessage(websocket.TextMessage, data)
}

func (c *client) Read(v any) error {
	_, msg, err := c.Conn.ReadMessage()
	if err != nil {
		return err
	}

	return json.Unmarshal(msg, v)
}
