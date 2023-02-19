package messages

import (
	"log"
	"sync"

	"github.com/armaaar/go-chat-app/database"
	"golang.org/x/net/websocket"
)

type wsMessage struct {
	Name    string
	Message string
}

var wsConnectionsMap = struct {
	sync.RWMutex
	connections map[*websocket.Conn]string
}{connections: make(map[*websocket.Conn]string)}

func messagesSocketHandler(ws *websocket.Conn) {
	if _, ok := wsConnectionsMap.connections[ws]; !ok {
		wsConnectionsMap.Lock()
		wsConnectionsMap.connections[ws] = ""
		wsConnectionsMap.Unlock()
	}

	for {
		var wsMsg wsMessage

		// Receive message
		if err := websocket.JSON.Receive(ws, &wsMsg); err != nil {
			log.Printf("Error: %s", err)
			break
		}

		// Register last known name
		wsConnectionsMap.Lock()
		wsConnectionsMap.connections[ws] = wsMsg.Name
		wsConnectionsMap.Unlock()

		// Create message
		message := database.Message{Name: wsMsg.Name, Message: wsMsg.Message}
		database.Connection.Create(&message)

		// Send message to all connections
		wsConnectionsMap.RLock()
		for wsConn, wsName := range wsConnectionsMap.connections {
			if err := websocket.JSON.Send(wsConn, message); err != nil {
				log.Printf("Error sending message to '%s': %s", wsName, err)
			}
		}
		wsConnectionsMap.RUnlock()
	}
	wsConnectionsMap.Lock()
	delete(wsConnectionsMap.connections, ws)
	wsConnectionsMap.Unlock()
	ws.Close()
}
