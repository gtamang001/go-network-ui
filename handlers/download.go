package handlers

import (
	"encoding/csv"
	"go-network-ui/storage"
	"net/http"
	"time"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=network_results.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"Server", "Port", "Status", "ResponseTime", "Timestamp"})

	for _, result := range storage.GetResults() {
		writer.Write([]string{
			result.Server,
			result.Port,
			result.Status,
			result.ResponseTime.String(),
			result.Timestamp.Format(time.RFC3339),
		})
	}
}
