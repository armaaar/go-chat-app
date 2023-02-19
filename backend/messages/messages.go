package messages

import (
	"net/http"

	"github.com/armaaar/go-chat-app/cors"
	"golang.org/x/net/websocket"
)

func RegisterRoutes(baseRoute string) {
	http.Handle(baseRoute+"messages", cors.Middleware(http.HandlerFunc(getMessagesHandler)))
	http.Handle(baseRoute+"messages/ws", websocket.Handler(messagesSocketHandler))
}
