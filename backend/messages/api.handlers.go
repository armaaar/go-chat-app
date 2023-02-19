package messages

import (
	"encoding/json"
	"net/http"

	"github.com/armaaar/go-chat-app/database"
)

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		messages := make([]database.Message, 0)
		database.Connection.Order("id desc").Limit(100).Find(&messages)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
