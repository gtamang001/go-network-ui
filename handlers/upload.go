package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go-network-ui/network"
	"go-network-ui/storage"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Received file: %s\n", header.Filename)

	// Save uploaded file temporarily
	tempFile, err := os.CreateTemp("", "upload-*.txt")
	if err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = tempFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}

	// Read lines (server,port) and run tests
	tempFile.Seek(0, 0)
	scanner := bufio.NewScanner(tempFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue // skip invalid lines
		}
		server := strings.TrimSpace(parts[0])
		port := strings.TrimSpace(parts[1])

		result := network.TestConnection(server, port)
		storage.AddResult(result)
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
