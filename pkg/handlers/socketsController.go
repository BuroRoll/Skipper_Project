package handlers

import (
	"Skipper/pkg/models/forms"
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var SocketServer *socketio.Server

// InitSocket handler for use by app
func InitSocket() error {
	SocketServer = socketio.NewServer(nil)
	return nil
}

// SocketEvents from websocket
func (h *Handler) SocketEvents() {
	SocketServer.OnConnect("/", func(conn socketio.Conn) error {
		url := conn.URL()
		roomId := url.Query().Get("roomId")
		conn.Join(roomId)
		return nil
	})

	SocketServer.OnEvent("/", "message", func(conn socketio.Conn, msg string) {
		var input forms.MessageInput
		if err := json.Unmarshal([]byte(msg), &input); err != nil {
			fmt.Println(err)
			return
		}
		message, err := h.services.CreateMessage(input)
		if err != nil {
			return
		}
		conn.SetContext("")

		SocketServer.BroadcastToRoom("/", input.ChatID, "message", message)
		conn.Emit("message"+input.SenderID, message)
	})

	SocketServer.OnEvent("/", "bye", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		log.Println(s.Close())
	})

	SocketServer.OnError("/", func(conn socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	SocketServer.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		fmt.Println("closed:", reason)
	})
}
