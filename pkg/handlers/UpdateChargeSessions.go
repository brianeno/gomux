package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brianeno/gomux/pkg/mocks"
	"github.com/brianeno/gomux/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateChargeSessions(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedCs models.ChargeSession
	json.Unmarshal(body, &updatedCs)

	// Iterate over all the mock CS
	for index, cs := range mocks.ChargeSession {
		if cs.Id == id {
			fmt.Println("Updating", cs.Id)
			// Update and send response when CS Id matches dynamic Id
			cs.Watt = updatedCs.Watt
			cs.Vin = updatedCs.Vin

			mocks.ChargeSession[index] = cs

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
