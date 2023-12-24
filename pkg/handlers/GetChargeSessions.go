package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brianeno/gomux/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetChargeSessions(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Iterate over all the mock cs's
	for _, cs := range mocks.ChargeSession {
		if cs.Id == id {
			// If ids are equal send cs as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cs)
			break
		}
	}
}
