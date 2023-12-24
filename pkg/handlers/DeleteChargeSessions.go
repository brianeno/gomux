package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brianeno/gomux/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteChargeSessions(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Iterate over all the mock charge sessions
	for index, cs := range mocks.ChargeSession {
		if cs.Id == id {
			// Delete bcs and if the cs Id matches Id
			mocks.ChargeSession = append(mocks.ChargeSession[:index], mocks.ChargeSession[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
