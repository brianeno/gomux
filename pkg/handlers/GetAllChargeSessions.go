package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brianeno/gomux/pkg/mocks"
)

// http://localhost:8000/chargesessions
func GetAllChargeSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.ChargeSession)
}
