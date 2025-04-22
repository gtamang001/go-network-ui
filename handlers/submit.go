package handlers

import (
	"go-network-ui/network"
	"go-network-ui/storage"
	"net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	server := r.FormValue("server")
	port := r.FormValue("port")

	result := network.TestConnection(server, port)
	storage.AddResult(result)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
