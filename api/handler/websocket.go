package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Websocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if string(msg) == "exit" {
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return err
		}

		fmt.Println(string(msg))

	}
	return nil
}
