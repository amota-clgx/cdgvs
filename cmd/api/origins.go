package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strconv"
)

func (app *application) getOriginsById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	origin_id, err := strconv.ParseInt(params.ByName("origin_id"), 10, 64)
	if err != nil || origin_id < 1 {
		http.NotFound(w, r)
		return
	}

	fileContents, err := os.ReadFile("resources/originData.json")
	if err != nil {
		http.Error(w, "File reading error", http.StatusNotFound)
		return
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileContents, &jsonData); err != nil {
		http.Error(w, "Error unmarshalling JSON data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)

}
