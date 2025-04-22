package storage

import (
	"go-network-ui/models"
)

var results []models.Result

func AddResult(result models.Result) {
	results = append(results, result)
}

func GetResults() []models.Result {
	return results
}
