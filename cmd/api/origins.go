package main

import (
	"fmt"
	"net/http"
)

func (app *application) getOriginsById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Retrieving origination context...")
}
