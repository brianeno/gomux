package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/brianeno/gomux/pkg/mocks"
	"github.com/brianeno/gomux/pkg/models"
)

func AddChargeSessions(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var cs models.ChargeSession
	json.Unmarshal(body, &cs)

	// Append to the CS mocks
	mocks.ChargeSession = append(mocks.ChargeSession, cs)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
