package handlers

import "net/http"

// Roothandler handles the root route
// if we name functon with lowercase then go recognize as private function
func RootHandler(w http.ResponseWriter, r *http.Request) { //interface,pointer passed through server

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running\n"))
}
