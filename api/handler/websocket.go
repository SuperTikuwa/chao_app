package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections map[string]*websocket.Conn

func init() {
	connections = make(map[string]*websocket.Conn)
}

func Websocket(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.String(400, "No name detected.")
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	connections[name] = conn
	broadCast(strings.Join([]string{"Player", name, "connected."}, " "))

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if string(msg) == "exit" {
			delete(connections, name)
			broadCast(strings.Join([]string{"Player", name, "disconnected."}, " "))
			break
		}

		broadCast(strings.Join([]string{name, string(msg)}, ":"))

		fmt.Println(strings.Join([]string{name, string(msg)}, ":"))

	}
	return nil
}

func broadCast(msg string) {
	for _, c := range connections {
		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Fatal(err)
		}
	}
}
