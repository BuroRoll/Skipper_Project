package handlers

import (
	"Skipper/pkg/models/forms"
	"encoding/json"
	"fmt"
	"github.com/alexandrevicenzi/go-sse"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var SocketServer *socketio.Server
var SseNotification *sse.Server

func InitSocket() error {
	SocketServer = socketio.NewServer(nil)
	return nil
}

func InitSseServe(sse *sse.Server) {
	SseNotification = sse
}

func (h *Handler) SocketEvents() {
	SocketServer.OnConnect("/", func(conn socketio.Conn) error {
		url := conn.URL()
		roomId := url.Query().Get("roomId")
		token := url.Query().Get("token")
		userId, _, err := h.services.Authorization.ParseToken(token)
		if err != nil {
			return nil
		}
		conn.Join(roomId)
		_ = h.services.ReadMessages(roomId, userId)
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
		SendMsgNotification(message, input.ReceiverID)
		fmt.Println(message)
		SocketServer.BroadcastToRoom("/", input.ChatID, "message", message)
		conn.Emit("message"+input.SenderID, message)
	})

	//SocketServer.OnEvent("/", "read_messages", func(s socketio.Conn, msg string) {
	//	var input forms.ReadChatInput
	//	if err := json.Unmarshal([]byte(msg), &input); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	userId, _, err := h.services.Authorization.ParseToken(s.RemoteHeader()["Authorization"][0][7:])
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	err = h.services.ReadMessages(input.ChatId, userId)
	//	if err != nil {
	//		SocketServer.BroadcastToRoom("/", input.ChatId, "read_messages", "read failed")
	//		return
	//	}
	//	SocketServer.BroadcastToRoom("/", input.ChatId, "read_messages", "read successfully")
	//})

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

func SendMsgNotification(message string, userId string) {
	SseNotification.SendMessage("/notifications/message/"+userId, sse.SimpleMessage(message))
}
