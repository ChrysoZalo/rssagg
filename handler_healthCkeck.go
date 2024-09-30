package main

import "net/http"

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, struct{}{})
}
