package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOriginsById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	origin_id, err := strconv.ParseInt(params.ByName("origin_id"), 10, 64)
	if err != nil || origin_id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "retrieving origination context for a specific user interaction with iframe")
}
