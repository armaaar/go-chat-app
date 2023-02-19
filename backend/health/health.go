package health

import (
	"encoding/json"
	"net/http"

	"github.com/armaaar/go-chat-app/cors"
)

type HealthState struct {
	IsLive bool
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthState{IsLive: true})
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RegisterRoutes(baseRoute string) {
	http.Handle(baseRoute+"health", cors.Middleware(http.HandlerFunc(healthHandler)))
}
