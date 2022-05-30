package handlers

import "github.com/alexandrevicenzi/go-sse"

var SseNotification *sse.Server

func InitSseServe(sse *sse.Server) {
	SseNotification = sse
}

func SendMsgNotification(message string, userId string) {
	SseNotification.SendMessage("/notifications/message/"+userId, sse.SimpleMessage(message))
}

func SendClassNotification(message string, userId string) {
	SseNotification.SendMessage("/notifications/class/"+userId, sse.SimpleMessage(message))
}
