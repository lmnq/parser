package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/lmnq/parser/internal/reqs/estp"
)

type EstpGetAuctions struct {
	KeyWord string `json:"key_word"`
}

func getEstpAuctions(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("server: could not read request body\nERROR: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	ga := EstpGetAuctions{}
	if err := json.Unmarshal(body, &ga); err != nil {
		log.Printf("server: could not parse request body\nERROR: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	res, err := estp.GetAuctions(ga.KeyWord)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(res.PageProps.Announces)
	if err != nil {
		log.Printf("server: could not parse data into response body\nERROR: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
