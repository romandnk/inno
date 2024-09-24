package main

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type client struct {
	conn *websocket.Conn
}

func newClient(ctx context.Context, address string, headers http.Header) (*client, error) {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, address, headers)
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	cl := &client{conn: conn}

	return cl, nil
}

func (c *client) WriteJSON(v any) error {
	err := c.conn.WriteJSON(v)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) ReadJSON(v any) error {
	err := c.conn.ReadJSON(v)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Close() error {
	return c.conn.Close()
}
