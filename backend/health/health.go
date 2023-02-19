package health

import (
	"encoding/json"
	"net/http"
)

type HealthState struct {
	IsLive bool
}

func healthRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthState{IsLive: true})
}

func RegisterRoutes(baseRoute string) {
	http.HandleFunc(baseRoute, healthRoute)
}
